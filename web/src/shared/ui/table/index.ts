import { h, list, spec } from "forest";
import { Store } from "effector";

export const Table = ({ columns, rows }: { columns: Store<string[]>; rows: Store<string[][]> }) => {
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
            list(columns, ({ store: column, key: idx }) => {
              h("th", () => {
                const $isFirstColumn = idx.map((idx) => idx === 0);

                spec({
                  attr: { scope: "col" },
                  classList: {
                    "px-6": true,
                    "py-3": true,
                    "w-48": $isFirstColumn,
                  },
                  text: column,
                });
              });
            });
          });
        });

        h("tbody", () => {
          list(rows, ({ store: row }) => {
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

              list(row, ({ store: cell }) => {
                h("td", {
                  classList: ["px-6", "py-4"],
                  text: cell,
                });
              });
            });
          });
        });
      });
    });
  });
};
