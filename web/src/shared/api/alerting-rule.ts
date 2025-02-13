import { BaseResponse, del, get, post, put } from "./base";
import { Level } from "./logs";

export type AlertingSeverity = "Info" | "Warn" | "Error" | "Fatal";
export type AlertingCondition = ">" | "<" | "==";

export type AlertingRule = {
  id: number;
  name: string;
  description: string;
  enabled: boolean;
  severity: AlertingSeverity;
  interval: string;
  threshold: number;
  condition: AlertingCondition;
  filter_level: Level;
  filter_schema_ids: number[];
  filter_schema_fields: string[];
  filter_schema_kinds: string[];
  aggregation_type: AggregationType;
  aggregation_group_by: string[];
  aggregation_time_window: string;
};

export type AggregationType = "count";

export type AlertingRuleToCreate = {
  name: string;
  description: string;
  enabled: boolean;
  severity: AlertingSeverity;
  interval: string;
  threshold: number;
  condition: AlertingCondition;
  filter_level: Level;
  filter_schema_ids: number[];
  filter_schema_fields: string[];
  filter_schema_kinds: string[];
  aggregation_type: AggregationType;
  aggregation_group_by: string[];
  aggregation_time_window: string;
};

export type AlertingRuleToUpdate = {
  id: number;
  name: string;
  description: string;
  enabled: boolean;
  severity: AlertingSeverity;
  interval: string;
  threshold: number;
  condition: AlertingCondition;
  filter_level: Level;
  filter_schema_ids: number[];
  filter_schema_fields: string[];
  filter_schema_kinds: string[];
  aggregation_type: AggregationType;
  aggregation_group_by: string[];
  aggregation_time_window: string;
};

export interface AlertingRuleResponse extends BaseResponse {
  data: AlertingRule;
}

export interface AlertingRuleListResponse extends BaseResponse {
  data: AlertingRule[];
}

export const getRules = (): Promise<AlertingRuleListResponse> => {
  return get({ url: "/api/alerting_rules" });
};

export const getRule = (id: number): Promise<AlertingRuleResponse> => {
  return get({ url: `/api/alerting_rules/${id}` });
};

export const createRule = (rule: AlertingRuleToCreate): Promise<AlertingRuleResponse> => {
  return post({ url: "/api/alerting_rules", body: JSON.stringify(rule) });
};

export const editRule = (rule: AlertingRuleToUpdate): Promise<AlertingRuleResponse> => {
  return put({ url: `/api/alerting_rules/${rule.id}`, body: JSON.stringify(rule) });
};

export const deleteRule = (id: number): Promise<AlertingRuleResponse> => {
  return del({ url: `/api/alerting_rules/${id}` });
};
