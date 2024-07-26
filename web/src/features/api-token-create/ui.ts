import { Button, ErrorHint, Input } from "@/shared/ui";
import { h, spec } from "forest";
import { $creationError, $freshToken, apiTokenForm } from "./model";
import { createEvent } from "effector";
import { i18n } from "@/shared/lib";

export const NewApiTokenForm = () => {
  h("form", () => {
    h("div", () => {
      spec({
        visible: $freshToken.map((t) => !t),
      });

      Input({
        type: "text",
        label: i18n("api_tokens.form.name.label"),
        value: apiTokenForm.fields.name.$value,
        inputChanged: apiTokenForm.fields.name.changed,
        errorText: apiTokenForm.fields.name.$errorText,
        hint: i18n("api_tokens.form.name.hint"),
      });

      Button({
        text: i18n("buttons.create"),
        event: apiTokenForm.submit,
        size: "base",
        prevent: true,
        variant: "default",
      });

      ErrorHint($creationError, $creationError.map(Boolean));
    });

    Input({
      type: "text",
      label: i18n("api_tokens.form.creation_hint"),
      value: $freshToken,
      visible: $freshToken.map(Boolean),
      inputChanged: createEvent(),
    });
  });
};
