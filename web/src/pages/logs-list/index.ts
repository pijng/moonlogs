import { withRoute } from "atomic-router-forest";
import { h, list, spec } from "forest";

import { logModel } from "@/entities/log";
import { logsRoute, showLogRoute } from "@/routing/shared";
import { CardHeaded, Table } from "@/shared/ui";
import { SearchBar } from "@/widgets";
import { combine, createStore } from "effector";
import { schemaModel } from "@/entities/schema";

export const LogsListPage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(logsRoute);

    const $schemaTitle = combine([logsRoute.$params, schemaModel.$schemas], ([params, schemas]) => {
      return schemas.find((s) => s.name === params.schemaName)?.title || "";
    });

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
      text: $schemaTitle,
    });

    h("div", () => {
      spec({
        classList: ["mt-5"],
      });

      SearchBar(logModel.events.queryChanged, logModel.$searchQuery);

      h("div", () => {
        spec({
          classList: ["flex", "flex-col", "space-y-6", "mt-2"],
        });

        list(logModel.$logsGroups, ({ store: group }) => {
          CardHeaded({
            tags: group.map((g) => g.tags),
            routeConfig: {
              route: showLogRoute,
              payload: group.map((g) => ({ schemaName: g.schema_name, hash: g.group_hash }) as Record<string, any>),
            },
            content: () => {
              Table({
                columns: createStore(["Time", "Text"]),
                rows: group.map((g) => g.logs.map((l) => [l.created_at, l.text])),
              });
            },
          });
        });
      });
    });
  });
};
