import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { schemaCreateRoute } from "@/shared/routing";
import { Header } from "@/shared/ui";
import { NewSchemaForm } from "@/features/schema/schema-create";
import { i18n } from "@/shared/lib/i18n";

export const SchemaCreatePage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(schemaCreateRoute);

    Header(i18n("log_groups.form.actions.create"));

    h("div", () => {
      spec({
        classList: ["pt-5", "max-w-3xl"],
      });

      NewSchemaForm();
    });
  });
};
