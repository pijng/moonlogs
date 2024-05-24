import { actionsRoute } from "@/shared/routing";
import { ActionToCreate, Condition, createAction } from "@/shared/api";
import { rules } from "@/shared/lib";
import { createEffect, createEvent, createStore, sample } from "effector";
import { createForm } from "effector-forms";

const addCondition = createEvent<Condition>();
const removeCondition = createEvent<number>();
const conditionAttributeChanged = createEvent<{ attribute: string; idx: number }>();
const conditionOperationChanged = createEvent<{ operation: Condition["operation"]; idx: number }>();
const conditionValueChanged = createEvent<{ value: string; idx: number }>();

const schemaChecked = createEvent<number>();
const schemaUnchecked = createEvent<number>();
const methodSelected = createEvent<ActionToCreate["method"]>();

export const actionForm = createForm<ActionToCreate>({
  fields: {
    name: {
      init: "",
      rules: [rules.required()],
    },
    schema_ids: {
      init: [],
      rules: [],
    },
    pattern: {
      init: "",
      rules: [rules.required()],
    },
    method: {
      init: "GET",
      rules: [rules.required()],
    },
    conditions: {
      init: [],
      rules: [],
    },
    disabled: {
      init: false,
      rules: [],
    },
  },
  validateOn: ["submit"],
});

sample({
  source: actionForm.fields.schema_ids.$value,
  clock: schemaChecked,
  fn: (schemas, newSchemaID) => [...schemas, newSchemaID],
  target: actionForm.fields.schema_ids.onChange,
});

sample({
  source: actionForm.fields.schema_ids.$value,
  clock: schemaUnchecked,
  fn: (schemas, newSchemaID) => schemas.filter((s) => s !== newSchemaID),
  target: actionForm.fields.schema_ids.onChange,
});

sample({
  clock: methodSelected,
  target: actionForm.fields.method.onChange,
});

sample({
  source: actionForm.fields.conditions.$value,
  clock: addCondition,
  fn: (conditions) => {
    const newCondition: Condition = { attribute: "", operation: "==", value: "" };
    return [...conditions, newCondition];
  },
  target: actionForm.fields.conditions.onChange,
});

sample({
  source: actionForm.fields.conditions.$value,
  clock: removeCondition,
  fn: (conditions, idx) => [...conditions.slice(0, idx), ...conditions.slice(idx + 1)],
  target: actionForm.fields.conditions.onChange,
});

sample({
  source: actionForm.fields.conditions.$value,
  clock: conditionAttributeChanged,
  fn: (conditions, { attribute, idx }) => {
    return conditions.map((condition, index) => (index === idx ? { ...condition, attribute: attribute } : condition));
  },
  target: actionForm.fields.conditions.onChange,
});

sample({
  source: actionForm.fields.conditions.$value,
  clock: conditionOperationChanged,
  fn: (conditions, { operation, idx }) => {
    return conditions.map((condition, index) => (index === idx ? { ...condition, operation: operation } : condition));
  },
  target: actionForm.fields.conditions.onChange,
});

sample({
  source: actionForm.fields.conditions.$value,
  clock: conditionValueChanged,
  fn: (conditions, { value, idx }) => {
    return conditions.map((condition, index) => (index === idx ? { ...condition, value: value } : condition));
  },
  target: actionForm.fields.conditions.onChange,
});

export const events = {
  addCondition,
  removeCondition,
  conditionAttributeChanged,
  conditionOperationChanged,
  conditionValueChanged,
  schemaChecked,
  schemaUnchecked,
  methodSelected,
};

export const $creationError = createStore("");

export const createActionFx = createEffect((action: ActionToCreate) => {
  return createAction(action);
});

sample({
  source: actionForm.formValidated,
  target: createActionFx,
});

sample({
  source: createActionFx.doneData,
  filter: (actionResponse) => actionResponse.success && Boolean(actionResponse.data.id),
  target: [actionForm.reset, $creationError.reinit, actionsRoute.open],
});

sample({
  source: createActionFx.doneData,
  filter: (actionResponse) => !actionResponse.success,
  fn: (actionResponse) => actionResponse.error,
  target: $creationError,
});
