import { logModel } from "@/entities/log";
import { schemaModel } from "@/entities/schema";
import { router } from "@/routing";
import { CardHeaded, LogsTable } from "@/shared/ui";
import { combine } from "effector";
import { h, list, spec } from "forest";

export const LogsList = () => {
  h("div", () => {
    spec({
      classList: ["flex", "flex-col", "space-y-6", "mt-2"],
    });

    const $activeSchema = combine([router.$activeRoutes, schemaModel.$schemas], ([activeRoutes, schemas]) => {
      const schemaName = activeRoutes[0]?.$params.getState().schemaName;
      return schemas.find((s) => s.name === schemaName) || null;
    });

    list(logModel.$logsGroups, ({ store: group }) => {
      CardHeaded({
        tags: group.map((g) => g.tags),
        kind: group.map((g) => g.kind),
        schema: $activeSchema,
        href: group.map((g) => `${g.schema_name}/${g.group_hash}`),
        content: () => {
          LogsTable(group.map((g) => g.logs));
        },
      });
    });
  });
};
