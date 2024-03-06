import { h, spec } from "forest";
import { withRoute } from "atomic-router-forest";

import { homeRoute, notFoundRoute } from "@/routing/shared";
import { Header, Link } from "@/shared/ui";
import { i18n } from "@/shared/lib/i18n";

export const NotFoundPage = () => {
  h("div", {
    fn() {
      // This allows to show/hide route if page is matched
      // It is required to call `withRoute` inside `h` call
      withRoute(notFoundRoute);

      h("div", () => {
        spec({
          classList: ["h-screen", "flex", "flex-col", "items-center", "justify-center", "px-7", "text-center"],
        });

        h("div", () => {
          spec({ classList: ["mb-4"] });
          Header(i18n("miscellaneous.not_found"));
        });

        // Duh
        Link(homeRoute, i18n("miscellaneous.to_home"));
      });
    },
  });
};
