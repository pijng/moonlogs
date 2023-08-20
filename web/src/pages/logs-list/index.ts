import { h, text } from "forest";
import { withRoute } from "atomic-router-forest";

import * as routes from "@/shared/routes";
import { Link } from "@/shared/ui";

export const LogsListPage = () => {
  h("div", {
    classList: ["flex", "flex-col"],
    fn() {
      // This allows to show/hide route if page is matched
      // It is required to call `withRoute` inside `h` call
      withRoute(routes.logsList);

      text`Hello from the logs list page`;

      Link(routes.home, "Go to home");
    },
  });
};
