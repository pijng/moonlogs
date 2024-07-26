import { invokeTemplate } from "./templating";
import { locales } from "./translations";
import { Store } from "effector";
import { $preferredLanguage } from "./locale";
import { Join, PathsToStringProps, Translation } from "./translations/types";

export { setLanguage, $preferredLanguage } from "./locale";

type TranslationPath = Join<PathsToStringProps<Translation>>;

export const i18n = (path: TranslationPath, vars?: Record<string, any>): Store<string> => {
  const $locale = $preferredLanguage.map((lang) => locales[lang]);

  const $template = $locale.map((locale) => {
    return path.split(".").reduce((prev: any, cur: string) => {
      return prev[cur];
    }, locale) as string;
  });

  return invokeTemplate($template, vars);
};
