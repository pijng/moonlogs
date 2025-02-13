import { Button, ErrorHint, Input, Label, Multiselect, PlusIcon, Select, TextArea, TrashIcon } from "@/shared/ui";
import { h, list, remap, spec } from "forest";
import { $creationError, events, notificationProfileForm } from "./model";
import { i18n, trigger } from "@/shared/lib";
import { NotificationProfileToCreate } from "@/shared/api";
import { createEvent, createStore, sample } from "effector";
import { alertingRuleModel } from "@/entities/alerting-rule";

export const NewNotificationProfileForm = () => {
  h("form", () => {
    Input({
      type: "text",
      label: i18n("notification_profiles.form.name.label"),
      hint: i18n("notification_profiles.form.name.hint"),
      value: notificationProfileForm.fields.name.$value,
      inputChanged: notificationProfileForm.fields.name.changed,
      errorText: notificationProfileForm.fields.name.$errorText,
    });

    Input({
      type: "text",
      label: i18n("notification_profiles.form.description.label"),
      hint: i18n("notification_profiles.form.name.hint"),
      value: notificationProfileForm.fields.description.$value,
      inputChanged: notificationProfileForm.fields.description.changed,
      errorText: notificationProfileForm.fields.description.$errorText,
    });

    h("div", () => {
      spec({ classList: ["mb-6"] });

      Multiselect({
        text: i18n("notification_profiles.form.rule_name.label"),
        hint: i18n("notification_profiles.form.rule_name.hint"),
        options: alertingRuleModel.$alertingRules.map((rule) => rule.map((s) => ({ name: s.name, id: s.id }))),
        selectedOptions: notificationProfileForm.fields.rule_ids.$value,
        optionChecked: events.ruleChecked,
        optionUnchecked: events.ruleUnchecked,
      });
    });

    Input({
      type: "checkbox",
      label: i18n("notification_profiles.form.enabled.label"),
      value: notificationProfileForm.fields.enabled.$value,
      inputChanged: notificationProfileForm.fields.enabled.changed,
      errorText: notificationProfileForm.fields.enabled.$errorText,
    });

    Input({
      type: "text",
      label: i18n("notification_profiles.form.url.label"),
      hint: i18n("notification_profiles.form.url.hint"),
      required: true,
      value: notificationProfileForm.fields.url.$value,
      inputChanged: notificationProfileForm.fields.url.changed,
      errorText: notificationProfileForm.fields.url.$errorText,
    });

    h("div", () => {
      spec({ classList: ["mb-6"] });

      Select({
        text: i18n("notification_profiles.form.method.label"),
        hint: i18n("notification_profiles.form.method.hint"),
        value: notificationProfileForm.fields.method.$value,
        options: createStore<NotificationProfileToCreate["method"][]>(["POST", "GET"]),
        optionSelected: notificationProfileForm.fields.method.onChange,
        withBlank: createStore(false),
      });
    });

    h("div", () => {
      spec({ classList: ["relative", "flex", "items-center", "mb-4", "pt-2"] });

      Label({
        text: i18n("notification_profiles.form.headers.label"),
        hint: i18n("notification_profiles.form.headers.hint"),
      });

      h("div", () => {
        spec({ classList: ["ml-1"] });

        Button({
          variant: "default",
          prevent: true,
          style: "round",
          size: "extra_small",
          event: events.addHeader,
          preIcon: PlusIcon,
        });
      });
    });

    h("div", () => {
      list(notificationProfileForm.fields.headers.$value, ({ store: header, key: idx }) => {
        const keyChanged = createEvent<string>();
        const valueChanged = createEvent<string>();

        sample({
          source: idx,
          clock: keyChanged,
          fn: (idx, key) => ({ key: key, idx: idx }),
          target: events.headerKeyChanged,
        });

        sample({
          source: idx,
          clock: valueChanged,
          fn: (idx, value) => ({ value: value, idx: idx }),
          target: events.headerValueChanged,
        });

        h("div", () => {
          spec({
            classList: ["grid", "gap-3", "place-items-stretch"],
            style: {
              gridTemplateColumns: "14fr 14fr 1fr",
            },
          });

          Input({
            type: "text",
            label: i18n("notification_profiles.form.headers.fields.key.label"),
            hint: i18n("notification_profiles.form.headers.fields.key.hint"),
            required: true,
            value: remap(header, "key"),
            inputChanged: keyChanged,
            errorText: notificationProfileForm.fields.url.$errorText,
          });

          Input({
            type: "text",
            label: i18n("notification_profiles.form.headers.fields.value.label"),
            hint: i18n("notification_profiles.form.headers.fields.value.hint"),
            required: true,
            value: remap(header, "value"),
            inputChanged: valueChanged,
            errorText: notificationProfileForm.fields.url.$errorText,
          });

          Button({
            event: trigger({ source: idx, target: events.removeHeader }),
            size: "plain",
            prevent: true,
            variant: "delete_icon",
            preIcon: TrashIcon,
          });
        });
      });
    });

    TextArea({
      type: "text",
      label: i18n("notification_profiles.form.payload.label"),
      hint: i18n("notification_profiles.form.payload.hint"),
      required: true,
      rows: 10,
      autoHeight: true,
      value: notificationProfileForm.fields.payload.$value,
      inputChanged: notificationProfileForm.fields.payload.changed,
      errorText: notificationProfileForm.fields.payload.$errorText,
    });

    h("div", () => {
      spec({ classList: ["pt-4"] });
      Button({
        text: i18n("buttons.create"),
        event: notificationProfileForm.submit,
        size: "base",
        prevent: true,
        variant: "default",
      });
    });

    ErrorHint($creationError, $creationError.map(Boolean));
  });
};
