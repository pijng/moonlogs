import { apiTokenModel } from "@/entities/api-token";
import { apiTokensRoute } from "@/shared/routing";
import { ApiTokenToUpdate, deleteApiToken, editApiToken } from "@/shared/api";
import { rules } from "@/shared/lib";
import { i18n } from "@/shared/lib/i18n";
import { redirect } from "atomic-router";
import { attach, createEffect, createEvent, createStore, sample } from "effector";
import { createForm } from "effector-forms";

export const apiTokenForm = createForm<Omit<ApiTokenToUpdate, "id">>({
  fields: {
    name: {
      init: "",
      rules: [rules.required()],
    },
    is_revoked: {
      init: false,
      rules: [],
    },
  },
  validateOn: ["submit"],
});

export const $editError = createStore("");

export const editApiTokenFx = createEffect((apiToken: ApiTokenToUpdate) => {
  return editApiToken(apiToken);
});

sample({
  source: apiTokenModel.$currentApiToken,
  target: apiTokenForm.setForm,
});

sample({
  source: apiTokenModel.$currentApiToken,
  clock: apiTokenForm.formValidated,
  fn: (currentApiToken, apiTokenToEdit) => {
    return { ...apiTokenToEdit, id: currentApiToken.id };
  },
  target: editApiTokenFx,
});

sample({
  source: editApiTokenFx.doneData,
  filter: (apiTokenResponse) => apiTokenResponse.success && Boolean(apiTokenResponse.data.id),
  target: [apiTokenForm.reset, apiTokensRoute.open],
});

sample({
  source: editApiTokenFx.doneData,
  filter: (apiTokenResponse) => !apiTokenResponse.success,
  fn: (apiTokenResponse) => apiTokenResponse.error,
  target: $editError,
});

const deleteApiTokenFx = createEffect((id: number) => {
  deleteApiToken(id);
});

export const deleteApiTokenClicked = createEvent<number>();
const alertDeleteFx = attach({
  source: i18n("api_tokens.alerts.delete"),
  effect(alertText, id: number) {
    const confirmed = confirm(alertText);

    return confirmed ? { confirmed: true, id: id } : { confirmed: false, id: id };
  },
});

sample({
  clock: deleteApiTokenClicked,
  target: alertDeleteFx,
});

sample({
  source: apiTokenModel.$currentApiToken,
  clock: alertDeleteFx.doneData,
  filter: (_, { confirmed }) => confirmed,
  fn: ({ id }) => id,
  target: deleteApiTokenFx,
});

redirect({
  clock: deleteApiTokenFx.done,
  route: apiTokensRoute,
});
