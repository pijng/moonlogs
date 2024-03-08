import { h } from "forest";
import { withRoute } from "atomic-router-forest";

import { homeRoute, schemaCreateRoute } from "@/shared/routing";
import { HeaderWithCreation, SchemasList } from "@/widgets";
import { i18n } from "@/shared/lib/i18n";

export const HomePage = () => {
  h("div", {
    fn() {
      // This allows to show/hide route if page is matched
      // It is required to call `withRoute` inside `h` call
      withRoute(homeRoute);

      HeaderWithCreation(i18n("log_groups.label"), schemaCreateRoute);

      SchemasList();
    },
  });
};
