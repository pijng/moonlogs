import { Button, ErrorHint, Input } from "@/shared/ui";
import { h, spec } from "forest";
import { $editError, apiTokenForm, deleteApiTokenClicked } from "./model";

export const EditApiTokenForm = () => {
  h("form", () => {
    Input({
      type: "text",
      label: "Name",
      value: apiTokenForm.fields.name.$value,
      inputChanged: apiTokenForm.fields.name.changed,
      errorText: apiTokenForm.fields.name.$errorText,
      hint: "Name - is used to indicate which service will use this API token. It does not affect the token functionally",
    });

    Input({
      type: "checkbox",
      label: "Revoked",
      value: apiTokenForm.fields.is_revoked.$value,
      inputChanged: apiTokenForm.fields.is_revoked.changed,
      errorText: apiTokenForm.fields.is_revoked.$errorText,
    });

    h("div", () => {
      spec({ classList: ["flex", "justify-start", "space-x-2"] });

      Button({
        text: "Save",
        event: apiTokenForm.submit,
        size: "base",
        prevent: true,
        variant: "default",
      });

      Button({
        text: "Delete",
        event: deleteApiTokenClicked,
        size: "base",
        prevent: true,
        variant: "delete",
      });
    });

    ErrorHint($editError, $editError.map(Boolean));
  });
};
