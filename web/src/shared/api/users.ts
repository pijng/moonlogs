import { BaseResponse, get, post, put } from "./base";

export type UserRole = "Member" | "Admin" | "System";

export type User = {
  id: number;
  name: string;
  email: string;
  role: UserRole;
  token: string;
};

export type UserToCreate = {
  name: string;
  email: string;
  role: UserRole;
  password: string;
  passwordConfirmation: string;
};

export type UserToUpdate = {
  id: number;
  name: string;
  email: string;
  role: UserRole;
  password: string;
  passwordConfirmation: string;
};

export interface UserReponse extends BaseResponse {
  data: User;
}

export interface UsersListReponse extends BaseResponse {
  data: User[];
}

export const getUsers = (): Promise<UsersListReponse> => {
  return get({ url: "/api/users" });
};

export const getUser = (id: number): Promise<UserReponse> => {
  return get({ url: `/api/users/${id}` });
};

export const createUser = (user: UserToCreate): Promise<UserReponse> => {
  return post({ url: "/api/users", body: JSON.stringify(user) });
};

export const editUser = (user: UserToUpdate): Promise<UserReponse> => {
  return put({ url: `/api/users/${user.id}`, body: JSON.stringify(user) });
};
