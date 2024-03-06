import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { apiTokenEditRoute } from "@/routing/shared";
import { Header } from "@/shared/ui";
import { EditApiTokenForm } from "@/features";
import { i18n } from "@/shared/lib/i18n";

export const ApiTokenEditPage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(apiTokenEditRoute);

    Header(i18n("api_tokens.form.actions.edit"));

    h("div", () => {
      spec({
        classList: ["pt-5", "max-w-2xl"],
      });

      EditApiTokenForm();
    });
  });
};
