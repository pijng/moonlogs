import { User, getUsers } from "@/shared/api";
import { getUser } from "@/shared/api/users";
import { createEffect, createEvent, createStore, sample } from "effector";

const THEME_KEY = "theme";

const getUsersFx = createEffect(() => {
  return getUsers();
});

const getUserFx = createEffect((id: number) => {
  return getUser(id);
});

export const $users = createStore<User[]>([]).on(getUsersFx.doneData, (_, usersResponse) => usersResponse.data);

export const $currentUser = createStore<User>({
  id: 0,
  name: "",
  email: "",
  role: "Member",
  tag_ids: [],
  token: "",
  is_revoked: false,
}).on(getUserFx.doneData, (_, userResponse) => userResponse.data);

export const $currentAccount = createStore<User>({
  id: 0,
  name: "",
  email: "",
  role: "Member",
  tag_ids: [],
  token: "",
  is_revoked: false,
});

const loadThemeFromStorageFx = createEffect(() => {
  return localStorage.getItem(THEME_KEY);
});

const setThemeToStorageFx = createEffect((theme: string) => {
  return localStorage.setItem(THEME_KEY, theme);
});

export const $currentTheme = createStore<string | null>("light");
export const themeChanged = createEvent<string>();

sample({
  source: themeChanged,
  fn: (theme) => theme.toLowerCase(),
  target: [setThemeToStorageFx, $currentTheme],
});

sample({
  source: loadThemeFromStorageFx.doneData,
  target: $currentTheme,
});

export const effects = {
  getUsersFx,
  getUserFx,
  loadThemeFromStorageFx,
};

export const events = {};
