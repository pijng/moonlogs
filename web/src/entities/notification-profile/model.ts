import { getNotificationProfile, getNotificationProfiles, NotificationProfile } from "@/shared/api";
import { createEffect, createStore } from "effector";

const getNotificationProfilesFx = createEffect(() => {
  return getNotificationProfiles();
});

const getNotificationProfileFx = createEffect((id: number) => {
  return getNotificationProfile(id);
});

export const $notificationProfiles = createStore<NotificationProfile[]>([]).on(
  getNotificationProfilesFx.doneData,
  (_, notificationProfilesResponse) => notificationProfilesResponse.data,
);

export const $currentNotificationProfile = createStore<NotificationProfile>({
  id: 0,
  name: "",
  description: "",
  enabled: true,
  silence_for: "1m",
  rule_ids: [],
  url: "",
  method: "POST",
  headers: [],
  payload: "",
}).on(getNotificationProfileFx.doneData, (_, notificationProfileResponse) => notificationProfileResponse.data);

export const effects = {
  getNotificationProfilesFx,
  getNotificationProfileFx,
};

export const events = {};
