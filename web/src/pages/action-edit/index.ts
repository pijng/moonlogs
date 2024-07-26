import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { actionEditRoute } from "@/shared/routing";
import { Header } from "@/shared/ui";
import { EditActionForm } from "@/features/action-edit";
import { i18n } from "@/shared/lib/i18n";

export const ActionEditPage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(actionEditRoute);

    Header(i18n("actions.form.actions.edit"));

    h("div", () => {
      spec({
        classList: ["pt-5", "max-w-2xl"],
      });

      EditActionForm();
    });
  });
};
