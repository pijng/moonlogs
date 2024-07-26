import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { actionCreateRoute, actionsRoute } from "@/shared/routing";
import { HeaderWithCreation } from "@/widgets/header-with-creation";
import { i18n } from "@/shared/lib/i18n";
import { ActionsList } from "@/widgets/actions-list";

export const ActionsListPage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(actionsRoute);

    HeaderWithCreation(i18n("actions.label"), actionCreateRoute);

    h("div", () => {
      spec({ classList: ["pt-3", "max-w-5xl"] });

      ActionsList();
    });
  });
};
