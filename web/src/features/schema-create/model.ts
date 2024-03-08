import { tagModel } from "@/entities/tag";
import { homeRoute } from "@/shared/routing";
import { SchemaField, SchemaKind, SchemaToCreate, createSchema } from "@/shared/api";
import { rules } from "@/shared/lib";
import { createEffect, createEvent, createStore, sample } from "effector";
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
  target: [schemaForm.reset, $creationError.reinit, homeRoute.open],
});

sample({
  source: createSchemaFx.doneData,
  filter: (schemaResponse) => !schemaResponse.success,
  fn: (schemaResponse) => schemaResponse.error,
  target: $creationError,
});
