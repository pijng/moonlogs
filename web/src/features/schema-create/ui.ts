import { Button, ErrorHint, Input } from "@/shared/ui";
import { h } from "forest";
import { $creationError, schemaForm } from "./model";

export const NewSchemaForm = () => {
  h("form", () => {
    Input({
      type: "text",
      label: "Title",
      inputChanged: schemaForm.fields.title.changed,
      errorText: schemaForm.fields.title.$errorText,
      errorVisible: schemaForm.fields.title.$errors.map(Boolean),
    });

    Input({
      type: "text",
      label: "Description",
      inputChanged: schemaForm.fields.description.changed,
      errorText: schemaForm.fields.description.$errorText,
      errorVisible: schemaForm.fields.description.$errors.map(Boolean),
    });

    Input({
      type: "text",
      label: "Name",
      required: true,
      inputChanged: schemaForm.fields.name.changed,
      errorText: schemaForm.fields.name.$errorText,
      errorVisible: schemaForm.fields.name.$errors.map(Boolean),
    });

    Button({
      text: "Create",
      event: schemaForm.submit,
      size: "base",
      prevent: true,
      variant: "default",
    });

    ErrorHint($creationError, $creationError.map(Boolean));
  });
};
