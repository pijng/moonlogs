import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { logsRoute } from "@/shared/routing";
import { LogsList, SchemaHeader, SearchBar } from "@/widgets";
import { Pagination, Spinner } from "@/shared/ui";
import { logModel } from "@/entities/log";

export const LogsListPage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(logsRoute);

    SchemaHeader();

    h("div", () => {
      spec({
        classList: ["pt-3"],
      });

      const $isLoadingLogs = logModel.effects.queryLogsFx.pending;

      SearchBar();

      h("div", () => {
        Pagination(logModel.$pages, logModel.$currentPage, logModel.events.pageChanged);

        h("div", () => {
          spec({ visible: $isLoadingLogs.map((p) => !p) });

          LogsList();
        });

        h("div", () => {
          spec({
            classList: ["pt-4"],
            visible: $isLoadingLogs.map((l) => !l),
          });
          Pagination(logModel.$pages, logModel.$currentPage, logModel.events.pageChanged);
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
  });
};
