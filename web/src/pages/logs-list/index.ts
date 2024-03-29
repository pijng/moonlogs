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

      SearchBar();

      h("div", () => {
        Pagination(logModel.$pages, logModel.$currentPage, logModel.events.pageChanged);

        Spinner({ visible: logModel.effects.queryLogsFx.pending });

        h("div", () => {
          spec({ visible: logModel.effects.queryLogsFx.pending.map((p) => !p) });

          LogsList();
        });

        h("div", () => {
          spec({ classList: ["pt-4"] });
          Pagination(logModel.$pages, logModel.$currentPage, logModel.events.pageChanged);
        });
      });
    });
  });
};
