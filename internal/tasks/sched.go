package tasks

import (
	"context"
	"log"
	"moonlogs/internal/services"
	"time"
)

const baseDuration = 1 * time.Second

func RunAlertingRulesSchedTask(ctx context.Context, aruc *services.AlertingRulesService) {
	ticker := time.NewTicker(baseDuration)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := aruc.Sched()
			if err != nil {
				log.Printf("failed scheding alerting rules: %v\n", err)
			}
		case <-ctx.Done():
			log.Println("alerting rules schedule task is finished")
			return
		}
	}
}
