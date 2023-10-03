import { BaseResponse, get, post, put } from "./base";

export type Schema = {
  id: number;
  title: string;
  description: string;
  name: string;
  fields: Array<{
    title: string;
    name: string;
    type?: string | number | boolean | Date;
  }>;
};

export type SchemaToCreate = {
  title: string;
  description: string;
  name: string;
  fields?: Array<{
    title: string;
    name: string;
    type?: string | number | boolean | Date;
  }>;
};

export type SchemaToUpdate = {
  id: number;
  title: string;
  description: string;
  fields?: Array<{
    title: string;
    name: string;
    type?: string | number | boolean | Date;
  }>;
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
