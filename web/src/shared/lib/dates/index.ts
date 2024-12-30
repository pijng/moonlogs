import { createEvent, createStore, sample } from "effector";

export const DATEFORMAT_OPTIONS: Intl.DateTimeFormatOptions = {
  hour: "numeric",
  minute: "numeric",
  second: "numeric",
  fractionalSecondDigits: 3,
  year: "numeric",
  month: "numeric",
  day: "numeric",
};

export const TIMEZONE = Intl.DateTimeFormat().resolvedOptions().timeZone;

export const setDateLanguage = createEvent<string>();

export const $intl = createStore<Intl.DateTimeFormat>(Intl.DateTimeFormat());

sample({
  source: setDateLanguage,
  fn: (lang) => {
    return Intl.DateTimeFormat(lang, DATEFORMAT_OPTIONS);
  },
  target: $intl,
});
