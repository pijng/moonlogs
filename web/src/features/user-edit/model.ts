import { userModel } from "@/entities/user";
import { membersRoute } from "@/routing/shared";
import { UserToUpdate, editUser } from "@/shared/api";
import { deleteUser } from "@/shared/api/users";
import { rules } from "@/shared/lib";
import { layoutClicked } from "@/shared/ui";
import { redirect } from "atomic-router";
import { createEffect, createEvent, createStore, restore, sample } from "effector";
import { createForm } from "effector-forms";

const tagChecked = createEvent<number>();
const tagUnchecked = createEvent<number>();
const tagSelectionClicked = createEvent<any>();

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

const deleteUserFx = createEffect((id: number) => {
  deleteUser(id);
});

export const deleteUserClicked = createEvent<number>();
const alertDeleteFx = createEffect((id: number): { confirmed: boolean; id: number } => {
  const confirmed = confirm("Are you sure you want to delete this user?");

  return confirmed ? { confirmed: true, id: id } : { confirmed: false, id: id };
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
  tagSelectionClicked,
  tagChecked,
  tagUnchecked,
};

export const $tagsDropwdownIsOpened = createStore(false);

sample({
  source: $tagsDropwdownIsOpened,
  clock: tagSelectionClicked,
  fn: (state) => !state,
  target: $tagsDropwdownIsOpened,
});

sample({
  source: [$tagsDropwdownIsOpened, restore(tagSelectionClicked, null)],
  clock: layoutClicked,
  filter: ([isOpened, clicked], layoutClicked) => {
    const path = layoutClicked.composedPath();

    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore
    return !path.includes(clicked?.target?.parentNode) && isOpened;
  },
  target: tagSelectionClicked,
});
