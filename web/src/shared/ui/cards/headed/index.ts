import { Store, createStore } from "effector";
import { h, list, spec } from "forest";
import { Button, ButtonVariant } from "@/shared/ui";

export const CardHeaded = ({ tags, content, href }: { tags: Store<string[]>; content: () => void; href?: Store<string> }) => {
  h("div", () => {
    spec({
      classList: [
        "w-full",
        "bg-white",
        "border",
        "border-gray-200",
        "rounded-lg",
        "shadow",
        "dark:bg-gray-800",
        "dark:border-gray-700",
      ],
    });

    h("div", () => {
      spec({
        classList: [
          "flex",
          "flex-row",
          "items-center",
          "bg-gray-50",
          "dark:bg-gray-800",
          "border-b",
          "border-gray-200",
          "rounded-t-lg",
          "dark:border-gray-700",
          "dark:text-gray-400",
        ],
      });

      h("ul", () => {
        spec({
          classList: [
            "flex",
            "flex-wrap",
            "gap-3",
            "basis-11/12",
            "flex-nowrap",
            "overflow-scroll",
            "text-sm",
            "p-4",
            "justify-start",
            "font-medium",
            "text-center",
            "text-gray-500",
          ],
          attr: { role: "tablist", id: "defaultTab" },
        });

        list(tags, ({ store: tag }) => {
          h("li", () => {
            spec({
              classList: ["min-w-fit"],
            });

            h("kbd", {
              classList: [
                "block",
                "px-2",
                "py-1.5",
                "text-xs",
                "font-semibold",
                "text-gray-800",
                "bg-gray-100",
                "border",
                "border-gray-200",
                "rounded-lg",
                "dark:bg-gray-600",
                "dark:text-gray-100",
                "dark:border-gray-500",
              ],
              text: tag,
            });
          });
        });
      });

      h("div", () => {
        spec({
          visible: createStore(Boolean(href)),
          classList: ["basis-1/12", "p-4", "min-w-fit", "border-l", "border-gray-200", "dark:border-gray-700"],
        });

        h("a", () => {
          spec({
            attr: { href: href || "" },
          });
          Button({ text: "Open", variant: createStore<ButtonVariant>("default"), size: "extra_small" });
        });
      });
    });

    h("div", () => {
      spec({
        attr: { id: "defaultTabContent" },
      });

      content();
    });
  });
};
