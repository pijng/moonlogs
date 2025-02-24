import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { alertingRuleCreateRoute } from "@/shared/routing";
import { Header } from "@/shared/ui";
import { i18n } from "@/shared/lib/i18n";
import { NewAlertingRuleForm } from "@/features/alerting-rule/alerting-rule-create";

export const AlertingRuleCreatePage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(alertingRuleCreateRoute);

    Header(i18n("alerting_rules.form.actions.create"));

    h("div", () => {
      spec({
        classList: ["pt-5", "max-w-2xl", "mb-40"],
      });

      NewAlertingRuleForm();
    });
  });
};
