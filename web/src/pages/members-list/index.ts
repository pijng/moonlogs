import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { membersRoute } from "@/routing/shared";
import { LogsList, SchemaHeader, SearchBar } from "@/widgets";
import { Pagination } from "@/shared/ui";
import { logModel } from "@/entities/log";

export const MembersListPage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(membersRoute);

    SchemaHeader();

    h("div", () => {
      spec({
        classList: ["mt-5"],
      });

      // SearchBar();
      // Pagination(logModel.$pages, logModel.$currentPage, logModel.events.pageChanged);
      // LogsList();
    });
  });
};
