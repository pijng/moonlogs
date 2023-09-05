import { h } from "forest";

type Size = "xl" | "2xl" | "3xl";

export const Logo = (size?: Size) => {
  const localSize = size ?? "xl";

  h("div", () => {
    h("span", {
      classList: ["mr-3", "leading-7", `text-${localSize}`],
      attr: { alt: "Moonlogs logo" },
      text: "🌘",
    });

    h("span", {
      classList: ["self-center", `text-${localSize}`, "font-semibold", "whitespace-nowrap", "dark:text-white"],
      text: "Moonlogs",
    });
  });
};
