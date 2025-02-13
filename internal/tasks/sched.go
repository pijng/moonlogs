package tasks

import (
	"context"
	"log"
	"moonlogs/internal/services"
	"time"
)

const baseDuration = 1 * time.Second

func RunAlertingRulesSchedTask(ctx context.Context, ars *services.AlertingRulesService) {
	ticker := time.NewTicker(baseDuration)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := ars.Sched()
			if err != nil {
				log.Printf("failed scheding alerting rules: %v\n", err)
			}
		case <-ctx.Done():
			log.Println("alerting rules schedule task is finished")
			return
		}
	}
}

func RunAlertManagerSchedTask(ctx context.Context, ams *services.AlertManagerService) {
	ticker := time.NewTicker(baseDuration)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := ams.Sched()
			if err != nil {
				log.Printf("failed scheding alertmanager: %v\n", err)
			}
		case <-ctx.Done():
			log.Println("alertmanager schedule task is finished")
			return
		}
	}
}
