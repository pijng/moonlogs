import { h } from "forest";
import { withRoute } from "atomic-router-forest";

import { homeRoute, schemaCreateRoute } from "@/routing/shared";
import { HeaderWithCreation, SchemasList } from "@/widgets";

export const HomePage = () => {
  h("div", {
    fn() {
      // This allows to show/hide route if page is matched
      // It is required to call `withRoute` inside `h` call
      withRoute(homeRoute);

      HeaderWithCreation("Log groups", schemaCreateRoute);

      SchemasList();
    },
  });
};
