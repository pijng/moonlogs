import { h, list, remap, spec } from "forest";
import { Event, Store, createEvent, sample } from "effector";
import { Action } from "@/shared/api";
import { Button } from "@/shared/ui";
import { i18n } from "@/shared/lib/i18n";

export const ActionsTable = (actions: Store<Action[]>, editClicked: Event<number>) => {
  h("div", () => {
    spec({
      classList: ["antialiased"],
    });

    h("div", () => {
      spec({
        classList: ["relative", "overflow-x-auto", "dark:scrollbar", "sm:rounded-b-lg"],
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
              "dark:bg-squid-ink",
              "dark:text-gray-200",
            ],
          });

          h("tr", () => {
            h("th", {
              attr: { scope: "col" },
              classList: ["px-6", "py-3"],
              text: i18n("tables.actions.name"),
            });

            h("th", {
              attr: { scope: "col" },
              classList: ["px-6", "py-3"],
              text: i18n("tables.actions.schema_name"),
            });

            h("th", {
              attr: { scope: "col" },
              classList: ["px-6", "py-3"],
              text: i18n("tables.actions.disabled"),
            });

            h("th", {
              attr: { scope: "col" },
              classList: ["px-6", "py-3", "max-w-28", "w-28"],
              text: i18n("tables.actions.actions"),
            });
          });
        });

        h("tbody", () => {
          list(actions, ({ store: action }) => {
            h("tr", () => {
              spec({
                classList: [
                  "border-t",
                  "w-full",
                  "dark:border-shadow-gray",
                  "hover:bg-gray-50",
                  "dark:hover:bg-slate-gray",
                  "text-gray-900",
                  "dark:text-gray-200",
                ],
              });

              h("td", () => {
                spec({
                  classList: ["px-6", "py-4"],
                });
                h("div", {
                  text: remap(action, "name"),
                });
              });

              h("td", () => {
                spec({
                  classList: ["px-6", "py-4"],
                });
                h("div", {
                  text: remap(action, "schema_name"),
                });
              });

              h("td", () => {
                spec({
                  classList: ["px-6", "py-4"],
                });
                h("div", {
                  text: action.map((a) => (a.disabled ? "Disabled" : "Active")),
                });
              });

              h("td", () => {
                spec({
                  classList: ["px-6", "py-4"],
                });

                const localEditClicked = createEvent();
                sample({
                  source: remap(action, "id"),
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
