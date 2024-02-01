import { h, spec } from "forest";
import { withRoute } from "atomic-router-forest";

import { profileRoute } from "@/routing/shared";
import { Header } from "@/shared/ui";
import { UserProfile } from "@/widgets";

export const ProfilePage = () => {
  h("div", {
    fn() {
      // This allows to show/hide route if page is matched
      // It is required to call `withRoute` inside `h` call
      withRoute(profileRoute);

      h("div", () => {
        spec({ classList: ["max-w-3xl"] });

        Header("Profile");

        UserProfile();
      });
    },
  });
};
