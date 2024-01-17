import { Button, ErrorHint, Input } from "@/shared/ui";
import { h, spec } from "forest";
import { $editError, tagForm, deleteTagClicked } from "./model";

export const EditTagForm = () => {
  h("form", () => {
    Input({
      type: "text",
      label: "Name",
      value: tagForm.fields.name.$value,
      inputChanged: tagForm.fields.name.changed,
      errorText: tagForm.fields.name.$errorText,
      hint: "Name - used for the human-readable name of the tag in the web interface",
    });

    h("div", () => {
      spec({ classList: ["flex", "justify-start", "space-x-2"] });

      Button({
        text: "Save",
        event: tagForm.submit,
        size: "base",
        prevent: true,
        variant: "default",
      });

      Button({
        text: "Delete",
        event: deleteTagClicked,
        size: "base",
        prevent: true,
        variant: "delete",
      });
    });

    ErrorHint($editError, $editError.map(Boolean));
  });
};
