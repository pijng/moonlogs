import { Button, ErrorHint, Input, Label, PlusIcon, Select, TrashIcon } from "@/shared/ui";
import { h, list, spec } from "forest";
import { $editError, deleteSchemaClicked, events, schemaForm } from "./model";
import { trigger } from "@/shared/lib";
import { combine, createEvent, sample } from "effector";
import { tagModel } from "@/entities/tag";

export const EditSchemaForm = () => {
  h("form", () => {
    Input({
      type: "text",
      label: "Title",
      value: schemaForm.fields.title.$value,
      inputChanged: schemaForm.fields.title.changed,
      errorText: schemaForm.fields.title.$errorText,
      hint: "Title - used for the human-readable name of the group in the web interface. Group search will also search for groups based on this characteristic",
    });

    Input({
      type: "text",
      label: "Description",
      value: schemaForm.fields.description.$value,
      inputChanged: schemaForm.fields.description.changed,
      errorText: schemaForm.fields.description.$errorText,
      hint: "Description - used for the human-readable description of group details in the web interface. Group search will also search for groups based on this characteristic",
    });

    h("div", () => {
      spec({ classList: ["mb-6"] });

      const $selectedTag = combine(tagModel.$tags, schemaForm.fields.tag_id.$value, (tags, id) => {
        return tags.find((t) => t.id === id)?.name || null;
      });

      Select({
        text: "Select a tag",
        value: $selectedTag,
        options: tagModel.$tags.map((tags) => tags.map((t) => t.name)),
        optionSelected: events.tagSelected,
      });
    });

    Input({
      type: "number",
      label: "Retention days",
      value: schemaForm.fields.retention_days.$value,
      inputChanged: schemaForm.fields.retention_days.changed,
      errorText: schemaForm.fields.retention_days.$errorText,
      hint: "Retention days - the number of days during which logs will be available after their creation. After the specified number of days elapses, the logs will be deleted. To set an infinite lifespan, specify 0 or leave the field empty",
    });

    h("div", () => {
      spec({ classList: ["relative", "flex", "items-center", "mb-4", "pt-4"] });

      Label({ text: "Group query fields", hint: "Group query fields - a set of fields by which log grouping will occur" });

      h("div", () => {
        spec({ classList: ["ml-1"] });

        Button({
          text: "",
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
            label: "Title",
            required: true,
            value: queryField.map((f) => f.title),
            inputChanged: titleChanged,
            errorText: schemaForm.fields.fields.$errorText,
            hint: "Title - used for the human-readable name of the field in the web interface for log filtering",
          });

          Input({
            type: "text",
            label: "Name",
            required: true,
            value: queryField.map((f) => f.name),
            inputChanged: nameChanged,
            errorText: schemaForm.fields.fields.$errorText,
            hint: "Name - used as a textual identifier for the group. Must be specified in Latin, in lowercase, and with underscores as separators",
          });

          Button({
            text: "",
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

      Label({ text: "Kinds", hint: "Kinds - a set of select options by which log grouping will occur" });

      h("div", () => {
        spec({ classList: ["ml-1"] });

        Button({
          text: "",
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
            label: "Title",
            required: true,
            value: kind.map((k) => k.title),
            inputChanged: titleChanged,
            errorText: schemaForm.fields.kinds.$errorText,
            hint: "Title - used for the human-readable name of the kind in the web interface for log filtering",
          });

          Input({
            type: "text",
            label: "Name",
            required: true,
            value: kind.map((k) => k.name),
            inputChanged: nameChanged,
            errorText: schemaForm.fields.fields.$errorText,
            hint: "Name - used as a textual identifier for the kind. Must be specified in Latin, in lowercase, and with underscores as separators",
          });

          Button({
            text: "",
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
        text: "Save",
        event: schemaForm.submit,
        size: "base",
        prevent: true,
        variant: "default",
      });

      Button({
        text: "Delete",
        event: deleteSchemaClicked,
        size: "base",
        prevent: true,
        variant: "delete",
      });
    });

    ErrorHint($editError, $editError.map(Boolean));
  });
};
