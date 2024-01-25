import { h, spec } from "forest";
import { withRoute } from "atomic-router-forest";

import { homeRoute, notFoundRoute } from "@/routing/shared";
import { Header, Link } from "@/shared/ui";

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
          Header("The requested resource could not be found");
        });

        Link(homeRoute, "Go to home page");
      });
    },
  });
};
