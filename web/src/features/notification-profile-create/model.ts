import { notificationProfileRoute } from "@/shared/routing";
import { createNotificationProfile, NotificationHeader, NotificationProfileToCreate } from "@/shared/api";
import { bindFieldList, manageSubmit, rules } from "@/shared/lib";
import { createEffect, createEvent, createStore, sample } from "effector";
import { createForm } from "effector-forms";

const addHeader = createEvent<NotificationHeader>();
const removeHeader = createEvent<number>();
const headerKeyChanged = createEvent<{ key: string; idx: number }>();
const headerValueChanged = createEvent<{ value: string; idx: number }>();
const ruleChecked = createEvent<number>();
const ruleUnchecked = createEvent<number>();

export const notificationProfileForm = createForm<NotificationProfileToCreate>({
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

export const $creationError = createStore("");

export const createNotificationProfileFx = createEffect((profile: NotificationProfileToCreate) => {
  return createNotificationProfile(profile);
});

manageSubmit({
  form: notificationProfileForm,
  actionFx: createNotificationProfileFx,
  error: $creationError,
  route: notificationProfileRoute,
});
