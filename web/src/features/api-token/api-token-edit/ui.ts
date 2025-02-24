import { Button, ErrorHint, Input } from "@/shared/ui";
import { h, spec } from "forest";
import { $editError, apiTokenForm, deleteApiTokenClicked } from "./model";
import { i18n } from "@/shared/lib";

export const EditApiTokenForm = () => {
  h("form", () => {
    Input({
      type: "text",
      label: i18n("api_tokens.form.name.label"),
      value: apiTokenForm.fields.name.$value,
      inputChanged: apiTokenForm.fields.name.changed,
      errorText: apiTokenForm.fields.name.$errorText,
      hint: i18n("api_tokens.form.name.hint"),
    });

    Input({
      type: "checkbox",
      label: i18n("api_tokens.form.revoked"),
      value: apiTokenForm.fields.is_revoked.$value,
      inputChanged: apiTokenForm.fields.is_revoked.changed,
      errorText: apiTokenForm.fields.is_revoked.$errorText,
    });

    h("div", () => {
      spec({ classList: ["flex", "justify-start", "space-x-2"] });

      Button({
        text: i18n("buttons.save"),
        event: apiTokenForm.submit,
        size: "base",
        prevent: true,
        variant: "default",
      });

      Button({
        text: i18n("buttons.delete"),
        event: deleteApiTokenClicked,
        size: "base",
        prevent: true,
        variant: "delete",
      });
    });

    ErrorHint($editError, $editError.map(Boolean));
  });
};
