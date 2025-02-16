import { tagModel } from "@/entities/tag";
import { tagsRoute } from "@/shared/routing";
import { TagToUpdate, deleteTag, editTag } from "@/shared/api";
import { rules, i18n, manageSubmit } from "@/shared/lib";
import { redirect } from "atomic-router";
import { attach, createEffect, createEvent, createStore, sample } from "effector";
import { createForm } from "effector-forms";

export const tagForm = createForm<Omit<TagToUpdate, "id">>({
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

export const $editError = createStore("");

export const editTagFx = createEffect((tag: TagToUpdate) => {
  return editTag(tag);
});

sample({
  source: tagModel.$currentTag,
  target: tagForm.setForm,
});

manageSubmit({
  form: tagForm,
  actionFx: editTagFx,
  error: $editError,
  currentModel: tagModel.$currentTag,
  route: tagsRoute,
});

const deleteTagFx = createEffect((id: number) => {
  deleteTag(id);
});

export const deleteTagClicked = createEvent<number>();
const alertDeleteFx = attach({
  source: i18n("tags.alerts.delete"),
  effect(alertText, id: number) {
    const confirmed = confirm(alertText);

    return confirmed ? { confirmed: true, id: id } : { confirmed: false, id: id };
  },
});

sample({
  clock: deleteTagClicked,
  target: alertDeleteFx,
});

sample({
  source: tagModel.$currentTag,
  clock: alertDeleteFx.doneData,
  filter: (_, { confirmed }) => confirmed,
  fn: ({ id }) => id,
  target: deleteTagFx,
});

redirect({
  clock: deleteTagFx.done,
  route: tagsRoute,
});
