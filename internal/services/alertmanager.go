package services

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"moonlogs/internal/entities"
	"moonlogs/internal/usecases"
	"net/http"
	"net/url"
	"sync"
	"time"
)

const httpTimeout = 5 * time.Second
const maxIdleConns = 10
const maxConnsPerHost = 10
const maxIdleConnsPerHost = 5

type httpClientConfig struct {
	timeout             time.Duration
	maxIdleConns        int
	maxConnsPerHost     int
	maxIdleConnsPerHost int
}

type AlertManagerService struct {
	ctx                        context.Context
	mu                         sync.RWMutex
	incidentMeta               map[incidentHash]profileIncidents
	notificationProfileUseCase *usecases.NotificationProfileUseCase
	incidentUseCase            *usecases.IncidentUseCase
}

type incidentMeta struct {
	silencedUntil time.Time
}

type profileID int
type incidentHash string

type profileIncidents map[profileID]incidentMeta

func NewAlertManagerService(ctx context.Context, nuc *usecases.NotificationProfileUseCase, iuc *usecases.IncidentUseCase) *AlertManagerService {
	return &AlertManagerService{
		ctx:                        ctx,
		incidentUseCase:            iuc,
		notificationProfileUseCase: nuc,
		incidentMeta:               make(map[incidentHash]profileIncidents),
	}
}

func (ams *AlertManagerService) Sched() error {
	now := time.Now()

	incidents, err := ams.incidentUseCase.GetAllIncidents(ams.ctx)
	if err != nil {
		return fmt.Errorf("getting incidents: %w", err)
	}

	profiles, err := ams.notificationProfileUseCase.GetAllNotificationProfiles(ams.ctx)
	if err != nil {
		return fmt.Errorf("getting notification profiles: %w", err)
	}

	for _, incident := range incidents {
		hash, err := incident.Hash()
		if err != nil {
			log.Printf("calculating hash for incident with ID: %v failed: %v", incident.ID, err)
			continue
		}

		for _, profile := range profiles {
			if !profile.Enabled {
				continue
			}

			meta, ok := ams.incidentInfo(incidentHash(hash), profileID(profile.ID))
			if ok && meta.silencedUntil.After(now) {
				continue
			}

			if ok && meta.silencedUntil.Before(now) {
				ams.removeIncidentInfo(incidentHash(hash), profileID(profile.ID))
			}

			notifier, err := buildNotifier(incident, profile, httpClientConfig{
				timeout:             httpTimeout,
				maxIdleConns:        maxIdleConns,
				maxConnsPerHost:     maxConnsPerHost,
				maxIdleConnsPerHost: maxIdleConnsPerHost,
			})

			if err != nil {
				log.Printf("building notifier: %v", err)
				continue
			}

			go func() {
				err = notifier.notify()
				if err != nil {
					log.Printf("notification for incident:'%v' and profile:'%v' failed: %v", incident.ID, profile.Name, err)
					return
				}

				ams.addIncidentInfo(incidentHash(hash), profileID(profile.ID), profile.SilenceFor.Duration)
			}()
		}
	}

	return nil
}

type notifier struct {
	incident      *entities.Incident
	profile       *entities.NotificationProfile
	parsedPayload string
	httpClient    *http.Client
}

func buildNotifier(incident *entities.Incident, profile *entities.NotificationProfile, cfg httpClientConfig) (*notifier, error) {
	transport := &http.Transport{
		MaxIdleConns:        cfg.maxIdleConns,
		MaxConnsPerHost:     cfg.maxConnsPerHost,
		MaxIdleConnsPerHost: cfg.maxIdleConnsPerHost,
	}

	client := &http.Client{
		Timeout:   cfg.timeout,
		Transport: transport,
	}

	parsedPayload, err := incident.Message(profile.Payload)
	if err != nil {
		return nil, fmt.Errorf("building message: %w", err)
	}

	return &notifier{
		incident:      incident,
		profile:       profile,
		httpClient:    client,
		parsedPayload: parsedPayload,
	}, nil
}

func (n *notifier) notify() error {
	url, err := url.Parse(n.profile.URL)
	if err != nil {
		return fmt.Errorf("parsing notification profile url: %w", err)
	}

	req := &http.Request{
		Method: n.profile.Method,
		URL:    url,
		Header: make(http.Header),
	}

	if n.parsedPayload != "" {
		req.Body = io.NopCloser(bytes.NewBuffer([]byte(n.parsedPayload)))
	}

	if n.profile.Headers != nil {
		for _, header := range n.profile.Headers {
			req.Header.Set(header.Key, header.Value)
		}
	}

	resp, err := n.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("to read response body: %w", err)
		}

		return fmt.Errorf("received non-ok status: %v\n%s", resp.Status, string(body))
	}

	return nil
}

func (ams *AlertManagerService) incidentInfo(incidentHash incidentHash, npID profileID) (incidentMeta, bool) {
	ams.mu.RLock()
	defer ams.mu.RUnlock()

	incidentsCollection, ok := ams.incidentMeta[incidentHash]
	if !ok {
		return incidentMeta{}, false
	}

	meta, ok := incidentsCollection[npID]

	return meta, ok
}

func (ams *AlertManagerService) addIncidentInfo(incidentHash incidentHash, npID profileID, silenceFor time.Duration) {
	ams.mu.Lock()
	defer ams.mu.Unlock()

	newMeta := incidentMeta{
		silencedUntil: time.Now().Add(silenceFor),
	}

	profileIncidents, ok := ams.incidentMeta[incidentHash]
	if !ok {
		newProfileIncidents := make(map[profileID]incidentMeta)
		newProfileIncidents[npID] = newMeta
		ams.incidentMeta[incidentHash] = newProfileIncidents
		return
	}

	profileIncidents[npID] = newMeta
}

func (ams *AlertManagerService) removeIncidentInfo(incidentHash incidentHash, npID profileID) {
	ams.mu.Lock()
	defer ams.mu.Unlock()

	profileIncidents, ok := ams.incidentMeta[incidentHash]
	if !ok {
		return
	}

	delete(profileIncidents, npID)
	ams.incidentMeta[incidentHash] = profileIncidents
}
