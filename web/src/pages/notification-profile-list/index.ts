import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { notificationProfileCreateRoute, notificationProfileRoute } from "@/shared/routing";
import { HeaderWithCreation } from "@/widgets/header-with-creation";

import { i18n } from "@/shared/lib/i18n";
import { NotificationProfilesList } from "@/widgets/notification-profiles-list";

export const NotificationProfileListPage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(notificationProfileRoute);

    HeaderWithCreation(i18n("notification_profiles.label"), notificationProfileCreateRoute);

    h("div", () => {
      spec({ classList: ["pt-3"] });

      NotificationProfilesList();
    });
  });
};
