import { h, spec } from "forest";
import { withRoute } from "atomic-router-forest";

import { profileRoute } from "@/shared/routing";
import { Header } from "@/shared/ui";
import { UserProfile } from "@/widgets/user-profile";
import { i18n } from "@/shared/lib/i18n";

export const ProfilePage = () => {
  h("div", {
    fn() {
      // This allows to show/hide route if page is matched
      // It is required to call `withRoute` inside `h` call
      withRoute(profileRoute);

      h("div", () => {
        spec({ classList: ["max-w-3xl"] });

        Header(i18n("profile.label"));

        UserProfile();
      });
    },
  });
};
