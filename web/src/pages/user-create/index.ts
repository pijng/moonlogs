import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { memberCreateRoute } from "@/routing/shared";
import { Header } from "@/shared/ui";
import { NewMemberForm } from "@/features/user-create";

export const UserCreatePage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(memberCreateRoute);

    Header("Create member");

    h("div", () => {
      spec({
        classList: ["pt-5", "max-w-2xl"],
      });

      NewMemberForm();
    });
  });
};
