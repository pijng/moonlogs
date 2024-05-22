package usecases

import (
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

func (uc *ActionUseCase) CreateAction(action entities.Action) (*entities.Action, error) {
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

	if action.SchemaID == 0 {
		return nil, fmt.Errorf("failed creating action: `schema_id` attribute is required")
	}

	var formattedConditions entities.Conditions

	for _, condition := range action.Conditions {
		var formattedCondition entities.Condition

		if condition.Attribute == "" || condition.Operation == "" || condition.Value == "" {
			return nil, fmt.Errorf("failed creating action: `attribute`, `operation` and `value` attributes must be present for each `condition` object")
		}

		formattedCondition.Attribute = condition.Attribute
		formattedCondition.Operation = condition.Operation
		formattedCondition.Value = condition.Value

		formattedConditions = append(formattedConditions, formattedCondition)
	}

	action.Conditions = formattedConditions

	return uc.actionStorage.CreateAction(action)
}

func (uc *ActionUseCase) GetAllActions() ([]*entities.Action, error) {
	return uc.actionStorage.GetAllActions()
}

func (uc *ActionUseCase) DeleteActionByID(id int) error {
	return uc.actionStorage.DeleteActionByID(id)
}

func (uc *ActionUseCase) GetActionByID(id int) (*entities.Action, error) {
	return uc.actionStorage.GetActionByID(id)
}

func (uc *ActionUseCase) UpdateActionByID(id int, action entities.Action) (*entities.Action, error) {
	return uc.actionStorage.UpdateActionByID(id, action)
}
