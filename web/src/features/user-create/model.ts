import { membersRoute } from "@/routing/shared";
import { UserToCreate, createUser } from "@/shared/api";
import { rules } from "@/shared/lib";
import { layoutClicked } from "@/shared/ui";
import { createEffect, createEvent, createStore, restore, sample } from "effector";
import { createForm } from "effector-forms";

const tagChecked = createEvent<number>();
const tagUnchecked = createEvent<number>();
const tagSelectionClicked = createEvent<any>();

export const memberForm = createForm<UserToCreate>({
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
      rules: [rules.required(), rules.password()],
    },
    passwordConfirmation: {
      init: "",
      rules: [rules.required(), rules.password()],
    },
  },
  validateOn: ["submit"],
});

export const $creationError = createStore("");

export const createUserFx = createEffect((user: UserToCreate) => {
  return createUser(user);
});

sample({
  source: memberForm.formValidated,
  target: createUserFx,
});

sample({
  source: createUserFx.doneData,
  filter: (userResponse) => userResponse.success && Boolean(userResponse.data.id),
  target: [membersRoute.open, memberForm.reset],
});

sample({
  source: createUserFx.doneData,
  filter: (userResponse) => !userResponse.success,
  fn: (userResponse) => userResponse.error,
  target: $creationError,
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
