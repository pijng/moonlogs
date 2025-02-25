package usecases

import "context"

type InsightsAdapter interface {
	GenerateContent(ctx context.Context, prompt string) (string, error)
}

type InsightsUseCase struct {
	insightsAdapter InsightsAdapter
}

func NewInsightsUseCase(insightsAdapter InsightsAdapter) *InsightsUseCase {
	return &InsightsUseCase{insightsAdapter: insightsAdapter}
}

func (uc *InsightsUseCase) Enabled() bool {
	return uc.insightsAdapter != nil
}

func (uc *InsightsUseCase) GenerateContent(ctx context.Context, prompt string) (string, error) {
	return uc.insightsAdapter.GenerateContent(ctx, prompt)
}
