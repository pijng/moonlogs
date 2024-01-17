import { BaseResponse, del, get, post, put } from "./base";

export type Schema = {
  id: number;
  title: string;
  description: string;
  name: string;
  fields: Array<SchemaField>;
  kinds: Array<SchemaKind>;
};

export type SchemaField = { title: string; name: string };
export type SchemaKind = { title: string; name: string };

export type SchemaToCreate = {
  title: string;
  description: string;
  name: string;
  retention_time: number | null;
  fields: Array<SchemaField>;
  kinds: Array<SchemaKind>;
};

export type SchemaToUpdate = {
  id: number;
  title: string;
  retention_time: number | null;
  description: string;
  fields: Array<SchemaField>;
  kinds: Array<SchemaKind>;
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
  return post({ url: "/api/schemas", body: JSON.stringify(schema) });
};

export const editSchema = (schema: SchemaToUpdate): Promise<SchemaResponse> => {
  return put({ url: `/api/schemas/${schema.id}`, body: JSON.stringify(schema) });
};

export const deleteSchema = (id: number): Promise<SchemaResponse> => {
  return del({ url: `/api/schemas/${id}` });
};
