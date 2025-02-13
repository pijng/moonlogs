import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { alertingRuleCreateRoute, alertingRulesRoute } from "@/shared/routing";
import { HeaderWithCreation } from "@/widgets/header-with-creation";

import { i18n } from "@/shared/lib/i18n";
import { AlertingRulesList } from "@/widgets/alerting-rules-list";

export const AlertingRulesListPage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(alertingRulesRoute);

    HeaderWithCreation(i18n("alerting_rules.label"), alertingRuleCreateRoute);

    h("div", () => {
      spec({ classList: ["pt-3"] });

      AlertingRulesList();
    });
  });
};
