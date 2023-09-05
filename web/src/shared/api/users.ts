import { get } from "./base";

export type UserRole = "Member" | "Admin" | "System";

export type User = {
  id: number;
  name: string;
  email: string;
  role: UserRole;
  token: string;
};

export type UsersReponse = {
  data: User[];
  meta: {
    page: number;
    count: number;
    per_page: number;
  };
};

export const getUsers = (): Promise<UsersReponse> => {
  return get({ url: "/api/users" });
};
