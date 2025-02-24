import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { schemaEditRoute } from "@/shared/routing";
import { Header } from "@/shared/ui";
import { EditSchemaForm } from "@/features/schema/schema-edit";
import { i18n } from "@/shared/lib/i18n";

export const SchemaEditPage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(schemaEditRoute);

    Header(i18n("log_groups.form.actions.edit"));

    h("div", () => {
      spec({
        classList: ["pt-5", "max-w-2xl"],
      });

      EditSchemaForm();
    });
  });
};
