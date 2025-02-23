import { $shouldCopyToClipboard, i18n, isObjectPresent } from "@/shared/lib";
import { Button, ChangesTable, Input, KBD, LegendIndicator, LevelBadge, Select, triggerTooltip } from "@/shared/ui";
import { ClassListArray, h, list, node, remap, spec } from "forest";
import { $fieldsOptions, $filterList, $insightLogs, $insightsSchemas, events, InsightSchema } from "./model";
import { combine, createEffect, createEvent, createStore, restore, sample, Store } from "effector";
import { logModel } from "@/entities/log";
import { Log } from "@/shared/api";

export const InsightsBuilder = () => {
  h("div", () => {
    InsightsFilters();

    h("div", () => {
      spec({ visible: $insightLogs.map((logs) => logs.length > 0) });

      InsightsSchemas();

      InsightLogsTable();
    });
  });
};

const InsightsFilters = () => {
  h("div", () => {
    spec({
      classList: ["max-w-5xl"],
    });

    list($filterList, ({ store: $filter, key: idx }) => {
      const fieldNameSelected = createEvent<string>();

      sample({
        source: idx,
        clock: fieldNameSelected,
        fn: (idx, fieldName) => ({ fieldName: fieldName, idx: idx }),
        target: events.fieldNameChanged,
      });

      const fieldValueChanged = createEvent<string>();
      sample({
        source: idx,
        clock: fieldValueChanged,
        fn: (idx, fieldValue) => ({ fieldValue: fieldValue, idx: idx }),
        target: events.fieldValueChanged,
      });

      h("div", () => {
        spec({ classList: ["grid", "grid-cols-3", "items-center", "gap-3"] });

        h("div", () => {
          spec({ classList: ["mb-6"] });

          Select({
            text: i18n("insights.form.field_name"),
            value: remap($filter, "fieldName"),
            options: $fieldsOptions,
            optionSelected: fieldNameSelected,
            withBlank: createStore(false),
          });
        });

        Input({
          type: "text",
          label: i18n("insights.form.field_value"),
          required: true,
          value: remap($filter, "fieldValue"),
          inputChanged: fieldValueChanged,
        });

        h("div", () => {
          Button({
            text: i18n("buttons.build_insight"),
            event: events.buildInsight,
            size: "base",
            prevent: true,
            variant: "default",
          });
        });
      });
    });
  });
};

const InsightsSchemas = () => {
  h("div", () => {
    spec({
      classList: [
        "ml-[1]",
        "px-2",
        "pt-3",
        "pb-3",
        "w-full",
        "sticky",
        "left-1",
        "top-0",
        "z-10",
        "dark:bg-eigengrau",
        "bg-white",
      ],
    });

    h("div", () => {
      spec({ classList: ["flex", "flex-wrap", "gap-3"] });

      list($insightsSchemas, ({ store: schema }) => {
        h("div", () => {
          spec({ classList: ["max-w-fit"] });
          LegendIndicator({
            text: remap(schema, "schemaTitle"),
            color: remap(schema, "schemaColor"),
          });
        });
      });
    });
  });
};

const InsightLogsTable = () => {
  h("div", () => {
    spec({
      classList: [
        "mt-2",
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
        classList: ["relative", "dark:scrollbar", "sm:rounded-b-lg"],
      });

      h("table", () => {
        spec({
          classList: ["border-separate", "border-spacing-0", "w-full", "text-sm", "text-left", "table-fixed"],
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
              // "sticky",
              // "top-0",
              // "z-10",
            ],
          });

          h("tr", () => {
            h("th", {
              attr: { scope: "col" },
              classList: ["px-4", "py-3", "w-32", "lg:w-52"],
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
          list($insightLogs, ({ store: log, key: idx }) => {
            h("tr", () => {
              const touch = createEvent();
              const touchClasses = createEvent<string>();
              const $classes = restore(touchClasses, "");

              sample({
                clock: touch,
                source: logRowClasses($insightsSchemas, $insightLogs, idx),
                target: touchClasses,
              });

              spec({
                classList: ["w-full", "text-gray-900", "dark:text-gray-100"],
              });

              h("td", () => {
                spec({
                  classList: [$classes] as ClassListArray,
                });
                h("div", {
                  text: remap(log, "created_at"),
                });

                node(() => {
                  touch();
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
                  source: [$formattedText, $shouldCopyToClipboard] as const,
                  clock: textClicked,
                  filter: ([, shouldCopy]) => !Boolean(window.getSelection()?.toString()) && shouldCopy,
                  fn: ([text]) => text,
                  target: copyTextFx,
                });

                sample({
                  source: i18n("miscellaneous.copied_to_clipboard"),
                  clock: copyTextFx.done,
                  fn: (text) => ({ text: text }),
                  target: triggerTooltip,
                });

                h("div", {
                  classList: {
                    "whitespace-pre-wrap": true,
                    "break-words": true,
                    "cursor-pointer": $shouldCopyToClipboard,
                  },
                  text: $formattedText,
                  handler: { click: textClicked },
                });

                ChangesTable(log);

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
                      target: logModel.events.requestURLClicked,
                    });

                    sample({
                      source: log.map((l) => parseInt(l.id)),
                      clock: localResponseClicked,
                      target: logModel.events.responseURLClicked,
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

const baseRowClasses = [
  "relative",
  "before:content-['']",
  "before:left-[-10]",
  "before:top-0",
  "before:h-full",
  "before:absolute",
  "before:border-l-10",
  "px-4",
  "py-4",
  "align-top",
];

const logRowClasses = (schemas: Store<InsightSchema[]>, logs: Store<Log[]>, idx: Store<number>): Store<string> => {
  const logsMatrix = combine(logs, idx, (logs, idx) => {
    const prevIdx = idx - 1;
    const nextIdx = idx + 1;
    const prev = prevIdx >= 0 ? logs.at(prevIdx) : null;
    const curr = logs.at(idx);
    const next = nextIdx < logs.length ? logs.at(nextIdx) : null;

    return { prev, curr, next };
  });

  const colors = combine(schemas, logsMatrix, (schemas, logs) => {
    const prevSchema = !!logs.prev ? schemas.find((s) => s.schemaId === logs.prev?.schema_id) : null;
    const currSchema = !!logs.curr ? schemas.find((s) => s.schemaId === logs.curr?.schema_id) : null;
    const nextSchema = !!logs.next ? schemas.find((s) => s.schemaId === logs.next?.schema_id) : null;

    const color = currSchema!.schemaColor;
    const btlr = !prevSchema || prevSchema.schemaId != currSchema?.schemaId ? "xl" : null;
    const bblr = !nextSchema || nextSchema.schemaId != currSchema?.schemaId ? "xl" : null;

    return [`before:border-${color}`, `before:rounded-tl-${btlr}`, `before:rounded-bl-${bblr}`] as string[];
  });

  return colors.map((colors) => baseRowClasses.concat(colors).join(" "));
};
