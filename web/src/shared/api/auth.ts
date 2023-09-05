import { get, post } from "./base";

export type SessionReponse = {
  data: { token: string };
  meta: {
    page: number;
    count: number;
    per_page: number;
  };
};

export const postSession = (email: string, password: string): Promise<SessionReponse> => {
  return post({ url: "/api/session", body: JSON.stringify({ email, password }) });
};

export const getSession = (): Promise<SessionReponse> => {
  return get({ url: "/api/session" });
};
