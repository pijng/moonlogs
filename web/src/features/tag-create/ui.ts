import { Button, ErrorHint, Input } from "@/shared/ui";
import { h } from "forest";
import { $creationError, tagForm } from "./model";
import { i18n } from "@/shared/lib/i18n";

export const NewTagForm = () => {
  h("form", () => {
    Input({
      type: "text",
      label: i18n("tags.form.name.label"),
      value: tagForm.fields.name.$value,
      inputChanged: tagForm.fields.name.changed,
      errorText: tagForm.fields.name.$errorText,
      hint: i18n("tags.form.name.hint"),
    });

    Button({
      text: i18n("buttons.create"),
      event: tagForm.submit,
      size: "base",
      prevent: true,
      variant: "default",
    });

    ErrorHint($creationError, $creationError.map(Boolean));
  });
};
