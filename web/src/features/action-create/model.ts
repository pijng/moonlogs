import { actionsRoute } from "@/shared/routing";
import { ActionToCreate, Condition, Schema, createAction } from "@/shared/api";
import { rules } from "@/shared/lib";
import { createEffect, createEvent, createStore, sample } from "effector";
import { createForm } from "effector-forms";
import { schemaModel } from "@/entities/schema";

const addCondition = createEvent<Condition>();
const removeCondition = createEvent<number>();
const conditionAttributeChanged = createEvent<{ attribute: string; idx: number }>();
const conditionOperationChanged = createEvent<{ operation: Condition["operation"]; idx: number }>();
const conditionValueChanged = createEvent<{ value: string; idx: number }>();

const schemaSelected = createEvent<string>();
const methodSelected = createEvent<ActionToCreate["method"]>();

const emptySchema = {
  id: 0,
  title: "",
  description: "",
  name: "",
  fields: [],
  kinds: [],
  tag_id: null,
};
export const $currentSchema = createStore<Schema>(emptySchema);

export const actionForm = createForm<ActionToCreate>({
  fields: {
    name: {
      init: "",
      rules: [rules.required()],
    },
    schema_id: {
      init: 0,
      rules: [rules.required()],
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
  source: schemaModel.$schemas,
  clock: schemaSelected,
  fn: (schemas, selectedSchema) => {
    return schemas.find((s) => s.name === selectedSchema)?.id || 0;
  },
  target: actionForm.fields.schema_id.onChange,
});

sample({
  source: schemaModel.$schemas,
  clock: schemaSelected,
  fn: (schemas, selectedSchema) => {
    return schemas.find((s) => s.name === selectedSchema) || emptySchema;
  },
  target: $currentSchema,
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
  schemaSelected,
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
