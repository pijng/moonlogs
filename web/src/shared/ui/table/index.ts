import { h, list, spec, variant } from "forest";
import { Store } from "effector";
import { LevelBadge } from "..";

export type Cell = {
  data: string;
  component: "text" | "badge";
};

export const Table = ({ columns, rows }: { columns: Store<string[]>; rows: Store<Cell[][]> }) => {
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
                const $isLastColumn = idx.map((idx) => ![0, 1].includes(idx));

                spec({
                  attr: { scope: "col" },
                  classList: {
                    "px-6": true,
                    "py-3": true,
                    "w-24": $isLastColumn.map((state) => !state),
                    "lg:w-48": $isLastColumn.map((state) => !state),
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
                const $componentName = cell.map((c) => ({ name: c.component }));

                h("td", () => {
                  spec({
                    classList: ["px-6", "py-4"],
                  });

                  variant({
                    source: $componentName,
                    key: "name",
                    cases: {
                      text: () => {
                        h("div", {
                          text: cell.map((c) => c.data),
                        });
                      },
                      badge: () => {
                        LevelBadge(cell.map((c) => c.data));
                      },
                      __: () => {
                        h("div", {
                          text: cell.map((c) => c.data),
                        });
                      },
                    },
                  });
                });
              });
            });
          });
        });
      });
    });
  });
};
