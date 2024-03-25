import { i18n } from "@/shared/lib/i18n";
import { h } from "forest";
import { LogoIcon } from "@/shared/ui";

type Size = "xl" | "2xl" | "3xl";

export const Logo = (size?: Size) => {
  const localSize = size ?? "xl";

  h("div", () => {
    LogoIcon();

    h("span", {
      classList: ["self-center", `text-${localSize}`, "font-semibold", "whitespace-nowrap", "dark:text-white"],
      text: i18n("miscellaneous.brand"),
    });
  });
};
