import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { alertingRuleEditRoute } from "@/shared/routing";
import { Header } from "@/shared/ui";
import { i18n } from "@/shared/lib/i18n";
import { EditAlertingRuleForm } from "@/features/alerting-rule-edit/ui";

export const AlertingRuleEditPage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(alertingRuleEditRoute);

    Header(i18n("alerting_rules.form.actions.edit"));

    h("div", () => {
      spec({
        classList: ["pt-5", "max-w-2xl"],
      });

      EditAlertingRuleForm();
    });
  });
};
