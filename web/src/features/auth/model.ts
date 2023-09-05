import { homeRoute } from "@/routing";
import { postSession } from "@/shared/api";
import { tokenReceived } from "@/shared/auth";
import { attach, createEffect, createEvent, restore, sample } from "effector";

export const emailChanged = createEvent<string>();
export const passwordChanged = createEvent<string>();

export const $email = restore(emailChanged, "");
export const $password = restore(passwordChanged, "");

export const logInFx = attach({
  source: { email: $email, password: $password },
  effect: createEffect(({ email, password }: { email: string; password: string }) => {
    return postSession(email, password);
  }),
});

export const logInSubmitted = createEvent();

sample({
  source: logInSubmitted,
  target: logInFx,
});

sample({
  source: logInFx.doneData,
  fn: (logInResponse) => logInResponse.data.token,
  target: [tokenReceived, homeRoute.open],
});
