import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { logModel } from "@/entities/log";
import { router, showLogRoute } from "@/shared/routing";
import { Spinner } from "@/shared/ui";
import { GroupActionsList } from "@/widgets/group-actions-list";
import { SchemaHeader } from "@/widgets/schema-header";
import { LogsCard } from "@/widgets/logs-card";
import { schemaModel } from "@/entities/schema";
import { combine } from "effector";

export const ShowLogPage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(showLogRoute);

    SchemaHeader();

    const $logsPresent = logModel.$groupedLogs.map((g) => g.logs.length > 0);

    h("div", () => {
      h("div", () => {
        spec({
          classList: ["inline-flex", "space-x-2", "pt-6", "pb-3"],
        });

        GroupActionsList({});
      });

      h("div", () => {
        spec({
          classList: ["flex", "flex-col", "space-y-6"],
          visible: $logsPresent,
        });

        const $activeSchema = combine([router.$activeRoutes, schemaModel.$schemas], ([activeRoutes, schemas]) => {
          const schemaName = activeRoutes[0]?.$params.getState().schemaName;
          return schemas.find((s) => s.name === schemaName) || null;
        });

        LogsCard({
          schema: $activeSchema,
          logsGroup: logModel.$groupedLogs,
        });
      });
    });

    h("div", () => {
      spec({
        classList: ["absolute", "top-1/2", "left-1/2"],
      });

      h("div", () => {
        spec({ classList: ["relative", "right-1/2"] });

        Spinner({ visible: $logsPresent.map((lp) => !lp) });
      });
    });
  });
};
