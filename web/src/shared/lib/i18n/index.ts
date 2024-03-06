import { TranslationPath } from "@/shared/types";
import { invokeTemplate } from "./templating";
import { locales } from "./translations";
import { Store } from "effector";
import { $preferredLanguage } from "./locale";
export { setLanguage } from "./locale";

export const i18n = (path: TranslationPath, vars?: Record<string, any>): Store<string> => {
  const $locale = $preferredLanguage.map((lang) => locales[lang]);

  const $template = $locale.map((locale) => {
    return path.split(".").reduce((prev: any, cur: string) => {
      return prev[cur];
    }, locale) as string;
  });

  return invokeTemplate($template, vars);
};
