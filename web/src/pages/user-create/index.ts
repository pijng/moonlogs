import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { memberCreateRoute } from "@/shared/routing";
import { Header } from "@/shared/ui";
import { NewMemberForm } from "@/features/user/user-create";
import { i18n } from "@/shared/lib/i18n";

export const UserCreatePage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(memberCreateRoute);

    Header(i18n("members.form.actions.create"));

    h("div", () => {
      spec({
        classList: ["pt-5", "max-w-2xl"],
      });

      NewMemberForm();
    });
  });
};
