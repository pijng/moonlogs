import { Button, ErrorHint, Input } from "@/shared/ui";
import { h } from "forest";
import { $editError, schemaForm } from "./model";

export const EditSchemaForm = () => {
  h("form", () => {
    Input({
      value: schemaForm.fields.title.$value,
      type: "text",
      label: "Title",
      inputChanged: schemaForm.fields.title.changed,
      errorText: schemaForm.fields.title.$errorText,
      errorVisible: schemaForm.fields.title.$errors.map(Boolean),
    });

    Input({
      value: schemaForm.fields.description.$value,
      type: "text",
      label: "Description",
      inputChanged: schemaForm.fields.description.changed,
      errorText: schemaForm.fields.description.$errorText,
      errorVisible: schemaForm.fields.description.$errors.map(Boolean),
    });

    Button({
      text: "Save",
      event: schemaForm.submit,
      size: "base",
      prevent: true,
      variant: "default",
    });

    ErrorHint($editError, $editError.map(Boolean));
  });
};
