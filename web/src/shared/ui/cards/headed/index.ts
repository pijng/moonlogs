import { Store, combine, createStore } from "effector";
import { h, list, spec } from "forest";
import { Button } from "@/shared/ui";
import { Schema } from "@/shared/api";

export const CardHeaded = ({
  tags,
  kind,
  schema,
  content,
  href,
  withMore,
}: {
  tags: Store<Array<[string, any]>>;
  kind: Store<string | null>;
  schema: Store<Schema | null>;
  content: () => void;
  href?: Store<string>;
  withMore?: boolean;
}) => {
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

      const localSchema = schema || createStore({});
      const localKind = kind || createStore(null);

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

        h("li", () => {
          const $kindTitle = combine(localKind, localSchema, (kind, schema) => {
            const schemaKind = schema?.kinds?.find((k) => k.name === kind);
            if (!schemaKind) return null;

            return schemaKind.title;
          });

          spec({ visible: $kindTitle.map(Boolean) });

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
            text: $kindTitle,
          });
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
              text: combine(tag, localSchema, (tag, schema) => {
                const [tagKey, tagValue] = tag;
                const field = schema?.fields.find((f) => f.name === tagKey);
                if (!field) return `${tagKey}: ${tagValue}`;

                return `${field.title}: ${tagValue}`;
              }),
            });
          });
        });
      });

      h("div", () => {
        spec({
          visible: createStore(Boolean(href)),
          classList: [
            "basis-1/12",
            "p-4",
            "min-w-fit",
            "border-l",
            "border-gray-200",
            "dark:border-gray-700",
            "flex",
            "justify-center",
          ],
        });

        h("a", () => {
          spec({
            attr: { href: href || "", target: "_blank" },
          });
          Button({ text: "Open", variant: "default", size: "extra_small" });
        });
      });
    });

    h("div", () => {
      spec({
        attr: { id: "defaultTabContent" },
      });

      content();
    });

    h("div", () => {
      spec({
        visible: createStore(withMore ?? false),
        classList: ["text-center", "cursor-default", "pb-1", "dark:text-gray-400"],
        text: "...",
      });
    });
  });
};
