import { userModel } from "@/entities/user";
import { membersRoute } from "@/shared/routing";
import { UserToUpdate, editUser } from "@/shared/api";
import { deleteUser } from "@/shared/api/users";
import { rules } from "@/shared/lib";
import { i18n } from "@/shared/lib/i18n";
import { redirect } from "atomic-router";
import { attach, createEffect, createEvent, createStore, sample } from "effector";
import { createForm } from "effector-forms";

const tagChecked = createEvent<number>();
const tagUnchecked = createEvent<number>();

export const memberForm = createForm<Omit<UserToUpdate, "id">>({
  fields: {
    name: {
      init: "",
      rules: [rules.required()],
    },
    email: {
      init: "",
      rules: [rules.required(), rules.email()],
    },
    is_revoked: {
      init: false,
      rules: [],
    },
    role: {
      init: "Member",
      rules: [rules.required()],
    },
    tag_ids: {
      init: [],
      rules: [],
    },
    password: {
      init: "",
      rules: [rules.password()],
    },
    passwordConfirmation: {
      init: "",
      rules: [rules.password()],
    },
  },
  validateOn: ["submit"],
});

export const $editError = createStore("");

export const editUserFx = createEffect((user: UserToUpdate) => {
  return editUser(user);
});

sample({
  source: userModel.$currentUser,
  target: memberForm.setInitialForm,
});

sample({
  source: userModel.$currentUser,
  clock: memberForm.formValidated,
  fn: (currentUser, userToEdit) => {
    return { ...userToEdit, id: currentUser.id };
  },
  target: editUserFx,
});

sample({
  source: editUserFx.doneData,
  filter: (userResponse) => userResponse.success && Boolean(userResponse.data.id),
  target: [memberForm.reset, membersRoute.open],
});

sample({
  source: editUserFx.doneData,
  filter: (userResponse) => !userResponse.success,
  fn: (userResponse) => userResponse.error,
  target: $editError,
});

export const deleteUserFx = createEffect((id: number) => {
  deleteUser(id);
});

export const deleteUserClicked = createEvent<number>();
const alertDeleteFx = attach({
  source: i18n("members.alerts.delete"),
  effect(alertText, id: number) {
    const confirmed = confirm(alertText);

    return confirmed ? { confirmed: true, id: id } : { confirmed: false, id: id };
  },
});

sample({
  clock: deleteUserClicked,
  target: alertDeleteFx,
});

sample({
  source: userModel.$currentUser,
  clock: alertDeleteFx.doneData,
  filter: (_, { confirmed }) => confirmed,
  fn: ({ id }) => id,
  target: deleteUserFx,
});

redirect({
  clock: deleteUserFx.done,
  route: membersRoute,
});

sample({
  source: memberForm.fields.tag_ids.$value,
  clock: tagChecked,
  fn: (tags, newTagID) => [...tags, newTagID],
  target: memberForm.fields.tag_ids.onChange,
});

sample({
  source: memberForm.fields.tag_ids.$value,
  clock: tagUnchecked,
  fn: (tags, newTagID) => tags.filter((t) => t !== newTagID),
  target: memberForm.fields.tag_ids.onChange,
});

export const events = {
  tagChecked,
  tagUnchecked,
};
