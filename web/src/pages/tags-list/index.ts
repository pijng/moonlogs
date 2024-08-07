import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { tagCreateRoute, tagsRoute } from "@/shared/routing";
import { HeaderWithCreation } from "@/widgets/header-with-creation";
import { TagsList } from "@/widgets/tags-list";
import { i18n } from "@/shared/lib/i18n";

export const TagsListPage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(tagsRoute);

    HeaderWithCreation(i18n("tags.label"), tagCreateRoute);

    h("div", () => {
      spec({ classList: ["pt-3", "max-w-3xl"] });

      TagsList();
    });
  });
};
