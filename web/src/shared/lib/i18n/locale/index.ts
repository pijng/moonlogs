import { createEvent, createStore, sample } from "effector";

export const setLanguage = createEvent<string>();
export const $preferredLanguage = createStore<string>("en");

sample({
  source: setLanguage,
  target: $preferredLanguage,
});
