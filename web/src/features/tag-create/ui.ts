import { Button, ErrorHint, Input } from "@/shared/ui";
import { h } from "forest";
import { $creationError, tagForm } from "./model";

export const NewTagForm = () => {
  h("form", () => {
    Input({
      type: "text",
      label: "Name",
      value: tagForm.fields.name.$value,
      inputChanged: tagForm.fields.name.changed,
      errorText: tagForm.fields.name.$errorText,
      hint: "Name - used for the human-readable name of the tag in the web interface",
    });

    Button({
      text: "Create",
      event: tagForm.submit,
      size: "base",
      prevent: true,
      variant: "default",
    });

    ErrorHint($creationError, $creationError.map(Boolean));
  });
};
