import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { apiTokenCreateRoute, apiTokensRoute } from "@/routing/shared";
import { HeaderWithCreation } from "@/widgets";
import { ApiTokensList } from "@/widgets";

export const ApiTokensListPage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(apiTokensRoute);

    HeaderWithCreation("API tokens", apiTokenCreateRoute);

    h("div", () => {
      spec({
        classList: ["pt-3"],
      });

      ApiTokensList();
    });
  });
};
