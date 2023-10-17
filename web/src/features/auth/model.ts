import { homeRoute } from "@/routing/shared";
import { postSession } from "@/shared/api";
import { tokenReceived } from "@/shared/auth";
import { rules } from "@/shared/lib";
import { createEffect, createEvent, createStore, sample } from "effector";
import { createForm } from "effector-forms";

export const authForm = createForm<{ email: string; password: string }>({
  fields: {
    email: {
      init: "",
      rules: [rules.required(), rules.email()],
    },
    password: {
      init: "",
      rules: [rules.required()],
    },
  },
  validateOn: ["submit"],
});

export const $authError = createStore("");

export const logInFx = createEffect(({ email, password }: { email: string; password: string }) => {
  return postSession(email, password);
});

sample({
  source: authForm.formValidated,
  target: logInFx,
});

export const logInSubmitted = createEvent();

sample({
  source: logInFx.doneData,
  fn: (logInResponse) => logInResponse.data.token,
  target: [tokenReceived, homeRoute.open],
});

sample({
  source: logInFx.doneData,
  filter: (logInResponse) => !logInResponse.success,
  fn: (logInResponse) => {
    if (logInResponse.code === 401) {
      return "Wrong email or password";
    }

    return logInResponse.error;
  },
  target: $authError,
});
