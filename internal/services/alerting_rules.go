package services

import (
	"context"
	"fmt"
	"log"
	"moonlogs/internal/entities"
	"moonlogs/internal/usecases"
	"sync"
	"time"
)

type AlertingRulesService struct {
	ctx                 context.Context
	mu                  sync.RWMutex
	rulesMeta           map[int]ruleMeta
	alertingRuleUseCase *usecases.AlertingRuleUseCase
	recordUseCase       *usecases.RecordUseCase
	incidentUseCase     *usecases.IncidentUseCase
}

type ruleMeta struct {
	timer    *time.Timer
	interval time.Duration
}

func NewAlertingRulesService(ctx context.Context, aruc *usecases.AlertingRuleUseCase, ruc *usecases.RecordUseCase, iuc *usecases.IncidentUseCase) *AlertingRulesService {
	return &AlertingRulesService{
		ctx:                 ctx,
		alertingRuleUseCase: aruc,
		rulesMeta:           make(map[int]ruleMeta),
		recordUseCase:       ruc,
		incidentUseCase:     iuc,
	}
}

func (aruc *AlertingRulesService) Sched() error {
	alertingRules, err := aruc.alertingRuleUseCase.GetAllRules(aruc.ctx)
	if err != nil {
		return fmt.Errorf("getting alerting rules: %w", err)
	}

	for _, rule := range alertingRules {
		if !rule.Enabled {
			continue
		}

		collector := aruc.buildCollector(rule)

		meta, ok := aruc.ruleInfo(rule.ID)
		if !ok {
			newTimer := time.NewTimer(rule.Interval.Duration)
			aruc.addRuleInfo(rule.ID, newTimer, rule.Interval.Duration)
			go aruc.runCollector(newTimer, collector)
			continue
		}

		// Reset rule timer to new interval. Forgive the offset
		if meta.interval != rule.Interval.Duration {
			aruc.updateRuleInfo(rule.ID, rule.Interval.Duration)
		}
	}

	return nil
}

func (aruc *AlertingRulesService) runCollector(timer *time.Timer, collector Collector) {
	<-timer.C
	defer func() {
		aruc.removeRuleInfo(collector.rule.ID)
	}()

	err := collector.trigger()
	if err != nil {
		log.Printf("collector of rule with ID %v failed: %v\n", collector.rule.ID, err)
	}
}

func (cl *Collector) trigger() error {
	to := time.Now()
	from := to.Add(-cl.rule.AggregationTimeWindow.Duration)

	filter := entities.RecordFilter{
		Level:        cl.rule.FilterLevel,
		SchemaIDs:    cl.rule.FilterSchemaIDs,
		SchemaFields: cl.rule.FilterSchemaFields,
		SchemaKinds:  cl.rule.FilterSchemaKinds,
		From:         from,
		To:           to,
	}

	aggregation := entities.RecordAggregation{
		GroupBy: cl.rule.AggregationGroupBy,
		Type:    cl.rule.AggregationType,
	}

	aggregatedGroups, err := cl.recordUseCase.AggregateRecords(cl.ctx, filter, aggregation)
	if err != nil {
		return fmt.Errorf("failed aggregating records: %w", err)
	}

	for _, group := range aggregatedGroups {
		if !isIncident(cl.rule.Condition, cl.rule.Threshold, int(group.Count)) {
			continue
		}

		payload := entities.Incident{
			RuleID: cl.rule.ID,
			Keys:   group.Keys,
			Count:  int(group.Count),
			TTL:    entities.RecordTime{Time: time.Now().Add(cl.rule.Interval.Duration)},
		}

		_, err := cl.incidentUseCase.CreateIncident(cl.ctx, payload)
		if err != nil {
			log.Printf("failed to create incident{keys: %v, count: %v, ttl: %v}: %v", payload.Keys, payload.Count, payload.TTL, err)
		}
	}

	return nil
}

func isIncident(op entities.AlertCondition, threshold int, got int) bool {
	switch op {
	case entities.LTAlertCondition:
		return got < threshold
	case entities.GTAlertCondition:
		return got > threshold
	case entities.EQAlertCondition:
		return got == threshold
	}

	return false
}

func (aruc *AlertingRulesService) buildCollector(rule *entities.AlertingRule) Collector {
	return Collector{
		ctx:                 aruc.ctx,
		alertingRuleUseCase: aruc.alertingRuleUseCase,
		recordUseCase:       aruc.recordUseCase,
		incidentUseCase:     aruc.incidentUseCase,
		rule:                rule,
	}
}

type Collector struct {
	ctx                 context.Context
	alertingRuleUseCase *usecases.AlertingRuleUseCase
	recordUseCase       *usecases.RecordUseCase
	incidentUseCase     *usecases.IncidentUseCase
	rule                *entities.AlertingRule
}

func (aruc *AlertingRulesService) ruleInfo(id int) (ruleMeta, bool) {
	aruc.mu.RLock()
	defer aruc.mu.RUnlock()

	mete, ok := aruc.rulesMeta[id]
	return mete, ok
}

func (aruc *AlertingRulesService) addRuleInfo(id int, timer *time.Timer, interval time.Duration) {
	aruc.mu.Lock()
	defer aruc.mu.Unlock()

	aruc.rulesMeta[id] = ruleMeta{
		timer:    timer,
		interval: interval,
	}
}

func (aruc *AlertingRulesService) updateRuleInfo(id int, interval time.Duration) {
	aruc.mu.Lock()
	defer aruc.mu.Unlock()

	meta := aruc.rulesMeta[id]
	meta.interval = interval
	meta.timer.Reset(interval)
	aruc.rulesMeta[id] = meta
}

func (aruc *AlertingRulesService) removeRuleInfo(id int) {
	aruc.mu.Lock()
	defer aruc.mu.Unlock()

	delete(aruc.rulesMeta, id)
}
