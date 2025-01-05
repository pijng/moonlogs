import { BaseResponse, del, get, post, put } from "./base";

export type Action = {
  id: number;
  name: string;
  pattern: string;
  method: "GET";
  conditions: Condition[];
  schema_ids: number[];
  disabled: boolean;
};

export type Condition = {
  attribute: string;
  operation: "==" | "!=" | ">" | ">=" | "<" | "<=" | "EXISTS" | "EMPTY";
  value: string;
};

export const operations: Condition["operation"][] = ["==", "!=", ">", ">=", "<", "<=", "EXISTS", "EMPTY"];
export const nonCmpOperations: Condition["operation"][] = ["EXISTS", "EMPTY"];

export type ActionToCreate = {
  name: string;
  pattern: string;
  method: "GET";
  conditions: Condition[];
  schema_ids: number[];
  disabled: boolean;
};

export type ActionToUpdate = {
  id: number;
  name: string;
  pattern: string;
  method: "GET";
  conditions: Condition[];
  schema_ids: number[];
  disabled: boolean;
};

export interface ActionResponse extends BaseResponse {
  data: Action;
}

export interface ActionListResponse extends BaseResponse {
  data: Action[];
}

export const getActions = (): Promise<ActionListResponse> => {
  return get({ url: "/api/actions" });
};

export const getAction = (id: number): Promise<ActionResponse> => {
  return get({ url: `/api/actions/${id}` });
};

export const createAction = (action: ActionToCreate): Promise<ActionResponse> => {
  return post({ url: "/api/actions", body: JSON.stringify(action) });
};

export const editAction = (action: ActionToUpdate): Promise<ActionResponse> => {
  return put({ url: `/api/actions/${action.id}`, body: JSON.stringify(action) });
};

export const deleteAction = (id: number): Promise<ActionResponse> => {
  return del({ url: `/api/actions/${id}` });
};
