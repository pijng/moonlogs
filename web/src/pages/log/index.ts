import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { logModel } from "@/entities/log";
import { showLogRoute } from "@/routing/shared";
import { CardHeaded, Table } from "@/shared/ui";
import { createStore } from "effector";
import { DATEFORMAT_OPTIONS, getLocale } from "@/shared/lib";

export const ShowLogPage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(showLogRoute);

    h("h1", {
      classList: [
        "inline-block",
        "text-2xl",
        "sm:text-3xl",
        "font-extrabold",
        "text-slate-900",
        "tracking-tight",
        "dark:text-slate-200",
      ],
      text: showLogRoute.$params.map((p) => p.schemaName),
    });

    h("div", () => {
      spec({
        classList: ["mt-5"],
      });

      h("div", () => {
        spec({
          classList: ["flex", "flex-col", "space-y-6"],
        });

        const $tags = logModel.$groupedLogs.map((g) => {
          const query = g[0]?.query ?? {};

          return Object.entries(query).map((q) => `${q[0]}: ${q[1]}`);
        });

        CardHeaded({
          tags: $tags,
          content: () => {
            Table({
              columns: createStore(["Time", "Text"]),
              rows: logModel.$groupedLogs.map((logs) =>
                logs.map((log) => {
                  const intl = Intl.DateTimeFormat(getLocale(), DATEFORMAT_OPTIONS);
                  return [intl.format(new Date(log.created_at)), log.text];
                }),
              ),
            });
          },
        });
      });
    });
  });
};
