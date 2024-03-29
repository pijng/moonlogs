import { BaseResponse, del, get, post, put } from "./base";

export type Schema = {
  id: number;
  title: string;
  description: string;
  name: string;
  fields: Array<SchemaField>;
  kinds: Array<SchemaKind>;
  tag_id: number | null;
};

export type SchemaField = { title: string; name: string };
export type SchemaKind = { title: string; name: string };

export type SchemaToCreate = {
  title: string;
  description: string;
  name: string;
  retention_days: number | string;
  fields: Array<SchemaField>;
  kinds: Array<SchemaKind>;
  tag_id: number | null;
};

export type SchemaToUpdate = {
  id: number;
  title: string;
  retention_days: number | string;
  description: string;
  fields: Array<SchemaField>;
  kinds: Array<SchemaKind>;
  tag_id: number | null;
};

export interface SchemasResponse extends BaseResponse {
  data: Schema[];
}

export interface SchemaResponse extends BaseResponse {
  data: Schema;
}

export const getSchemas = (): Promise<SchemasResponse> => {
  return get({ url: "/api/schemas" });
};

export const getSchema = (id: number): Promise<SchemaResponse> => {
  return get({ url: `/api/schemas/${id}` });
};

export const querySchemas = (query: Record<string, any>): Promise<SchemasResponse> => {
  return post({ url: "/api/schemas/search", body: JSON.stringify(query) });
};

export const createSchema = (schema: SchemaToCreate): Promise<SchemaResponse> => {
  const retention_days = schema.retention_days === "" ? 0 : schema.retention_days;
  const modifiedSchema: SchemaToCreate = { ...schema, retention_days: retention_days };

  return post({ url: "/api/schemas", body: JSON.stringify(modifiedSchema) });
};

export const editSchema = (schema: SchemaToUpdate): Promise<SchemaResponse> => {
  const retention_days = schema.retention_days === "" ? 0 : schema.retention_days;
  const modifiedSchema: SchemaToUpdate = { ...schema, retention_days: retention_days };

  return put({ url: `/api/schemas/${schema.id}`, body: JSON.stringify(modifiedSchema) });
};

export const deleteSchema = (id: number): Promise<SchemaResponse> => {
  return del({ url: `/api/schemas/${id}` });
};
