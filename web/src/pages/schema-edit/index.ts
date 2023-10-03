import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { schemaEditRoute } from "@/routing/shared";
import { Header } from "@/shared/ui";
import { EditSchemaForm } from "@/features/schema-edit";

export const SchemaEditPage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(schemaEditRoute);

    Header("Edit category");

    h("div", () => {
      spec({
        classList: ["pt-5", "max-w-2xl"],
      });

      EditSchemaForm();
    });
  });
};
