import { Button, ErrorHint, Input, Label, PlusIcon, TrashIcon } from "@/shared/ui";
import { h, list, spec } from "forest";
import { $editError, deleteSchemaClicked, events, schemaForm } from "./model";
import { trigger } from "@/shared/lib";
import { createEvent, sample } from "effector";

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

    Input({
      type: "number",
      label: "Retention days",
      value: schemaForm.fields.retention_time.$value,
      inputChanged: schemaForm.fields.retention_time.changed,
      errorText: schemaForm.fields.retention_time.$errorText,
      hint: "Retention days - the number of days during which logs will be available after their creation. After the specified number of days elapses, the logs will be deleted. To set an infinite lifespan, specify 0 or leave the field empty",
    });

    h("div", () => {
      spec({ classList: ["relative", "flex", "items-center", "mb-6"] });

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
      spec({ classList: ["flex", "justify-start", "space-x-2"] });

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
