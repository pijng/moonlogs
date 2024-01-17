import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { tagCreateRoute, tagsRoute } from "@/routing/shared";
import { HeaderWithCreation } from "@/widgets";
import { TagsList } from "@/widgets";

export const TagsListPage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(tagsRoute);

    HeaderWithCreation("Tags", tagCreateRoute);

    h("div", () => {
      spec({
        classList: ["pt-3"],
      });

      TagsList();
    });
  });
};
