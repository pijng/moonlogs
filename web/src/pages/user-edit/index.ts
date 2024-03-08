import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { memberEditRoute } from "@/shared/routing";
import { Header } from "@/shared/ui";
import { EditMemberForm } from "@/features/user-edit";
import { i18n } from "@/shared/lib/i18n";

export const UserEditPage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(memberEditRoute);

    Header(i18n("members.form.actions.edit"));

    h("div", () => {
      spec({
        classList: ["pt-5", "max-w-2xl"],
      });

      EditMemberForm();
    });
  });
};
