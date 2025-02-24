import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { notificationProfileCreateRoute } from "@/shared/routing";
import { Header } from "@/shared/ui";
import { i18n } from "@/shared/lib/i18n";
import { NewNotificationProfileForm } from "@/features/notification-profile/notification-profile-create";

export const NotificationProfileCreatePage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(notificationProfileCreateRoute);

    Header(i18n("notification_profiles.form.actions.create"));

    h("div", () => {
      spec({
        classList: ["pt-5", "max-w-2xl", "mb-40"],
      });

      NewNotificationProfileForm();
    });
  });
};
