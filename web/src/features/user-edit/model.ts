import { userModel } from "@/entities/user";
import { membersRoute } from "@/routing/shared";
import { UserToUpdate, editUser } from "@/shared/api";
import { rules } from "@/shared/lib";
import { createEffect, createStore, sample } from "effector";
import { createForm } from "effector-forms";

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
    role: {
      init: "Member",
      rules: [rules.required()],
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
