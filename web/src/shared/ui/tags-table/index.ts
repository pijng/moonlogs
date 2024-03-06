import { h, list, remap, spec } from "forest";
import { Event, Store, createEvent, sample } from "effector";
import { Tag } from "@/shared/api";
import { Button } from "@/shared/ui";
import { i18n } from "@/shared/lib/i18n";

export const TagsTable = (tags: Store<Tag[]>, editClicked: Event<number>) => {
  h("div", () => {
    spec({
      classList: ["antialiased"],
    });

    h("div", () => {
      spec({
        classList: ["relative", "overflow-x-auto", "sm:rounded-b-lg"],
      });

      h("table", () => {
        spec({
          classList: ["min-w-full", "w-max", "text-sm", "text-left", "table-fixed"],
        });

        h("thead", () => {
          spec({
            classList: [
              "text-xs",
              "text-gray-700",
              "uppercase",
              "bg-gray-50",
              "w-full",
              "dark:bg-gray-700",
              "dark:text-gray-200",
            ],
          });

          h("tr", () => {
            h("th", {
              attr: { scope: "col" },
              classList: ["px-6", "py-3"],
              text: i18n("tables.tags.name"),
            });

            h("th", {
              attr: { scope: "col" },
              classList: ["px-6", "py-3", "max-w-28", "w-28"],
              text: i18n("tables.tags.actions"),
            });
          });
        });

        h("tbody", () => {
          list(tags, ({ store: tag }) => {
            h("tr", () => {
              spec({
                classList: [
                  "border-t",
                  "w-full",
                  "dark:border-gray-700",
                  "hover:bg-gray-50",
                  "dark:hover:bg-gray-600",
                  "text-gray-900",
                  "dark:text-gray-200",
                ],
              });

              h("td", () => {
                spec({
                  classList: ["px-6", "py-4"],
                });
                h("div", {
                  text: remap(tag, "name"),
                });
              });

              h("td", () => {
                spec({
                  classList: ["px-6", "py-4"],
                });

                const localEditClicked = createEvent();
                sample({
                  source: remap(tag, "id"),
                  clock: localEditClicked,
                  target: editClicked,
                });

                Button({
                  text: i18n("buttons.edit"),
                  event: localEditClicked,
                  variant: "default",
                  size: "extra_small",
                });
              });
            });
          });
        });
      });
    });
  });
};
