import { homeRoute } from "@/routing/shared";
import { SchemaToCreate, createSchema } from "@/shared/api";
import { rules } from "@/shared/lib";
import { createEffect, createStore, sample } from "effector";
import { createForm } from "effector-forms";

export const schemaForm = createForm<SchemaToCreate>({
  fields: {
    title: {
      init: "",
      rules: [rules.required()],
    },
    description: {
      init: "",
      rules: [rules.required()],
    },
    name: {
      init: "",
      rules: [rules.required()],
    },
    fields: {
      init: [{ title: "", name: "" }],
      rules: [],
    },
  },
  validateOn: ["submit"],
});

export const $creationError = createStore("");

export const createSchemaFx = createEffect((schema: SchemaToCreate) => {
  return createSchema(schema);
});

sample({
  source: schemaForm.formValidated,
  target: createSchemaFx,
});

sample({
  source: createSchemaFx.doneData,
  filter: (schemaResponse) => schemaResponse.success && Boolean(schemaResponse.data.id),
  target: homeRoute.open,
});

sample({
  source: createSchemaFx.doneData,
  filter: (schemaResponse) => !schemaResponse.success,
  fn: (schemaResponse) => schemaResponse.error,
  target: $creationError,
});
