import { BaseResponse, del, get, post, put } from "./base";

export type ApiToken = {
  id: number;
  name: string;
  token: string;
  is_revoked: boolean;
};

export type ApiTokenToCreate = {
  name: string;
};

export type ApiTokenToUpdate = {
  id: number;
  name: string;
  is_revoked: boolean;
};

export interface ApiTokenReponse extends BaseResponse {
  data: ApiToken;
}

export interface ApiTokensListReponse extends BaseResponse {
  data: ApiToken[];
}

export const getApiTokens = (): Promise<ApiTokensListReponse> => {
  return get({ url: "/api/api_tokens" });
};

export const getApiToken = (id: number): Promise<ApiTokenReponse> => {
  return get({ url: `/api/api_tokens/${id}` });
};

export const createApiToken = (apiToken: ApiTokenToCreate): Promise<ApiTokenReponse> => {
  return post({ url: "/api/api_tokens", body: JSON.stringify(apiToken) });
};

export const editApiToken = (apiToken: ApiTokenToUpdate): Promise<ApiTokenReponse> => {
  return put({ url: `/api/api_tokens/${apiToken.id}`, body: JSON.stringify(apiToken) });
};

export const deleteApiToken = (id: number): Promise<ApiTokenReponse> => {
  return del({ url: `/api/api_tokens/${id}` });
};
