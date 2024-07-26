import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { apiTokenCreateRoute } from "@/shared/routing";
import { Header } from "@/shared/ui";
import { NewApiTokenForm } from "@/features/api-token-create";
import { i18n } from "@/shared/lib/i18n";

export const ApiTokenCreatePage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(apiTokenCreateRoute);

    Header(i18n("api_tokens.form.actions.create"));

    h("div", () => {
      spec({
        classList: ["pt-5", "max-w-2xl"],
      });

      NewApiTokenForm();
    });
  });
};
