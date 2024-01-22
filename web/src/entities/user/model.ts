import { User, getUsers } from "@/shared/api";
import { getUser } from "@/shared/api/users";
import { createEffect, createStore } from "effector";

const getUsersFx = createEffect(() => {
  return getUsers();
});

const getUserFx = createEffect((id: number) => {
  return getUser(id);
});

export const $users = createStore<User[]>([]).on(getUsersFx.doneData, (_, usersResponse) => usersResponse.data);

export const $currentUser = createStore<User>({ id: 0, name: "", email: "", role: "Member", tag_ids: [], token: "" }).on(
  getUserFx.doneData,
  (_, userResponse) => userResponse.data,
);

export const $currentAccount = createStore<User>({ id: 0, name: "", email: "", role: "Member", tag_ids: [], token: "" });

export const effects = {
  getUsersFx,
  getUserFx,
};

export const events = {};
