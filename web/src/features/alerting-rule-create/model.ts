import { alertingRulesRoute } from "@/shared/routing";
import { AlertingRuleToCreate, createRule } from "@/shared/api";
import { rules } from "@/shared/lib";
import { createEffect, createEvent, createStore, sample } from "effector";
import { createForm } from "effector-forms";

const schemaChecked = createEvent<number>();
const schemaUnchecked = createEvent<number>();
const schemaFieldChecked = createEvent<string>();
const schemaFieldUnchecked = createEvent<string>();
const schemaKindChecked = createEvent<string>();
const schemaKindUnchecked = createEvent<string>();
const aggregationGroupByChecked = createEvent<string>();
const aggregationGroupByUnchecked = createEvent<string>();
const aggregationTimeWindowChanged = createEvent<string>();

export const ruleForm = createForm<AlertingRuleToCreate>({
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
      rules: [rules.required()],
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
};

export const $creationError = createStore("");

export const createRuleFx = createEffect((rule: AlertingRuleToCreate) => {
  return createRule(rule);
});

sample({
  source: ruleForm.fields.filter_schema_ids.$value,
  clock: schemaChecked,
  fn: (schemas, schemaID) => [...schemas, schemaID],
  target: ruleForm.fields.filter_schema_ids.onChange,
});

sample({
  source: ruleForm.fields.filter_schema_ids.$value,
  clock: schemaUnchecked,
  fn: (schemas, schemaID) => schemas.filter((s) => s !== schemaID),
  target: ruleForm.fields.filter_schema_ids.onChange,
});

sample({
  source: ruleForm.fields.filter_schema_fields.$value,
  clock: schemaFieldChecked,
  fn: (fields, field) => [...fields, field],
  target: ruleForm.fields.filter_schema_fields.onChange,
});

sample({
  source: ruleForm.fields.filter_schema_fields.$value,
  clock: schemaFieldUnchecked,
  fn: (fields, field) => fields.filter((f) => f !== field),
  target: ruleForm.fields.filter_schema_fields.onChange,
});

sample({
  source: ruleForm.fields.filter_schema_kinds.$value,
  clock: schemaKindChecked,
  fn: (kinds, kind) => [...kinds, kind],
  target: ruleForm.fields.filter_schema_kinds.onChange,
});

sample({
  source: ruleForm.fields.filter_schema_kinds.$value,
  clock: schemaKindUnchecked,
  fn: (kinds, kind) => kinds.filter((k) => k !== kind),
  target: ruleForm.fields.filter_schema_kinds.onChange,
});

sample({
  source: ruleForm.fields.aggregation_group_by.$value,
  clock: aggregationGroupByChecked,
  fn: (groups, groupBy) => [...groups, groupBy],
  target: ruleForm.fields.aggregation_group_by.onChange,
});

sample({
  source: ruleForm.fields.aggregation_group_by.$value,
  clock: aggregationGroupByUnchecked,
  fn: (groups, groupBy) => groups.filter((g) => g !== groupBy),
  target: ruleForm.fields.aggregation_group_by.onChange,
});

sample({
  source: ruleForm.formValidated,
  target: createRuleFx,
});

sample({
  source: createRuleFx.doneData,
  filter: (ruleResponse) => ruleResponse.success && Boolean(ruleResponse.data.id),
  target: [ruleForm.reset, $creationError.reinit, alertingRulesRoute.open],
});

sample({
  source: createRuleFx.doneData,
  filter: (ruleResponse) => !ruleResponse.success,
  fn: (ruleResponse) => ruleResponse.error,
  target: $creationError,
});
