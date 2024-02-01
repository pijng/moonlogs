import { tagModel } from "@/entities/tag";
import { userModel } from "@/entities/user";
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

        h("p", { text: "Name" });

        h("p", {
          classList: ["truncate", "tracking-tight", "font-semibold", "text-gray-900", "dark:text-white"],
          text: remap(userModel.$currentAccount, "name"),
        });
      });

      h("div", () => {
        spec({ classList: ["py-3"] });

        h("p", { text: "Email" });

        h("p", {
          classList: ["truncate", "tracking-tight", "font-semibold", "text-gray-900", "dark:text-white"],
          text: remap(userModel.$currentAccount, "email"),
        });
      });

      h("div", () => {
        spec({ classList: ["py-3"] });

        h("p", { text: "Role" });

        h("p", {
          classList: ["truncate", "tracking-tight", "font-semibold", "text-gray-900", "dark:text-white"],
          text: remap(userModel.$currentAccount, "role"),
        });
      });

      h("div", () => {
        spec({ classList: ["py-3"] });

        h("p", { text: "Tags" });

        const $userTagNames = combine([tagModel.$tags, remap(userModel.$currentAccount, "tag_ids")], ([tags, tagIds]) => {
          const appliedTags = tags
            .filter((t) => tagIds.includes(t.id))
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

        h("p", { text: "Theme" });

        Select({
          value: userModel.$currentTheme,
          text: "",
          options: createStore(["dark", "light"]),
          optionSelected: userModel.themeChanged,
          withBlank: createStore(false),
        });
      });
    });
  });
};
