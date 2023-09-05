import { withRoute } from "atomic-router-forest";
import { h } from "forest";

import { membersRoute } from "@/routing/shared";
import { Header } from "@/shared/ui";
import { UsersList } from "@/widgets/users-list";

export const MembersListPage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(membersRoute);

    Header("Members");

    UsersList();
  });
};
