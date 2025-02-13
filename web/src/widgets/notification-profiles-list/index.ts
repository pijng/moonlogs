import { notificationProfileModel } from "@/entities/notification-profile";
import { notificationProfileEditRoute } from "@/shared/routing";
import { NotificationProfilesTable } from "@/shared/ui";
import { createEvent, sample } from "effector";
import { h, spec } from "forest";

export const NotificationProfilesList = () => {
  h("div", () => {
    spec({
      classList: ["flex", "flex-col", "space-y-6", "mt-3"],
    });

    const editNotificationProfileClicked = createEvent<number>();
    sample({
      source: editNotificationProfileClicked,
      fn: (id) => ({ id }),
      target: notificationProfileEditRoute.open,
    });

    NotificationProfilesTable(notificationProfileModel.$notificationProfiles, editNotificationProfileClicked);
  });
};
