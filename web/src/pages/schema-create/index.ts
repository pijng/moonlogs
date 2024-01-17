import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { schemaCreateRoute } from "@/routing/shared";
import { Header } from "@/shared/ui";
import { NewSchemaForm } from "@/features/schema-create";

export const SchemaCreatePage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(schemaCreateRoute);

    Header("Create log group");

    h("div", () => {
      spec({
        classList: ["pt-5", "max-w-3xl"],
      });

      NewSchemaForm();
    });
  });
};
