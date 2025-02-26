import { actionsRoute } from "@/shared/routing";
import { ActionToCreate, ActionToUpdate, Condition, deleteAction, editAction } from "@/shared/api";
import { rules, i18n, bindFieldList, manageSubmit } from "@/shared/lib";
import { attach, createEffect, createEvent, createStore, sample } from "effector";
import { createForm } from "effector-forms";
import { actionModel } from "@/entities/action";
import { redirect } from "atomic-router";

const addCondition = createEvent<Condition>();
const removeCondition = createEvent<number>();
const conditionAttributeChanged = createEvent<{ attribute: string; idx: number }>();
const conditionOperationChanged = createEvent<{ operation: Condition["operation"]; idx: number }>();
const conditionValueChanged = createEvent<{ value: string; idx: number }>();

const schemaChecked = createEvent<number>();
const schemaUnchecked = createEvent<number>();
const methodSelected = createEvent<ActionToCreate["method"]>();

export const actionForm = createForm<Omit<ActionToUpdate, "id">>({
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

bindFieldList({ field: actionForm.fields.schema_ids, added: schemaChecked, removed: schemaUnchecked });

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

export const $editError = createStore("");

export const editActionFx = createEffect((action: ActionToUpdate) => {
  return editAction(action);
});

sample({
  source: actionModel.$currentAction,
  target: [actionForm.setForm],
});

manageSubmit({
  form: actionForm,
  currentModel: actionModel.$currentAction,
  actionFx: editActionFx,
  error: $editError,
  route: actionsRoute,
});

export const deleteActionFx = createEffect((id: number) => {
  deleteAction(id);
});

export const deleteActionClicked = createEvent<number>();
const alertDeleteFx = attach({
  source: i18n("actions.alerts.delete"),
  effect(alertText, id: number) {
    const confirmed = confirm(alertText);

    return confirmed ? { confirmed: true, id: id } : { confirmed: false, id: id };
  },
});

sample({
  clock: deleteActionClicked,
  target: alertDeleteFx,
});

sample({
  source: actionModel.$currentAction,
  clock: alertDeleteFx.doneData,
  filter: (_, { confirmed }) => confirmed,
  fn: ({ id }) => id,
  target: deleteActionFx,
});

redirect({
  clock: deleteActionFx.done,
  route: actionsRoute,
});
