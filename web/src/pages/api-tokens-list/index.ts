import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { apiTokenCreateRoute, apiTokensRoute } from "@/shared/routing";
import { HeaderWithCreation } from "@/widgets/header-with-creation";
import { ApiTokensList } from "@/widgets/api-tokens-list";
import { i18n } from "@/shared/lib/i18n";

export const ApiTokensListPage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(apiTokensRoute);

    HeaderWithCreation(i18n("api_tokens.label"), apiTokenCreateRoute);

    h("div", () => {
      spec({
        classList: ["pt-3"],
      });

      ApiTokensList();
    });
  });
};
