import { Button, ErrorHint, Input, Label, PlusIcon, Select, TrashIcon } from "@/shared/ui";
import { h, list, spec } from "forest";
import { $editError, deleteSchemaClicked, events, schemaForm } from "./model";
import { trigger, i18n } from "@/shared/lib";
import { createEvent, sample } from "effector";
import { tagModel } from "@/entities/tag";

export const EditSchemaForm = () => {
  h("form", () => {
    Input({
      type: "text",
      label: i18n("log_groups.form.title.label"),
      value: schemaForm.fields.title.$value,
      inputChanged: schemaForm.fields.title.changed,
      errorText: schemaForm.fields.title.$errorText,
      hint: i18n("log_groups.form.title.hint"),
    });

    Input({
      type: "text",
      label: i18n("log_groups.form.description.label"),
      value: schemaForm.fields.description.$value,
      inputChanged: schemaForm.fields.description.changed,
      errorText: schemaForm.fields.description.$errorText,
      hint: i18n("log_groups.form.description.hint"),
    });

    h("div", () => {
      spec({ classList: ["mb-6"] });

      Select({
        text: i18n("log_groups.form.tag.label"),
        value: schemaForm.fields.tag_id.$value,
        options: tagModel.$tags,
        optionSelected: events.tagSelected,
      });
    });

    Input({
      type: "number",
      label: i18n("log_groups.form.retention_days.label"),
      value: schemaForm.fields.retention_days.$value,
      inputChanged: schemaForm.fields.retention_days.changed,
      errorText: schemaForm.fields.retention_days.$errorText,
      hint: i18n("log_groups.form.retention_days.hint"),
    });

    h("div", () => {
      spec({ classList: ["relative", "flex", "items-center", "mb-4", "pt-4"] });

      Label({
        text: i18n("log_groups.form.group_query_fields.label"),
        hint: i18n("log_groups.form.group_query_fields.hint"),
      });

      h("div", () => {
        spec({ classList: ["ml-1"] });

        Button({
          variant: "default",
          prevent: true,
          style: "round",
          size: "extra_small",
          event: events.addField,
          preIcon: PlusIcon,
        });
      });
    });

    h("div", () => {
      list(schemaForm.fields.fields.$value, ({ store: queryField, key: idx }) => {
        const titleChanged = createEvent<string>();
        const nameChanged = createEvent<string>();

        sample({
          source: idx,
          clock: titleChanged,
          fn: (idx, title) => ({ title: title, idx: idx }),
          target: events.fieldTitleChanged,
        });

        sample({
          source: idx,
          clock: nameChanged,
          fn: (idx, name) => ({ name: name, idx: idx }),
          target: events.fieldNameChanged,
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
            label: i18n("log_groups.form.group_query_fields.fields.title.label"),
            required: true,
            value: queryField.map((f) => f.title),
            inputChanged: titleChanged,
            errorText: schemaForm.fields.fields.$errorText,
            hint: i18n("log_groups.form.group_query_fields.fields.title.hint"),
          });

          Input({
            type: "text",
            label: i18n("log_groups.form.group_query_fields.fields.name.label"),
            required: true,
            value: queryField.map((f) => f.name),
            inputChanged: nameChanged,
            errorText: schemaForm.fields.fields.$errorText,
            hint: i18n("log_groups.form.group_query_fields.fields.name.hint"),
          });

          Button({
            event: trigger({ source: idx, target: events.removeField }),
            size: "plain",
            prevent: true,
            variant: "delete_icon",
            preIcon: TrashIcon,
          });
        });
      });
    });

    h("div", () => {
      spec({ classList: ["relative", "flex", "items-center", "mb-4", "pt-4"] });

      Label({
        text: i18n("log_groups.form.kinds.label"),
        hint: i18n("log_groups.form.kinds.hint"),
      });

      h("div", () => {
        spec({ classList: ["ml-1"] });

        Button({
          variant: "default",
          prevent: true,
          style: "round",
          size: "extra_small",
          event: events.addKind,
          preIcon: PlusIcon,
        });
      });
    });

    h("div", () => {
      list(schemaForm.fields.kinds.$value, ({ store: kind, key: idx }) => {
        const titleChanged = createEvent<string>();
        const nameChanged = createEvent<string>();

        sample({
          source: idx,
          clock: titleChanged,
          fn: (idx, title) => ({ title: title, idx: idx }),
          target: events.kindTitleChanged,
        });

        sample({
          source: idx,
          clock: nameChanged,
          fn: (idx, name) => ({ name: name, idx: idx }),
          target: events.kindNameChanged,
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
            label: i18n("log_groups.form.kinds.fields.title.label"),
            required: true,
            value: kind.map((k) => k.title),
            inputChanged: titleChanged,
            errorText: schemaForm.fields.kinds.$errorText,
            hint: i18n("log_groups.form.kinds.fields.title.hint"),
          });

          Input({
            type: "text",
            label: i18n("log_groups.form.kinds.fields.name.label"),
            required: true,
            value: kind.map((k) => k.name),
            inputChanged: nameChanged,
            errorText: schemaForm.fields.fields.$errorText,
            hint: i18n("log_groups.form.kinds.fields.name.hint"),
          });

          Button({
            event: trigger({ source: idx, target: events.removeKind }),
            size: "plain",
            prevent: true,
            variant: "delete_icon",
            preIcon: TrashIcon,
          });
        });
      });
    });

    h("div", () => {
      spec({ classList: ["flex", "justify-start", "space-x-2", "pt-4"] });

      Button({
        text: i18n("buttons.save"),
        event: schemaForm.submit,
        size: "base",
        prevent: true,
        variant: "default",
      });

      Button({
        text: i18n("buttons.delete"),
        event: deleteSchemaClicked,
        size: "base",
        prevent: true,
        variant: "delete",
      });
    });

    ErrorHint($editError, $editError.map(Boolean));
  });
};
