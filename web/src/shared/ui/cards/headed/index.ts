import { Store, combine, createStore } from "effector";
import { h, list, spec } from "forest";
import { Button, KBD } from "@/shared/ui";
import { Schema } from "@/shared/api";
import { i18n } from "@/shared/lib/i18n";

export const CardHeaded = ({
  tags,
  kind,
  schema,
  content,
  href,
}: {
  tags: Store<Array<[string, any]>>;
  kind: Store<string | null>;
  schema: Store<Schema | null>;
  content: () => void;
  href?: Store<string>;
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
        "dark:bg-raisin-black",
        "dark:border-shadow-gray",
      ],
    });

    h("div", () => {
      spec({
        classList: [
          "flex",
          "flex-row",
          "items-center",
          "bg-gray-50",
          "dark:bg-raisin-black",
          "border-b",
          "border-gray-200",
          "rounded-t-lg",
          "dark:border-shadow-gray",
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
            "overflow-auto",
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

          KBD({ text: $kindTitle });
        });

        list(tags, ({ store: tag }) => {
          h("li", () => {
            spec({
              classList: ["min-w-fit"],
            });

            const text = combine(tag, localSchema, (tag, schema) => {
              const [tagKey, tagValue] = tag;
              const field = schema?.fields.find((f) => f.name === tagKey);
              if (!field) return `${tagKey}: ${tagValue}`;

              return `${field.title}: ${tagValue}`;
            });

            KBD({ text: text });
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
            "dark:border-shadow-gray",
            "flex",
            "justify-center",
          ],
        });

        h("a", () => {
          spec({
            attr: { href: href || "", target: "_blank" },
          });
          Button({ text: i18n("buttons.open"), variant: "default", size: "extra_small" });
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
