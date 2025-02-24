import { tagsRoute } from "@/shared/routing";
import { TagToCreate, createTag } from "@/shared/api";
import { manageSubmit, rules } from "@/shared/lib";
import { createEffect, createStore } from "effector";
import { createForm } from "effector-forms";

export const tagForm = createForm<TagToCreate>({
  fields: {
    name: {
      init: "",
      rules: [rules.required()],
    },
    view_order: {
      init: "",
      rules: [],
    },
  },
  validateOn: ["submit"],
});

export const $creationError = createStore("");

export const createTagFx = createEffect((tag: TagToCreate) => {
  return createTag(tag);
});

manageSubmit({
  form: tagForm,
  actionFx: createTagFx,
  error: $creationError,
  route: tagsRoute,
});
