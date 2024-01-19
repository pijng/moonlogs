import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { logModel } from "@/entities/log";
import { router, showLogRoute } from "@/routing/shared";
import { CardHeaded, LogsTable } from "@/shared/ui";
import { SchemaHeader } from "@/widgets";
import { schemaModel } from "@/entities/schema";
import { combine } from "effector";

export const ShowLogPage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(showLogRoute);

    SchemaHeader();

    h("div", () => {
      spec({
        classList: ["mt-3"],
      });

      h("div", () => {
        spec({
          classList: ["flex", "flex-col", "space-y-6"],
        });

        const $activeSchema = combine([router.$activeRoutes, schemaModel.$schemas], ([activeRoutes, schemas]) => {
          const schemaName = activeRoutes[0]?.$params.getState().schemaName;
          return schemas.find((s) => s.name === schemaName) || null;
        });

        CardHeaded({
          tags: logModel.$groupedLogs.map((g) => g.tags),
          schema: $activeSchema,
          kind: logModel.$groupedLogs.map((g) => g.kind),
          content: () => {
            LogsTable(logModel.$groupedLogs.map((g) => g.logs));
          },
        });
      });
    });
  });
};
