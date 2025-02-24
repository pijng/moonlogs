import { Schema, getSchema, getSchemas, querySchemas } from "@/shared/api";
import { createEffect, createStore } from "effector";

const getSchemasFx = createEffect(() => {
  return getSchemas();
});

const getSchemaFx = createEffect((id: number) => {
  return getSchema(id);
});

const querySchemasFx = createEffect((query: Record<string, any>) => {
  return querySchemas(query);
});

export const $schemas = createStore<Schema[]>([])
  .on(getSchemasFx.doneData, (_, schemasResponse) => (!!schemasResponse ? schemasResponse.data : []))
  .on(querySchemasFx.doneData, (_, schemasResponse) => schemasResponse.data);

export const $currentSchema = createStore<Schema>({
  id: 0,
  title: "",
  description: "",
  name: "",
  fields: [],
  kinds: [],
  tag_id: null,
}).on(getSchemaFx.doneData, (_, schemaResponse) => schemaResponse.data);

export const effects = {
  getSchemasFx,
  getSchemaFx,
  querySchemasFx,
};
