import { tagsRoute } from "@/routing/shared";
import { TagToCreate, createTag } from "@/shared/api";
import { rules } from "@/shared/lib";
import { createEffect, createStore, sample } from "effector";
import { createForm } from "effector-forms";

export const tagForm = createForm<TagToCreate>({
  fields: {
    name: {
      init: "",
      rules: [rules.required()],
    },
  },
  validateOn: ["submit"],
});

export const $creationError = createStore("");

export const createTagFx = createEffect((tag: TagToCreate) => {
  return createTag(tag);
});

sample({
  source: tagForm.formValidated,
  target: createTagFx,
});

sample({
  source: createTagFx.doneData,
  filter: (tagResponse) => tagResponse.success && Boolean(tagResponse.data.id),
  target: [tagForm.reset, tagsRoute.open],
});

sample({
  source: createTagFx.doneData,
  filter: (tagResponse) => !tagResponse.success,
  fn: (tagResponse) => tagResponse.error,
  target: $creationError,
});
