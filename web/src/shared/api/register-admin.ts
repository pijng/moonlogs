import { SessionResponse } from "./auth";
import { post } from "./base";

export const registerAdmin = (admin: { name: string; email: string; password: string }): Promise<SessionResponse> => {
  return post({ url: "/api/setup/register_admin", body: JSON.stringify(admin) });
};
