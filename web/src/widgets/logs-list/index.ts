import { logModel } from "@/entities/log";
import { showLogRoute } from "@/routing/shared";
import { CardHeaded, Table } from "@/shared/ui";
import { createStore } from "effector";
import { h, list, spec } from "forest";

export const LogsList = () => {
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
            rows: group.map((g) => g.formattedLogs),
          });
        },
      });
    });
  });
};
