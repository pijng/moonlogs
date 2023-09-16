import { BaseResponse, get, post } from "./base";

export interface SessionResponse extends BaseResponse {
  data: { token: string };
}

export const postSession = (email: string, password: string): Promise<SessionResponse> => {
  return post({ url: "/api/session", body: JSON.stringify({ email, password }) });
};

export const getSession = (): Promise<SessionResponse> => {
  return get({ url: "/api/session" });
};
