import { Schema, getSchema, getSchemas, querySchemas } from "@/shared/api";
import { combine, createEffect, createEvent, createStore } from "effector";

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
  .on(getSchemasFx.doneData, (_, schemasResponse) => schemasResponse.data)
  .on(querySchemasFx.doneData, (_, schemasResponse) => schemasResponse.data);

export const $searchQuery = createStore("");
const queryChanged = createEvent<string>();
$searchQuery.on(queryChanged, (_, query) => query);

export const $filteredSchemas = combine([$schemas, $searchQuery], ([schemas, searchQuery]) => {
  return schemas.filter((s) => {
    const titleMatched = s.title.toLowerCase().includes(searchQuery.toLocaleLowerCase());
    const descriptionMatched = s.description.toLocaleLowerCase().includes(searchQuery.toLowerCase());
    return titleMatched || descriptionMatched;
  });
});

export const $currentSchema = createStore<Schema>({ id: 0, title: "", description: "", name: "", fields: [] }).on(
  getSchemaFx.doneData,
  (_, schemaResponse) => schemaResponse.data,
);

export const effects = {
  getSchemasFx,
  getSchemaFx,
  querySchemasFx,
};

export const events = {
  queryChanged,
};
