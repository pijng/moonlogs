import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { tagEditRoute } from "@/shared/routing";
import { Header } from "@/shared/ui";
import { EditTagForm } from "@/features";
import { i18n } from "@/shared/lib/i18n";

export const TagEditPage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(tagEditRoute);

    Header(i18n("tags.form.actions.edit"));

    h("div", () => {
      spec({
        classList: ["pt-5", "max-w-2xl"],
      });

      EditTagForm();
    });
  });
};
