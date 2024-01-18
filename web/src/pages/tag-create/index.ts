import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { tagCreateRoute } from "@/routing/shared";
import { Header } from "@/shared/ui";
import { NewTagForm } from "@/features";

export const TagCreatePage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(tagCreateRoute);

    Header("Create tag");

    h("div", () => {
      spec({
        classList: ["pt-5", "max-w-2xl"],
      });

      NewTagForm();
    });
  });
};