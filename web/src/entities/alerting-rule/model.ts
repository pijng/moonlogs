import { AlertingRule, getRule, getRules } from "@/shared/api";
import { createEffect, createStore } from "effector";

const getRulesFx = createEffect(() => {
  return getRules();
});

const getRuleFx = createEffect((id: number) => {
  return getRule(id);
});

export const $alertingRules = createStore<AlertingRule[]>([]).on(getRulesFx.doneData, (_, rulesResponse) => rulesResponse.data);

export const $currentRule = createStore<AlertingRule>({
  id: 0,
  name: "",
  description: "",
  enabled: false,
  severity: "Error",
  interval: "1m",
  threshold: 0,
  condition: ">",
  filter_level: "Error",
  filter_schema_ids: [],
  filter_schema_fields: [],
  filter_schema_kinds: [],
  aggregation_type: "count",
  aggregation_group_by: [],
  aggregation_time_window: "",
}).on(getRuleFx.doneData, (_, ruleResponse) => ruleResponse.data);

export const effects = {
  getRulesFx,
  getRuleFx,
};

export const events = {};
