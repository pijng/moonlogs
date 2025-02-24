import { router } from "@/shared/routing";
import { ApiTokenToCreate, createApiToken } from "@/shared/api";
import { rules } from "@/shared/lib";
import { createEffect, createEvent, createStore, sample } from "effector";
import { createForm } from "effector-forms";

export const apiTokenForm = createForm<ApiTokenToCreate>({
  fields: {
    name: {
      init: "",
      rules: [rules.required()],
    },
  },
  validateOn: ["submit"],
});

export const $creationError = createStore("");

export const createApiTokenFx = createEffect((apiToken: ApiTokenToCreate) => {
  return createApiToken(apiToken);
});

const resetToken = createEvent();
export const $freshToken = createStore("").reset(resetToken);

sample({
  source: apiTokenForm.formValidated,
  target: createApiTokenFx,
});

sample({
  source: createApiTokenFx.doneData,
  filter: (apiTokenResponse) => apiTokenResponse.success && Boolean(apiTokenResponse.data.id),
  fn: (apiTokenResponse) => apiTokenResponse.data.token,
  target: [$freshToken, apiTokenForm.reset],
});

sample({
  clock: router.$path,
  target: resetToken,
});

sample({
  source: createApiTokenFx.doneData,
  filter: (apiTokenResponse) => !apiTokenResponse.success,
  fn: (apiTokenResponse) => apiTokenResponse.error,
  target: $creationError,
});
