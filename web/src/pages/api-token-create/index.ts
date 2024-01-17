import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { apiTokenCreateRoute } from "@/routing/shared";
import { Header } from "@/shared/ui";
import { NewApiTokenForm } from "@/features";

export const ApiTokenCreatePage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(apiTokenCreateRoute);

    Header("Create API token");

    h("div", () => {
      spec({
        classList: ["pt-5", "max-w-2xl"],
      });

      NewApiTokenForm();
    });
  });
};
