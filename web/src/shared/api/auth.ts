import { BaseResponse, get, post } from "./base";
import { UserRole } from "./users";

export interface SessionResponse extends BaseResponse {
  data: {
    token: string;
    should_create_initial_admin: boolean;
    id: number;
    name: string;
    email: string;
    role: UserRole;
    tag_ids: number[];
    is_revoked: boolean;
    insights_enabled: boolean;
  };
}

export const postSession = (email: string, password: string): Promise<SessionResponse> => {
  return post({ url: "/api/session", body: JSON.stringify({ email, password }) });
};

export const getSession = (): Promise<SessionResponse> => {
  return get({ url: "/api/session" });
};
