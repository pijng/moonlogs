import { schemaModel } from "@/entities/schema";
import { tagModel } from "@/entities/tag";
import { homeRoute } from "@/shared/routing";
import { SchemaField, SchemaKind, SchemaToUpdate, deleteSchema, editSchema } from "@/shared/api";
import { rules, i18n } from "@/shared/lib";
import { redirect } from "atomic-router";
import { attach, createEffect, createEvent, createStore, sample } from "effector";
import { createForm } from "effector-forms";

const addField = createEvent<SchemaField>();
const removeField = createEvent<number>();
const fieldTitleChanged = createEvent<{ title: string; idx: number }>();
const fieldNameChanged = createEvent<{ name: string; idx: number }>();

const addKind = createEvent<SchemaKind>();
const removeKind = createEvent<number>();
const kindTitleChanged = createEvent<{ title: string; idx: number }>();
const kindNameChanged = createEvent<{ name: string; idx: number }>();

const tagSelected = createEvent<string>();

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
    retention_days: {
      init: "",
      rules: [],
    },
    fields: {
      init: [],
      rules: [],
    },
    kinds: {
      init: [],
      rules: [],
    },
    tag_id: {
      init: null,
      rules: [],
    },
  },
  validateOn: ["submit"],
});

sample({
  source: tagModel.$tags,
  clock: tagSelected,
  fn: (tags, selectedTag) => {
    return tags.find((t) => t.name === selectedTag)?.id || null;
  },
  target: schemaForm.fields.tag_id.onChange,
});

sample({
  source: schemaForm.fields.fields.$value,
  clock: addField,
  fn: (fields) => [...fields, { title: "", name: "" }],
  target: schemaForm.fields.fields.onChange,
});

sample({
  source: schemaForm.fields.fields.$value,
  clock: removeField,
  fn: (fields, idx) => [...fields.slice(0, idx), ...fields.slice(idx + 1)],
  target: schemaForm.fields.fields.onChange,
});

sample({
  source: schemaForm.fields.fields.$value,
  clock: fieldTitleChanged,
  fn: (fields, { title, idx }) => {
    return fields.map((field, index) => (index === idx ? { ...field, title: title } : field));
  },
  target: schemaForm.fields.fields.onChange,
});

sample({
  source: schemaForm.fields.fields.$value,
  clock: fieldNameChanged,
  fn: (fields, { name, idx }) => {
    return fields.map((field, index) => (index === idx ? { ...field, name: name } : field));
  },
  target: schemaForm.fields.fields.onChange,
});

sample({
  source: schemaForm.fields.kinds.$value,
  clock: addKind,
  fn: (kinds) => [...(kinds || []), { title: "", name: "" }],
  target: schemaForm.fields.kinds.onChange,
});

sample({
  source: schemaForm.fields.kinds.$value,
  clock: removeKind,
  fn: (kinds, idx) => [...kinds.slice(0, idx), ...kinds.slice(idx + 1)],
  target: schemaForm.fields.kinds.onChange,
});

sample({
  source: schemaForm.fields.kinds.$value,
  clock: kindTitleChanged,
  fn: (kinds, { title, idx }) => {
    return kinds.map((kind, index) => (index === idx ? { ...kind, title: title } : kind));
  },
  target: schemaForm.fields.kinds.onChange,
});

sample({
  source: schemaForm.fields.kinds.$value,
  clock: kindNameChanged,
  fn: (kinds, { name, idx }) => {
    return kinds.map((kind, index) => (index === idx ? { ...kind, name: name } : kind));
  },
  target: schemaForm.fields.kinds.onChange,
});

export const events = {
  addField,
  removeField,
  fieldTitleChanged,
  fieldNameChanged,
  addKind,
  removeKind,
  kindTitleChanged,
  kindNameChanged,
  tagSelected,
};

export const $editError = createStore("");

export const editSchemaFx = createEffect((schema: SchemaToUpdate) => {
  return editSchema(schema);
});

sample({
  source: schemaModel.$currentSchema,
  target: [schemaForm.setForm],
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

export const deleteSchemaFx = createEffect((id: number) => {
  deleteSchema(id);
});

export const deleteSchemaClicked = createEvent<number>();
const alertDeleteFx = attach({
  source: i18n("log_groups.alerts.delete"),
  effect(alertText, id: number) {
    const confirmed = confirm(alertText);

    return confirmed ? { confirmed: true, id: id } : { confirmed: false, id: id };
  },
});

sample({
  clock: deleteSchemaClicked,
  target: alertDeleteFx,
});

sample({
  source: schemaModel.$currentSchema,
  clock: alertDeleteFx.doneData,
  filter: (_, { confirmed }) => confirmed,
  fn: ({ id }) => id,
  target: deleteSchemaFx,
});

redirect({
  clock: deleteSchemaFx.done,
  route: homeRoute,
});
