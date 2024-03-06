import { h, spec } from "forest";
import { withRoute } from "atomic-router-forest";

import { forbiddenRoute, homeRoute } from "@/routing/shared";
import { Header, Link } from "@/shared/ui";
import { i18n } from "@/shared/lib/i18n";

export const ForbiddenPage = () => {
  h("div", {
    fn() {
      // This allows to show/hide route if page is matched
      // It is required to call `withRoute` inside `h` call
      withRoute(forbiddenRoute);

      h("div", () => {
        spec({
          classList: ["h-screen", "flex", "flex-col", "items-center", "justify-center", "px-7", "text-center"],
        });

        h("div", () => {
          spec({ classList: ["mb-4"] });
          Header(i18n("miscellaneous.forbidden"));
        });

        Link(homeRoute, i18n("miscellaneous.to_home"));
      });
    },
  });
};
