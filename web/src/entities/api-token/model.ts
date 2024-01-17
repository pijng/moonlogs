import { ApiToken, getApiToken, getApiTokens } from "@/shared/api/api-tokens";
import { createEffect, createStore } from "effector";

const getApiTokensFx = createEffect(() => {
  return getApiTokens();
});

const getApiTokenFx = createEffect((id: number) => {
  return getApiToken(id);
});

export const $apiTokens = createStore<ApiToken[]>([]).on(
  getApiTokensFx.doneData,
  (_, apiTokensResponse) => apiTokensResponse.data,
);

export const $currentApiToken = createStore<ApiToken>({ id: 0, name: "", token: "", is_revoked: false }).on(
  getApiTokenFx.doneData,
  (_, apiTokenResponse) => apiTokenResponse.data,
);

export const effects = {
  getApiTokensFx,
  getApiTokenFx,
};

export const events = {};
