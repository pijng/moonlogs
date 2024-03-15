import { Button, ErrorHint, Input } from "@/shared/ui";
import { h, spec } from "forest";
import { $editError, tagForm, deleteTagClicked } from "./model";
import { i18n } from "@/shared/lib/i18n";

export const EditTagForm = () => {
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

    h("div", () => {
      spec({ classList: ["flex", "justify-start", "space-x-2"] });

      Button({
        text: i18n("buttons.save"),
        event: tagForm.submit,
        size: "base",
        prevent: true,
        variant: "default",
      });

      Button({
        text: i18n("buttons.delete"),
        event: deleteTagClicked,
        size: "base",
        prevent: true,
        variant: "delete",
      });
    });

    ErrorHint($editError, $editError.map(Boolean));
  });
};
