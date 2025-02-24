import { alertingRulesRoute } from "@/shared/routing";
import { AlertingRuleToCreate, createRule } from "@/shared/api";
import { bindFieldList, manageSubmit, rules } from "@/shared/lib";
import { createEffect, createEvent, createStore } from "effector";
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

bindFieldList({ field: ruleForm.fields.filter_schema_ids, added: schemaChecked, removed: schemaUnchecked });
bindFieldList({ field: ruleForm.fields.filter_schema_fields, added: schemaFieldChecked, removed: schemaFieldUnchecked });
bindFieldList({ field: ruleForm.fields.filter_schema_kinds, added: schemaKindChecked, removed: schemaKindUnchecked });
bindFieldList({
  field: ruleForm.fields.aggregation_group_by,
  added: aggregationGroupByChecked,
  removed: aggregationGroupByUnchecked,
});

manageSubmit({ form: ruleForm, actionFx: createRuleFx, error: $creationError, route: alertingRulesRoute });
