import { h, text } from "forest";
import { withRoute } from "atomic-router-forest";

import * as routes from "@/shared/routes";
import { CardLink, Link } from "@/shared/ui";

export const HomePage = () => {
  h("div", {
    classList: ["flex", "flex-col"],
    fn() {
      // This allows to show/hide route if page is matched
      // It is required to call `withRoute` inside `h` call
      withRoute(routes.home);

      text`Hello from the home page`;

      CardLink({
        title: "Procart",
        description:
          "Интеграция с RKeeper через модуль Procart от компании Carbis и я вообще норм сижу такой и туда сюда делаю без всяких преколов ну и все собсна пиздец карточка большая",
        route: routes.logsList,
        link: "logs",
      });

      Link(routes.logsList, "Show logs list");
    },
  });
};
