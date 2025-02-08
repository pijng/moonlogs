import { h, list, remap, spec } from "forest";
import { Store } from "effector";
import { Log } from "@/shared/api";
import { isObjectPresent } from "@/shared/lib";
import { Change, diffWords } from "diff";

export const ChangesTable = (log: Store<Log>) => {
  h("div", () => {
    spec({
      visible: log.map((l) => isObjectPresent(l.changes)),
      classList: ["relative", "mt-4"],
    });

    const $changesList = log.map((l) => {
      const changesValues = l.changes || {};
      const changesContainer: { name: string; before: Change[]; after: Change[] }[] = [];

      for (const [name, values] of Object.entries(changesValues)) {
        const changes = diffWords(JSON.stringify(values.old_value), JSON.stringify(values.new_value));
        const before = changes.filter((p) => !p.added);
        const after = changes.filter((p) => !p.removed);

        changesContainer.push({ name: name, before: before, after: after });
      }

      return changesContainer;
    });

    list($changesList, ({ store: changes }) => {
      h("table", () => {
        spec({ classList: ["w-full", "text-sm", "text-left", "rtl:text-right", "text-gray-900", "dark:text-gray-100"] });

        h("thead", () => {
          spec({
            classList: [
              "border",
              "border-gray-200",
              "text-xs",
              "text-gray-900",
              "dark:text-gray-100",
              "uppercase",
              "bg-gray-50",
              "dark:border-shadow-gray",
              "dark:bg-squid-ink",
            ],
          });

          h("tr", () => {
            h("th", {
              classList: ["px-3", "py-2", "text-center"],
              attr: { scope: "col" },
              text: remap(changes, "name"),
            });
          });
        });

        h("tbody", () => {
          spec({ classList: ["border", "dark:border-shadow-gray", "border-gray-200"] });

          h("tr", () => {
            spec({
              classList: ["grid", "grid-cols-2", "bg-white", "dark:bg-raisin-black"],
            });

            h("td", () => {
              spec({
                classList: ["px-3", "py-2", "border-r-2", "pr-4", "border-gray-200", "border-dashed", "dark:border-shadow-gray"],
              });

              const $before = remap(changes, "before");
              list($before, ({ store: before }) => {
                h("span", {
                  text: remap(before, "value"),
                  classList: {
                    "whitespace-pre-wrap": true,
                    "break-words": true,
                    "bg-red-200": before.map((p) => Boolean(p.removed)),
                    "dark:bg-red-800": before.map((p) => Boolean(p.removed)),
                  },
                });
              });
            });

            h("td", () => {
              spec({
                classList: ["px-3", "py-2"],
              });

              const $after = remap(changes, "after");

              list($after, ({ store: after }) => {
                h("span", {
                  text: remap(after, "value"),
                  classList: {
                    "whitespace-pre-wrap": true,
                    "break-words": true,
                    "bg-green-200": after.map((p) => Boolean(p.added)),
                    "dark:bg-green-800": after.map((p) => Boolean(p.added)),
                  },
                });
              });
            });
          });
        });
      });
    });
  });
};
