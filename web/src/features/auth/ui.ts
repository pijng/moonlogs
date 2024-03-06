import { Button, ErrorHint, Input } from "@/shared/ui";
import { $authError, authForm } from "./model";
import { h, spec } from "forest";
import { i18n } from "@/shared/lib/i18n";

export const Auth = () => {
  h("div", () => {
    spec({
      classList: ["mt-10", "max-w-sm", "w-full"],
    });

    h("form", () => {
      Input({
        type: "email",
        label: i18n("auth.email"),
        autofocus: true,
        inputChanged: authForm.fields.email.changed,
        errorText: authForm.fields.email.$errorText,
      });

      Input({
        type: "password",
        label: i18n("auth.password"),
        inputChanged: authForm.fields.password.changed,
        errorText: authForm.fields.password.$errorText,
      });

      Button({
        text: i18n("buttons.log_in"),
        event: authForm.submit,
        size: "base",
        prevent: true,
        variant: "default",
      });

      ErrorHint($authError, $authError.map(Boolean));
    });
  });
};
