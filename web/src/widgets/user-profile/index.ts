import { tagModel } from "@/entities/tag";
import { userModel } from "@/entities/user";
import { $currentTheme, themeChanged } from "@/shared/lib";
import { i18n } from "@/shared/lib/i18n";
import { Select } from "@/shared/ui";
import { combine, createStore } from "effector";
import { h, remap, spec } from "forest";

export const UserProfile = () => {
  h("div", () => {
    spec({
      classList: ["container", "mx-auto", "mt-3"],
    });

    h("div", () => {
      spec({
        classList: ["mt-5"],
      });

      h("div", () => {
        spec({ classList: ["pb-3"] });

        h("p", { text: i18n("profile.name") });

        h("p", {
          classList: ["truncate", "tracking-tight", "font-semibold", "text-gray-900", "dark:text-white"],
          text: remap(userModel.$currentAccount, "name"),
        });
      });

      h("div", () => {
        spec({ classList: ["py-3"] });

        h("p", { text: i18n("profile.email") });

        h("p", {
          classList: ["truncate", "tracking-tight", "font-semibold", "text-gray-900", "dark:text-white"],
          text: remap(userModel.$currentAccount, "email"),
        });
      });

      h("div", () => {
        spec({ classList: ["py-3"] });

        h("p", { text: i18n("profile.role") });

        h("p", {
          classList: ["truncate", "tracking-tight", "font-semibold", "text-gray-900", "dark:text-white"],
          text: remap(userModel.$currentAccount, "role"),
        });
      });

      h("div", () => {
        spec({ classList: ["py-3"] });

        h("p", { text: i18n("profile.tags") });

        const $userTagNames = combine([tagModel.$tags, remap(userModel.$currentAccount, "tag_ids")], ([tags, tagIds]) => {
          const appliedTags = tags
            .filter((t) => tagIds?.includes(t.id))
            .map((t) => t.name)
            .join(", ");

          return appliedTags || "—";
        });

        h("p", {
          classList: ["truncate", "tracking-tight", "font-semibold", "text-gray-900", "dark:text-white"],
          text: $userTagNames,
        });
      });

      h("div", () => {
        spec({ classList: ["py-3", "max-w-5"] });

        h("p", { text: i18n("profile.language") });

        Select({
          value: userModel.$currentLocale,
          options: createStore(["en", "ru"]),
          optionSelected: userModel.localeChanged,
          withBlank: createStore(false),
        });
      });

      h("div", () => {
        spec({ classList: ["py-3", "max-w-5"] });

        h("p", { text: i18n("profile.theme") });

        Select({
          value: $currentTheme,
          options: createStore(["dark", "light"]),
          optionSelected: themeChanged,
          withBlank: createStore(false),
        });
      });
    });
  });
};
