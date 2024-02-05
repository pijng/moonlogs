import { h, list, remap, spec } from "forest";
import { Store } from "effector";
import { LevelBadge } from "..";
import { Log } from "@/shared/api";

export const LogsTable = (logs: Store<Log[]>) => {
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
              classList: ["px-4", "py-3", "w-24", "lg:w-48"],
              text: "Time",
            });

            h("th", {
              attr: { scope: "col" },
              classList: ["px-4", "py-3", "w-16"],
              text: "Level",
            });

            h("th", {
              attr: { scope: "col" },
              classList: ["px-4", "py-3"],
              text: "Text",
            });
          });
        });

        h("tbody", () => {
          list(logs, ({ store: log }) => {
            h("tr", () => {
              spec({
                classList: [
                  "border-t",
                  "w-full",
                  "dark:border-gray-700",
                  "hover:bg-gray-50",
                  "dark:hover:bg-gray-600",
                  "text-gray-900",
                  "dark:text-gray-100",
                ],
              });

              h("td", () => {
                spec({
                  classList: ["px-4", "py-4"],
                });
                h("div", {
                  text: remap(log, "created_at"),
                });
              });

              h("td", () => {
                spec({
                  classList: ["px-4", "py-4"],
                });
                LevelBadge(remap(log, "level"));
              });

              h("td", () => {
                spec({
                  classList: ["px-4", "py-4"],
                });
                h("div", {
                  text: remap(log, "text"),
                });
              });
            });
          });
        });
      });
    });
  });
};
