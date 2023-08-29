import { get, post } from "./base";

export type Schema = {
  title: string;
  description: string;
  name: string;
  fields: {
    title: string;
    name: string;
    type: string | number | boolean | Date;
  }[];
};

export type SchemasReponse = {
  data: Schema[];
  meta: {
    page: number;
    count: number;
    per_page: number;
  };
};

export const getSchemas = (): Promise<SchemasReponse> => {
  return get({ url: "/api/schemas", headers: {} });
};

export const querySchemas = (query: Record<string, any>): Promise<SchemasReponse> => {
  return post({ url: "/api/schemas/search", headers: {}, body: JSON.stringify(query) });
};
