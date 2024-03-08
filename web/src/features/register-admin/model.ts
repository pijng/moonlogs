import { loginRoute } from "@/shared/routing";
import { registerAdmin } from "@/shared/api";
import { rules } from "@/shared/lib";
import { createEffect, createEvent, createStore, sample } from "effector";
import { createForm } from "effector-forms";

export const registerAdminForm = createForm<{ name: string; email: string; password: string; passwordConfirmation: string }>({
  fields: {
    name: {
      init: "",
      rules: [rules.required()],
    },
    email: {
      init: "",
      rules: [rules.required(), rules.email()],
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

export const $registrationError = createStore("");

export const registerAdminFx = createEffect((adminInfo: { name: string; email: string; password: string }) => {
  return registerAdmin(adminInfo);
});

sample({
  source: registerAdminForm.formValidated,
  target: registerAdminFx,
});

export const registerSubmitted = createEvent();

sample({
  source: registerAdminFx.doneData,
  fn: (registrationResponse) => registrationResponse.data.token,
  target: [loginRoute.open, registerAdminForm.reset],
});

sample({
  source: registerAdminFx.doneData,
  filter: (registrationResponse) => !registrationResponse.success,
  fn: (registrationResponse) => registrationResponse.error,
  target: $registrationError,
});
