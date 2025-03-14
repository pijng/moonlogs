import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { actionCreateRoute } from "@/shared/routing";
import { Header } from "@/shared/ui";
import { NewActionForm } from "@/features/actions/action-create";
import { i18n } from "@/shared/lib/i18n";

export const ActionCreatePage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(actionCreateRoute);

    Header(i18n("actions.form.actions.create"));

    h("div", () => {
      spec({
        classList: ["pt-5", "max-w-2xl"],
      });

      NewActionForm();
    });
  });
};
