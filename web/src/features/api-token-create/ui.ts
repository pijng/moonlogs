import { Button, ErrorHint, Input } from "@/shared/ui";
import { h, spec } from "forest";
import { $creationError, $freshToken, apiTokenForm } from "./model";
import { createEvent } from "effector";

export const NewApiTokenForm = () => {
  h("form", () => {
    h("div", () => {
      spec({
        visible: $freshToken.map((t) => !t),
      });

      Input({
        type: "text",
        label: "Name",
        value: apiTokenForm.fields.name.$value,
        inputChanged: apiTokenForm.fields.name.changed,
        errorText: apiTokenForm.fields.name.$errorText,
        hint: "Name - is used to indicate which service will use this API token. It does not affect the token functionally",
      });

      Button({
        text: "Create",
        event: apiTokenForm.submit,
        size: "base",
        prevent: true,
        variant: "default",
      });

      ErrorHint($creationError, $creationError.map(Boolean));
    });

    Input({
      type: "text",
      label:
        "Your integration API token has been successfully created. Make sure to save it securely now, as it won't be displayed again for security reasons",
      value: $freshToken,
      visible: $freshToken.map(Boolean),
      inputChanged: createEvent(),
    });
  });
};
