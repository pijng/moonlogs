import { BaseResponse, del, get, post, put } from "./base";

export type UserRole = "Member" | "Admin";

export type User = {
  id: number;
  name: string;
  email: string;
  role: UserRole;
  tag_ids: number[];
  token: string;
  is_revoked: boolean;
  insights_enabled?: boolean;
};

export type UserToCreate = {
  name: string;
  email: string;
  role: UserRole;
  tag_ids: number[];
  password: string;
  passwordConfirmation: string;
};

export type UserToUpdate = {
  id: number;
  name: string;
  email: string;
  is_revoked: boolean;
  role: UserRole;
  tag_ids: number[];
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

export const deleteUser = (id: number): Promise<UserReponse> => {
  return del({ url: `/api/users/${id}` });
};
