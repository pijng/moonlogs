import { User, getUsers } from "@/shared/api";
import { getUser } from "@/shared/api/users";
import { setLanguage } from "@/shared/lib/i18n";
import { createEffect, createEvent, createStore, sample } from "effector";

const THEME_KEY = "theme";
const LOCALE_KEY = "locale";

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
  return localStorage.getItem(THEME_KEY) || "light";
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

const loadLocaleFromStorageFx = createEffect(() => {
  return localStorage.getItem(LOCALE_KEY) || "en";
});

const setLocaleToStorageFx = createEffect((locale: string) => {
  return localStorage.setItem(LOCALE_KEY, locale);
});

export const $currentLocale = createStore<string>("en");

// This event for ui (click in profile)
export const localeChanged = createEvent<string>();

sample({
  source: localeChanged,
  fn: (locale) => locale.toLowerCase(),
  target: [setLocaleToStorageFx, $currentLocale],
});

sample({
  source: loadLocaleFromStorageFx.doneData,
  target: $currentLocale,
});

sample({
  source: $currentLocale,
  target: setLanguage,
});

export const effects = {
  getUsersFx,
  getUserFx,
  loadThemeFromStorageFx,
  loadLocaleFromStorageFx,
};

export const events = {};
