import { withRoute } from "atomic-router-forest";
import { h, spec } from "forest";

import { notificationProfileEditRoute } from "@/shared/routing";

import { i18n } from "@/shared/lib/i18n";
import { Header } from "@/shared/ui";
import { EditNotificationProfileForm } from "@/features/notification-profile-edit";

export const NotificationProfleEditPage = () => {
  h("div", () => {
    // This allows to show/hide route if page is matched
    // It is required to call `withRoute` inside `h` call
    withRoute(notificationProfileEditRoute);

    Header(i18n("notification_profiles.form.actions.edit"));

    h("div", () => {
      spec({
        classList: ["pt-5", "max-w-2xl"],
      });

      EditNotificationProfileForm();
    });
  });
};
