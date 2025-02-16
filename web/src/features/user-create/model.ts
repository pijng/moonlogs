import { membersRoute } from "@/shared/routing";
import { UserToCreate, createUser } from "@/shared/api";
import { bindFieldList, manageSubmit, rules } from "@/shared/lib";
import { createEffect, createEvent, createStore } from "effector";
import { createForm } from "effector-forms";

const tagChecked = createEvent<number>();
const tagUnchecked = createEvent<number>();

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

bindFieldList({ field: memberForm.fields.tag_ids, added: tagChecked, removed: tagUnchecked });

manageSubmit({
  form: memberForm,
  actionFx: createUserFx,
  error: $creationError,
  route: membersRoute,
});

export const events = {
  tagChecked,
  tagUnchecked,
};
