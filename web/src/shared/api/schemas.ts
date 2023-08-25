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

export const getSchemas = (): Promise<Schema[]> => {
  return get({ url: "/api/schemas", headers: {} });
};

export const querySchemas = (query: Record<string, any>): Promise<Schema[]> => {
  return post({ url: "/api/schemas/search", headers: {}, body: JSON.stringify(query) });
};
