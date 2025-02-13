import { alertingRuleModel } from "@/entities/alerting-rule";
import { alertingRuleEditRoute } from "@/shared/routing";
import { AlertingRulesTable } from "@/shared/ui";
import { createEvent, sample } from "effector";
import { h, spec } from "forest";

export const AlertingRulesList = () => {
  h("div", () => {
    spec({
      classList: ["flex", "flex-col", "space-y-6", "mt-3"],
    });

    const editRuleClicked = createEvent<number>();
    sample({
      source: editRuleClicked,
      fn: (id) => ({ id }),
      target: alertingRuleEditRoute.open,
    });

    AlertingRulesTable(alertingRuleModel.$alertingRules, editRuleClicked);
  });
};
