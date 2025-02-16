import { notificationProfileModel } from "@/entities/notification-profile";
import {
  deleteNotificationProfile,
  editNotificationProfile,
  NotificationHeader,
  NotificationProfileToUpdate,
} from "@/shared/api";
import { bindFieldList, i18n, manageSubmit, rules } from "@/shared/lib";
import { notificationProfileRoute } from "@/shared/routing";
import { redirect } from "atomic-router";
import { attach, createEffect, createEvent, createStore, sample } from "effector";
import { createForm } from "effector-forms";

const addHeader = createEvent<NotificationHeader>();
const removeHeader = createEvent<number>();
const headerKeyChanged = createEvent<{ key: string; idx: number }>();
const headerValueChanged = createEvent<{ value: string; idx: number }>();
const ruleChecked = createEvent<number>();
const ruleUnchecked = createEvent<number>();

export const deleteNotificationProfileClicked = createEvent<number>();

export const notificationProfileForm = createForm<Omit<NotificationProfileToUpdate, "id">>({
  fields: {
    name: {
      init: "",
      rules: [rules.required()],
    },
    description: {
      init: "",
      rules: [rules.required()],
    },
    rule_ids: {
      init: [],
    },
    enabled: {
      init: true,
      rules: [],
    },
    silence_for: {
      init: "1m",
      rules: [rules.required()],
    },
    url: {
      init: "",
      rules: [rules.required()],
    },
    method: {
      init: "POST",
      rules: [rules.required()],
    },
    headers: {
      init: [],
      rules: [],
    },
    payload: {
      init: "",
      rules: [rules.required()],
    },
  },
  validateOn: ["submit"],
});

export const events = {
  addHeader,
  removeHeader,
  headerKeyChanged,
  headerValueChanged,
  ruleChecked,
  ruleUnchecked,
  deleteNotificationProfileClicked,
};

sample({
  source: notificationProfileForm.fields.headers.$value,
  clock: addHeader,
  fn: (headers) => {
    const newHeader: NotificationHeader = { key: "", value: "" };
    return [...headers, newHeader];
  },
  target: notificationProfileForm.fields.headers.onChange,
});

sample({
  source: notificationProfileForm.fields.headers.$value,
  clock: removeHeader,
  fn: (headers, idx) => headers.filter((_, index) => index !== idx),
  target: notificationProfileForm.fields.headers.onChange,
});

sample({
  source: notificationProfileForm.fields.headers.$value,
  clock: headerKeyChanged,
  fn: (headers, { key, idx }) => {
    return headers.map((header, index) => (index === idx ? { ...header, key: key } : header));
  },
  target: notificationProfileForm.fields.headers.onChange,
});

sample({
  source: notificationProfileForm.fields.headers.$value,
  clock: headerValueChanged,
  fn: (headers, { value, idx }) => {
    return headers.map((header, index) => (index === idx ? { ...header, value: value } : header));
  },
  target: notificationProfileForm.fields.headers.onChange,
});

bindFieldList({ field: notificationProfileForm.fields.rule_ids, added: ruleChecked, removed: ruleUnchecked });

export const $editError = createStore("");

export const editNotificationProfileFx = createEffect((rule: NotificationProfileToUpdate) => {
  return editNotificationProfile(rule);
});

sample({
  source: notificationProfileModel.$currentNotificationProfile,
  target: [notificationProfileForm.setForm],
});

manageSubmit({
  form: notificationProfileForm,
  actionFx: editNotificationProfileFx,
  error: $editError,
  currentModel: notificationProfileModel.$currentNotificationProfile,
  route: notificationProfileRoute,
});

export const deleteNotificationProfileFx = createEffect((id: number) => {
  deleteNotificationProfile(id);
});

const alertDeleteFx = attach({
  source: i18n("notification_profiles.alerts.delete"),
  effect(alertText, id: number) {
    const confirmed = confirm(alertText);

    return confirmed ? { confirmed: true, id: id } : { confirmed: false, id: id };
  },
});

sample({
  clock: deleteNotificationProfileClicked,
  target: alertDeleteFx,
});

sample({
  source: notificationProfileModel.$currentNotificationProfile,
  clock: alertDeleteFx.doneData,
  filter: (_, { confirmed }) => confirmed,
  fn: ({ id }) => id,
  target: deleteNotificationProfileFx,
});

redirect({
  clock: deleteNotificationProfileFx.done,
  route: notificationProfileRoute,
});
