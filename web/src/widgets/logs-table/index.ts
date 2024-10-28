import { h, list, remap, spec } from "forest";
import { Event, Store, createEffect, createEvent, sample } from "effector";
import { Log } from "@/shared/api";
import { isObjectPresent } from "@/shared/lib";
import { i18n } from "@/shared/lib/i18n";
import { DiffText, KBD, LevelBadge, triggerTooltip } from "@/shared/ui";

export const LogsTable = ({
  logs,
  requestClicked,
  responseClicked,
}: {
  logs: Store<Log[]>;
  requestClicked: Event<number>;
  responseClicked: Event<number>;
}) => {
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
          classList: ["w-full", "text-sm", "text-left", "table-fixed"],
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
              classList: ["px-4", "py-3", "w-24", "lg:w-48"],
              text: i18n("tables.log_groups.time"),
            });

            h("th", {
              attr: { scope: "col" },
              classList: ["px-4", "py-3", "w-24"],
              text: i18n("tables.log_groups.level"),
            });

            h("th", {
              attr: { scope: "col" },
              classList: ["px-4", "py-3"],
              text: i18n("tables.log_groups.text"),
            });
          });
        });

        h("tbody", () => {
          list(logs, ({ store: log }) => {
            h("tr", () => {
              spec({
                classList: ["border-t", "w-full", "dark:border-shadow-gray", "text-gray-900", "dark:text-gray-100"],
              });

              h("td", () => {
                spec({
                  classList: ["px-4", "py-4", "align-top"],
                });
                h("div", {
                  text: remap(log, "created_at"),
                });
              });

              h("td", () => {
                spec({
                  classList: ["px-4", "py-4", "align-top"],
                });
                LevelBadge(remap(log, "level"));
              });

              h("td", () => {
                spec({
                  classList: ["relative", "px-4", "py-4"],
                });

                const $formattedText = remap(log, "text").map((t) => t.replaceAll("\\n", "\n"));
                const textClicked = createEvent<MouseEvent>();
                const copyTextFx = createEffect((clickedText: string) => {
                  return navigator.clipboard.writeText(clickedText);
                });

                sample({
                  source: $formattedText,
                  clock: textClicked,
                  filter: () => !Boolean(window.getSelection()?.toString()),
                  target: copyTextFx,
                });

                sample({
                  source: i18n("miscellaneous.copied_to_clipboard"),
                  clock: copyTextFx.done,
                  fn: (text) => ({ text: text }),
                  target: triggerTooltip,
                });

                h("div", {
                  classList: ["whitespace-pre-wrap", "break-words", "cursor-pointer"],
                  text: $formattedText,
                  handler: { click: textClicked },
                });

                h("div", () => {
                  spec({ visible: log.map((l) => isObjectPresent(l.changes)) });

                  const $changesList = log.map((l) => {
                    const changesValues = l.changes || {};
                    const changesContainer: { name: string; old_value: any; new_value: any }[] = [];

                    for (const [name, changes] of Object.entries(changesValues)) {
                      changesContainer.push({ name: name, old_value: changes.old_value, new_value: changes.new_value });
                    }

                    return changesContainer;
                  });

                  list($changesList, ({ store: changes }) => {
                    h("div", () => {
                      spec({ classList: ["mt-3"] });

                      h("div", {
                        classList: ["whitespace-pre-wrap", "break-words"],
                        text: changes.map((c) => `{${c.name}}`),
                      });

                      DiffText({
                        oldText: remap(changes, "old_value"),
                        newText: remap(changes, "new_value"),
                      });
                    });
                  });
                });

                h("div", () => {
                  spec({
                    visible: log.map((l) => Boolean(l.old_value) || Boolean(l.new_value)),
                    classList: ["mt-3"],
                  });

                  DiffText({ oldText: log.map((l) => l.old_value), newText: log.map((l) => l.new_value) });
                });

                const $netFieldsPresent = log.map((l) => isObjectPresent(l.request) || isObjectPresent(l.response));

                h("div", () => {
                  spec({ visible: $netFieldsPresent });

                  h("ul", () => {
                    spec({
                      classList: [
                        "flex",
                        "flex-wrap",
                        "gap-3",
                        "basis-11/12",
                        "flex-nowrap",
                        "pt-2",
                        "overflow-auto",
                        "dark:scrollbar",
                        "text-sm",
                        "justify-start",
                        "font-medium",
                        "text-center",
                        "text-gray-500",
                      ],
                    });

                    const localRequestClicked = createEvent();
                    const localResponseClicked = createEvent();

                    sample({
                      source: log.map((l) => parseInt(l.id)),
                      clock: localRequestClicked,
                      target: requestClicked,
                    });

                    sample({
                      source: log.map((l) => parseInt(l.id)),
                      clock: localResponseClicked,
                      target: responseClicked,
                    });

                    KBD({
                      text: i18n("tables.log_groups.request"),
                      event: localRequestClicked,
                      visible: remap(log, "request").map(isObjectPresent),
                    });

                    KBD({
                      text: i18n("tables.log_groups.response"),
                      event: localResponseClicked,
                      visible: remap(log, "response").map(isObjectPresent),
                    });
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
