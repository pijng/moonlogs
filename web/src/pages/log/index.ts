import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { logModel } from "@/entities/log";
import { showLogRoute } from "@/routing/shared";
import { CardHeaded, LogsTable } from "@/shared/ui";
import { SchemaHeader } from "@/widgets";

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

        CardHeaded({
          tags: logModel.$groupedLogs.map((g) => g.tags),
          content: () => {
            LogsTable(logModel.$groupedLogs.map((g) => g.logs));
          },
        });
      });
    });
  });
};
