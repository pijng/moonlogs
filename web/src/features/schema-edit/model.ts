import { schemaModel } from "@/entities/schema";
import { homeRoute } from "@/routing/shared";
import { SchemaToUpdate, editSchema } from "@/shared/api";
import { rules } from "@/shared/lib";
import { createEffect, createStore, sample } from "effector";
import { createForm } from "effector-forms";

export const schemaForm = createForm<Omit<SchemaToUpdate, "id">>({
  fields: {
    title: {
      init: "",
      rules: [rules.required()],
    },
    description: {
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

export const $editError = createStore("");

export const editSchemaFx = createEffect((schema: SchemaToUpdate) => {
  return editSchema(schema);
});

sample({
  source: schemaModel.$currentSchema,
  target: schemaForm.setInitialForm,
});

sample({
  source: schemaModel.$currentSchema,
  clock: schemaForm.formValidated,
  fn: (currentSchema, schemaToEdit) => {
    return { ...schemaToEdit, id: currentSchema.id };
  },
  target: editSchemaFx,
});

sample({
  source: editSchemaFx.doneData,
  filter: (schemaResponse) => schemaResponse.success && Boolean(schemaResponse.data.id),
  target: [schemaForm.reset, homeRoute.open],
});

sample({
  source: editSchemaFx.doneData,
  filter: (schemaResponse) => !schemaResponse.success,
  fn: (schemaResponse) => schemaResponse.error,
  target: $editError,
});
