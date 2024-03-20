import { User, getUsers, getUser } from "@/shared/api";
import { setCurrentAccount } from "@/shared/auth";
import { setLanguage } from "@/shared/lib/i18n";
import { createEffect, createEvent, createStore, sample } from "effector";

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

sample({
  source: $currentAccount,
  target: setCurrentAccount,
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
  loadLocaleFromStorageFx,
};

export const events = {};
