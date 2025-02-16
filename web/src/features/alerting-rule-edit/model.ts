import { alertingRulesRoute } from "@/shared/routing";
import { AlertingRuleToUpdate, deleteRule, editRule } from "@/shared/api";
import { bindFieldList, i18n, manageSubmit, rules } from "@/shared/lib";
import { attach, createEffect, createEvent, createStore, sample } from "effector";
import { createForm } from "effector-forms";
import { alertingRuleModel } from "@/entities/alerting-rule";
import { redirect } from "atomic-router";

const schemaChecked = createEvent<number>();
const schemaUnchecked = createEvent<number>();
const schemaFieldChecked = createEvent<string>();
const schemaFieldUnchecked = createEvent<string>();
const schemaKindChecked = createEvent<string>();
const schemaKindUnchecked = createEvent<string>();
const aggregationGroupByChecked = createEvent<string>();
const aggregationGroupByUnchecked = createEvent<string>();
const aggregationTimeWindowChanged = createEvent<string>();

export const deleteRuleClicked = createEvent<number>();

export const ruleForm = createForm<Omit<AlertingRuleToUpdate, "id">>({
  fields: {
    name: {
      init: "",
      rules: [rules.required()],
    },
    description: {
      init: "",
      rules: [rules.required()],
    },
    enabled: {
      init: true,
      rules: [],
    },
    severity: {
      init: "Error",
      rules: [rules.required()],
    },
    interval: {
      init: "1m",
      rules: [rules.required()],
    },
    threshold: {
      init: 0,
      rules: [rules.required()],
    },
    condition: {
      init: ">",
      rules: [rules.required()],
    },
    filter_level: {
      init: "Error",
      rules: [],
    },
    filter_schema_ids: {
      init: [],
    },
    filter_schema_fields: {
      init: [],
    },
    filter_schema_kinds: {
      init: [],
    },
    aggregation_type: {
      init: "count",
      rules: [rules.required()],
    },
    aggregation_group_by: {
      init: [],
      rules: [rules.required()],
    },
    aggregation_time_window: {
      init: "5m",
      rules: [rules.required()],
    },
  },
  validateOn: ["submit"],
});

export const events = {
  schemaChecked,
  schemaUnchecked,
  schemaFieldChecked,
  schemaFieldUnchecked,
  schemaKindChecked,
  schemaKindUnchecked,
  aggregationGroupByChecked,
  aggregationGroupByUnchecked,
  aggregationTimeWindowChanged,
  deleteRuleClicked,
};

bindFieldList({ field: ruleForm.fields.filter_schema_ids, added: schemaChecked, removed: schemaUnchecked });
bindFieldList({ field: ruleForm.fields.filter_schema_fields, added: schemaFieldChecked, removed: schemaFieldUnchecked });
bindFieldList({ field: ruleForm.fields.filter_schema_kinds, added: schemaKindChecked, removed: schemaKindUnchecked });
bindFieldList({
  field: ruleForm.fields.aggregation_group_by,
  added: aggregationGroupByChecked,
  removed: aggregationGroupByUnchecked,
});

export const $editError = createStore("");

export const editRuleFx = createEffect((rule: AlertingRuleToUpdate) => {
  return editRule(rule);
});

sample({
  source: alertingRuleModel.$currentRule,
  target: [ruleForm.setForm],
});

manageSubmit({
  form: ruleForm,
  currentModel: alertingRuleModel.$currentRule,
  actionFx: editRuleFx,
  error: $editError,
  route: alertingRulesRoute,
});

export const deleteRuleFx = createEffect((id: number) => {
  deleteRule(id);
});

const alertDeleteFx = attach({
  source: i18n("alerting_rules.alerts.delete"),
  effect(alertText, id: number) {
    const confirmed = confirm(alertText);

    return confirmed ? { confirmed: true, id: id } : { confirmed: false, id: id };
  },
});

sample({
  clock: deleteRuleClicked,
  target: alertDeleteFx,
});

sample({
  source: alertingRuleModel.$currentRule,
  clock: alertDeleteFx.doneData,
  filter: (_, { confirmed }) => confirmed,
  fn: ({ id }) => id,
  target: deleteRuleFx,
});

redirect({
  clock: deleteRuleFx.done,
  route: alertingRulesRoute,
});
