import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { logsRoute } from "@/routing/shared";
import { LogsList, SchemaHeader, SearchBar } from "@/widgets";

export const LogsListPage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(logsRoute);

    SchemaHeader();

    h("div", () => {
      spec({
        classList: ["mt-5"],
      });

      SearchBar();
      LogsList();
    });
  });
};
