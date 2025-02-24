import { Store, combine, createStore } from "effector";
import { h, list, remap, spec } from "forest";
import { Button, KBD } from "@/shared/ui";
import { LogsGroup, Schema } from "@/shared/api";
import { i18n } from "@/shared/lib/i18n";
import { logModel } from "@/entities/log";
import { GroupActionsList } from "../group-actions-list";
import { LogsTable } from "@/features/logs-table";

export const LogsCard = ({
  schema,
  logsGroup,
  href,
}: {
  schema: Store<Schema | null>;
  logsGroup: Store<LogsGroup>;
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
      const localKind = remap(logsGroup, "kind");

      h("ul", () => {
        spec({
          classList: [
            "flex",
            "flex-wrap",
            "gap-3",
            "basis-full",
            "flex-nowrap",
            "overflow-auto",
            "dark:scrollbar",
            "text-sm",
            "px-4",
            "py-2",
            "justify-start",
            "font-medium",
            "text-center",
            "text-gray-500",
          ],
          attr: { role: "tablist", id: "defaultTab" },
        });

        h("li", () => {
          spec({
            classList: ["min-w-max"],
          });

          const $kindTitle = combine(localKind, localSchema, (kind, schema) => {
            const schemaKind = schema?.kinds?.find((k) => k.name === kind);
            if (!schemaKind) return null;

            return schemaKind.title;
          });

          spec({ visible: $kindTitle.map(Boolean) });

          KBD({ text: $kindTitle });
        });

        list(remap(logsGroup, "tags"), ({ store: tag }) => {
          h("li", () => {
            spec({
              classList: ["min-w-max"],
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
          classList: ["p-2", "min-w-fit", "border-l", "border-gray-200", "dark:border-shadow-gray", "inline-flex", "space-x-2"],
        });

        h("div", () => {
          spec({
            classList: ["lg:hidden"],
          });

          GroupActionsList({ logGroup: remap(logsGroup, "logs"), style: "alternative", collapsed: true });
        });

        h("div", () => {
          spec({
            classList: ["space-x-2", "hidden", "lg:inline-flex"],
          });

          GroupActionsList({ logGroup: remap(logsGroup, "logs"), style: "alternative" });
        });

        h("a", () => {
          spec({
            attr: { href: href || "", target: "_blank" },
            visible: createStore(Boolean(href)),
          });
          Button({ text: i18n("buttons.open"), variant: "default", size: "extra_small" });
        });
      });
    });

    h("div", () => {
      spec({
        attr: { id: "defaultTabContent" },
      });

      LogsTable({
        logs: remap(logsGroup, "logs"),
        requestClicked: logModel.events.requestURLClicked,
        responseClicked: logModel.events.responseURLClicked,
      });
    });
  });
};
