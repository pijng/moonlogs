import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { logModel } from "@/entities/log";
import { router, showLogRoute } from "@/shared/routing";
import { CardHeaded, LogsTable, Spinner } from "@/shared/ui";
import { SchemaHeader } from "@/widgets";
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
      spec({
        classList: ["mt-3"],
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

        CardHeaded({
          tags: logModel.$groupedLogs.map((g) => g.tags),
          schema: $activeSchema,
          kind: logModel.$groupedLogs.map((g) => g.kind),
          content: () => {
            LogsTable({
              logs: logModel.$groupedLogs.map((g) => g.logs),
              requestClicked: logModel.events.requestURLClicked,
              responseClicked: logModel.events.responseURLClicked,
            });
          },
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
