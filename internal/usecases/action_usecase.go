package usecases

import (
	"context"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/storage"
	"slices"
)

type ActionUseCase struct {
	actionStorage storage.ActionStorage
}

func NewActionUseCase(actionStorage storage.ActionStorage) *ActionUseCase {
	return &ActionUseCase{actionStorage: actionStorage}
}

func (uc *ActionUseCase) CreateAction(ctx context.Context, action entities.Action) (*entities.Action, error) {
	if action.Name == "" {
		return nil, fmt.Errorf("failed creating action: `name` attribute is required")
	}

	if action.Pattern == "" {
		return nil, fmt.Errorf("failed creating action: `pattern` attribute is required")
	}

	isValidMethod := slices.Contains(entities.AppropriateActions, string(action.Method))
	if !isValidMethod {
		return nil, fmt.Errorf("failed creating action: `method` field should be one of: %v", entities.AppropriateActionsInfo)
	}

	var formattedConditions entities.Conditions

	for _, condition := range action.Conditions {
		var formattedCondition entities.Condition

		if condition.Attribute == "" || condition.Operation == "" || (condition.Value == "" && !slices.Contains([]string{"EXISTS", "EMPTY"}, condition.Operation)) {
			return nil, fmt.Errorf("failed creating action: `attribute`, `operation` and `value` attributes must be present for each `condition` object")
		}

		isValidOperation := slices.Contains(entities.AppropriateOperations, string(condition.Operation))
		if !isValidOperation {
			return nil, fmt.Errorf("failed creating action: `condition[].operation` field should be one of: %v", entities.AppropriateOperationsInfo)
		}

		formattedCondition.Attribute = condition.Attribute
		formattedCondition.Operation = condition.Operation
		formattedCondition.Value = condition.Value

		formattedConditions = append(formattedConditions, formattedCondition)
	}

	action.Conditions = formattedConditions

	return uc.actionStorage.CreateAction(ctx, action)
}

func (uc *ActionUseCase) GetAllActions(ctx context.Context) ([]*entities.Action, error) {
	return uc.actionStorage.GetAllActions(ctx)
}

func (uc *ActionUseCase) DeleteActionByID(ctx context.Context, id int) error {
	return uc.actionStorage.DeleteActionByID(ctx, id)
}

func (uc *ActionUseCase) GetActionByID(ctx context.Context, id int) (*entities.Action, error) {
	return uc.actionStorage.GetActionByID(ctx, id)
}

func (uc *ActionUseCase) UpdateActionByID(ctx context.Context, id int, action entities.Action) (*entities.Action, error) {
	return uc.actionStorage.UpdateActionByID(ctx, id, action)
}
