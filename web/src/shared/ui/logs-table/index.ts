import { h, list, remap, spec } from "forest";
import { Event, Store, createEvent, sample } from "effector";
import { KBD, LevelBadge } from "..";
import { Log } from "@/shared/api";
import { isObjectPresent } from "@/shared/lib";

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
                classList: ["border-t", "w-full", "dark:border-gray-700", "text-gray-900", "dark:text-gray-100"],
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
                        "overflow-scroll",
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
                      text: "Request",
                      event: localRequestClicked,
                      visible: remap(log, "request").map(isObjectPresent),
                    });

                    KBD({
                      text: "Response",
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
