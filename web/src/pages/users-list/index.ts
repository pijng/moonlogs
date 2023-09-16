import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { memberCreateRoute, membersRoute } from "@/routing/shared";
import { UsersList } from "@/widgets/users-list";
import { HeaderWithCreation } from "@/widgets";

export const UsersListPage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(membersRoute);

    HeaderWithCreation("Members", memberCreateRoute);

    h("div", () => {
      spec({
        classList: ["pt-3"],
      });

      UsersList();
    });
  });
};
