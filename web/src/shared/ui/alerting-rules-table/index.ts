import { h, list, remap, spec } from "forest";
import { Event, Store, createEvent, sample } from "effector";
import { AlertingRule } from "@/shared/api";
import { Button } from "@/shared/ui";
import { i18n } from "@/shared/lib/i18n";

export const AlertingRulesTable = (rules: Store<AlertingRule[]>, editClicked: Event<number>) => {
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
              text: i18n("tables.alerting_rules.name"),
            });

            h("th", {
              attr: { scope: "col" },
              classList: ["px-6", "py-3"],
              text: i18n("tables.alerting_rules.enabled"),
            });

            h("th", {
              attr: { scope: "col" },
              classList: ["px-6", "py-3"],
              text: i18n("tables.alerting_rules.severity"),
            });

            h("th", {
              attr: { scope: "col" },
              classList: ["px-6", "py-3"],
              text: i18n("tables.alerting_rules.interval"),
            });

            h("th", {
              attr: { scope: "col" },
              classList: ["px-6", "py-3"],
              text: i18n("tables.alerting_rules.threshold"),
            });

            h("th", {
              attr: { scope: "col" },
              classList: ["px-6", "py-3", "max-w-28", "w-28"],
              text: i18n("tables.alerting_rules.actions"),
            });
          });
        });

        h("tbody", () => {
          list(rules, ({ store: rule }) => {
            h("tr", () => {
              spec({
                classList: {
                  "border-t": true,
                  "w-full": true,
                  "dark:border-shadow-gray": true,
                  "hover:bg-gray-50": true,
                  "dark:hover:bg-slate-gray": true,
                  "text-gray-900": true,
                  "dark:text-gray-200": true,
                  "opacity-50": rule.map((r) => !r.enabled),
                },
              });

              h("td", () => {
                spec({
                  classList: ["px-6", "py-4"],
                });
                h("div", {
                  text: remap(rule, "name"),
                });
              });

              h("td", () => {
                spec({
                  classList: ["px-6", "py-4"],
                });
                h("div", {
                  text: rule.map((r) => (r.enabled ? "Active" : "Disabled")),
                });
              });

              h("td", () => {
                spec({
                  classList: ["px-6", "py-4"],
                });
                h("div", {
                  text: remap(rule, "severity"),
                });
              });

              h("td", () => {
                spec({
                  classList: ["px-6", "py-4"],
                });
                h("div", {
                  text: remap(rule, "interval"),
                });
              });

              h("td", () => {
                spec({
                  classList: ["px-6", "py-4"],
                });
                h("div", {
                  text: remap(rule, "threshold"),
                });
              });

              h("td", () => {
                spec({
                  classList: ["px-6", "py-4"],
                });

                const localEditClicked = createEvent();
                sample({
                  source: remap(rule, "id"),
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
