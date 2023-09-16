import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { memberEditRoute } from "@/routing/shared";
import { Header } from "@/shared/ui";
import { EditMemberForm } from "@/features/user-edit";

export const UserEditPage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(memberEditRoute);

    Header("Edit member");

    h("div", () => {
      spec({
        classList: ["pt-5", "max-w-2xl"],
      });

      EditMemberForm();
    });
  });
};
