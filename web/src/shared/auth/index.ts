import { createEffect, createEvent, createStore, sample } from "effector";
import { User, getSession } from "@/shared/api";

export const tokenReceived = createEvent<string>();
export const tokenErased = createEvent();

export const $token = createStore("");

export const $isAuthorized = $token.map(Boolean);

export const unauthorizedTriggered = createEvent();
export const notAllowedTriggered = createEvent();
export const notFoundTriggered = createEvent();

$token.on(tokenReceived, (_, token) => token).reset(tokenErased);

export const getSessionFx = createEffect(() => {
  return getSession();
});

export const obtainSession = createEvent();

sample({
  source: getSessionFx.pending,
  clock: obtainSession,
  filter: (pending) => !pending,
  target: getSessionFx,
});

sample({
  source: $isAuthorized,
  clock: getSessionFx.doneData,
  filter: (isAuthorized, sessionResponse) => !isAuthorized && !!sessionResponse?.data?.token && sessionResponse.success,
  fn: (_, sessionResponse) => sessionResponse.data.token,
  target: tokenReceived,
});

export const createInitialAdmin = createEvent();

sample({
  clock: getSessionFx.doneData,
  filter: (sessionResponse) => sessionResponse.success && sessionResponse.data.should_create_initial_admin,
  target: createInitialAdmin,
});

export const setCurrentAccount = createEvent<User>();
export const $currentAccount = createStore<User>({
  id: 0,
  name: "",
  email: "",
  role: "Member",
  tag_ids: [],
  token: "",
  is_revoked: false,
});

sample({
  source: setCurrentAccount,
  target: $currentAccount,
});
