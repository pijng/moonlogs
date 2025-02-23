import { h, spec } from "forest";
import { withRoute } from "atomic-router-forest";

import { insightsRoute } from "@/shared/routing";
import { Header } from "@/shared/ui";
import { i18n } from "@/shared/lib/i18n";
import { Insights } from "@/widgets/insights";

export const InsightsPage = () => {
  h("div", {
    fn() {
      // This allows to show/hide route if page is matched
      // It is required to call `withRoute` inside `h` call
      withRoute(insightsRoute);

      Header(i18n("insights.label"));

      h("div", () => {
        spec({ classList: ["pt-3"] });

        Insights();
      });
    },
  });
};
