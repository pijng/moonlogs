import { createEffect, createEvent, createStore, sample } from "effector";
import { getSession } from "@/shared/api";

export const tokenReceived = createEvent<string>();
export const tokenErased = createEvent();

export const $token = createStore("");

export const $isAuthorized = $token.map(Boolean);

export const unauthorizedTriggered = createEvent();

$token.on(tokenReceived, (_, token) => token).reset(tokenErased);

export const getSessionFx = createEffect(() => {
  return getSession();
});

sample({
  source: $isAuthorized,
  clock: getSessionFx.doneData,
  filter: (isAuthorized, sessionResponse) => !isAuthorized && !!sessionResponse?.data?.token,
  fn: (_, sessionResponse) => sessionResponse.data.token,
  target: tokenReceived,
});