package usecases

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const httpTimeout = 60 * time.Second

type InsightsAdapter interface {
	GenerateContent(ctx context.Context, prompt string, httpClient *http.Client) (string, error)
}

type InsightsUseCase struct {
	insightsAdapter InsightsAdapter
	proxyHost       string
	proxyPort       string
	proxyUser       string
	proxyPass       string
}

func NewInsightsUseCase(insightsAdapter InsightsAdapter, proxyUser, proxyPass, proxyHost, proxyPort string) *InsightsUseCase {
	return &InsightsUseCase{
		insightsAdapter: insightsAdapter,
		proxyHost:       proxyHost,
		proxyPort:       proxyPort,
		proxyUser:       proxyUser,
		proxyPass:       proxyPass,
	}
}

func (uc *InsightsUseCase) Enabled() bool {
	return uc.insightsAdapter != nil
}

func (uc *InsightsUseCase) GenerateContent(ctx context.Context, prompt string) (string, error) {
	httpClient, err := uc.httpClient(httpTimeout)
	if err != nil {
		return "", fmt.Errorf("creating HTTP client: %w", err)
	}

	return uc.insightsAdapter.GenerateContent(ctx, prompt, httpClient)
}

func (uc *InsightsUseCase) httpClient(timeout time.Duration) (*http.Client, error) {
	if uc.proxyHost != "" && uc.proxyPort != "" {
		return uc.proxyHttpClient(timeout)
	}

	return uc.baseHttpClient(timeout)
}

func (uc *InsightsUseCase) baseHttpClient(timeout time.Duration) (*http.Client, error) {
	client := &http.Client{
		Timeout: timeout,
	}

	return client, nil
}

func (uc *InsightsUseCase) proxyHttpClient(timeout time.Duration) (*http.Client, error) {

	var proxyStr string
	if uc.proxyUser != "" && uc.proxyPass != "" {
		proxyStr = fmt.Sprintf("http://%s:%s@%s:%s", uc.proxyUser, uc.proxyPass, uc.proxyHost, uc.proxyPort)
	} else {
		proxyStr = fmt.Sprintf("http://%s:%s", uc.proxyHost, uc.proxyPort)
	}

	proxyURL, err := url.Parse(proxyStr)
	if err != nil {
		return nil, fmt.Errorf("parsing proxy url: %w", err)
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   timeout,
	}

	return client, nil
}
