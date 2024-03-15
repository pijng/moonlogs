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

    Input({
      type: "number",
      label: i18n("tags.form.view_order.label"),
      value: tagForm.fields.view_order.$value,
      inputChanged: tagForm.fields.view_order.changed,
      errorText: tagForm.fields.view_order.$errorText,
      hint: i18n("tags.form.view_order.hint"),
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
