import { $shouldCopyToClipboard, bindLinkNavigation, i18n, isObjectPresent } from "@/shared/lib";
import {
  Button,
  ChangesTable,
  Infobox,
  Input,
  KBD,
  LegendIndicator,
  LevelBadge,
  Select,
  Spinner,
  triggerTooltip,
} from "@/shared/ui";
import { ClassListArray, h, list, node, remap, spec } from "forest";
import {
  $aiSummary,
  $fieldsOptions,
  $filterList,
  $insightLogs,
  $insightsSchemas,
  $isLoadingLogs,
  events,
  InsightSchema,
} from "./model";
import { createEffect, createEvent, createStore, sample } from "effector";
import { logModel } from "@/entities/log";
import { Log } from "@/shared/api";
import { userModel } from "@/entities/user";
import { Link, logsRoute } from "@/shared/routing";

export const InsightsBuilder = () => {
  h("div", () => {
    InsightsFilters();

    h("div", () => {
      spec({ visible: $insightLogs.map((logs) => logs.length > 0) });

      h("div", () => {
        spec({ visible: $isLoadingLogs.map((l) => !l) });

        h("div", () => {
          spec({ classList: ["mt-6"] });
          Infobox({ text: $aiSummary, emoji: "✨", visible: userModel.$currentAccount.map((u) => !!u.insights_enabled) });
        });

        InsightsSchemas();
        InsightLogsTable();
      });
    });

    h("div", () => {
      spec({
        classList: ["absolute", "top-1/2", "left-1/2"],
      });

      h("div", () => {
        spec({ classList: ["relative", "right-1/2"] });

        Spinner({ visible: $isLoadingLogs });
      });
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
        spec({ classList: ["grid", "grid-cols-2", "md:grid-cols-3", "items-center", "gap-x-3", "gap-y-3"] });

        h("div", () => {
          spec({ classList: [] });

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
          disableMargin: true,
        });

        h("div", () => {
          spec({ classList: ["mt-auto"] });
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
        "md:sticky",
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
        const params = {
          schemaName: remap(schema, "schemaName"),
        };
        const query = $filterList.map((filters) => {
          return { f: filters.map((f) => `${f.fieldName}=${f.fieldValue}`).join("&") };
        });

        const { click, mounted } = bindLinkNavigation({ params, route: logsRoute });

        Link(logsRoute, {
          params,
          query,
          handler: {
            config: { prevent: true, capture: true, stop: true },
            on: { click },
          },
          fn() {
            h("div", () => {
              spec({ classList: ["max-w-fit"] });
              LegendIndicator({
                text: remap(schema, "schemaTitle"),
                color: remap(schema, "schemaColor"),
              });
            });

            node((el) => {
              mounted(el);
            });
          },
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
          list($insightLogs, ({ store: log, key: $idx }) => {
            const $classes = createStore("");
            const touch = createEvent();
            sample({
              source: { schemas: $insightsSchemas, logs: $insightLogs, idx: $idx },
              clock: touch,
              fn: ({ schemas, logs, idx }) => logRowClasses(schemas, logs, idx),
              target: $classes,
            });

            h("tr", () => {
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

const logRowClasses = (schemas: InsightSchema[], logs: Log[], idx: number): string => {
  const prevIdx = idx - 1;
  const nextIdx = idx + 1;
  const prev = prevIdx >= 0 ? logs.at(prevIdx) : null;
  const curr = logs.at(idx);
  const next = nextIdx < logs.length ? logs.at(nextIdx) : null;

  const prevSchema = schemas.find((s) => s.schemaId === prev?.schema_id);
  const currSchema = schemas.find((s) => s.schemaId === curr?.schema_id);
  const nextSchema = schemas.find((s) => s.schemaId === next?.schema_id);

  const color = currSchema?.schemaColor;
  const btlr = !prevSchema || prevSchema.schemaId != currSchema?.schemaId ? "xl" : null;
  const bblr = !nextSchema || nextSchema.schemaId != currSchema?.schemaId ? "xl" : null;

  const colorClasses = [`before:border-${color}`, `before:rounded-tl-${btlr}`, `before:rounded-bl-${bblr}`];
  const classes = colorClasses.concat(baseRowClasses);

  return classes.join(" ");
};
