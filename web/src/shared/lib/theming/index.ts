import { createEffect, createEvent, createStore, sample } from "effector";

const THEME_KEY = "theme";

type Theme = "dark" | "light";

export const $currentTheme = createStore<Theme>("light");
export const themeChanged = createEvent<Theme>();

export const loadThemeFromStorageFx = createEffect(() => {
  const theme: Theme = (localStorage.getItem(THEME_KEY) || "light") as Theme;

  return theme;
});

const setThemeToStorageFx = createEffect((theme: Theme) => {
  return localStorage.setItem(THEME_KEY, theme);
});

const applyThemeFx = createEffect((theme: Theme) => {
  if (theme === "dark") {
    return document.querySelector("html")?.classList?.add("dark");
  }

  return document.querySelector("html")?.classList?.remove("dark");
});

sample({
  clock: [themeChanged, loadThemeFromStorageFx.doneData],
  target: [applyThemeFx, setThemeToStorageFx, $currentTheme],
});
