import { User, getUsers } from "@/shared/api";
import { Cell } from "@/shared/ui";
import { createEffect, createStore } from "effector";

const getUsersFx = createEffect(() => {
  return getUsers();
});

export const $users = createStore<User[]>([]).on(getUsersFx.doneData, (_, usersResponse) => usersResponse.data);

export const $formattedUsers = $users.map((users) => {
  return users.map<Cell[]>((user) => {
    return [
      { data: user.email, component: "text" },
      { data: user.name, component: "text" },
      { data: user.role, component: "text" },
    ];
  });
});

export const effects = {
  getUsersFx,
};

export const events = {};
