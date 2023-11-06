import { Button, ErrorHint, Input } from "@/shared/ui";
import { $authError, authForm } from "./model";
import { h, spec } from "forest";

export const Auth = () => {
  h("div", () => {
    spec({
      classList: ["mt-10", "max-w-sm", "w-full"],
    });

    h("form", () => {
      Input({
        type: "email",
        label: "Email",
        inputChanged: authForm.fields.email.changed,
        errorText: authForm.fields.email.$errorText,
        errorVisible: authForm.fields.email.$errors.map(Boolean),
      });

      Input({
        type: "password",
        label: "Password",
        inputChanged: authForm.fields.password.changed,
        errorText: authForm.fields.password.$errorText,
        errorVisible: authForm.fields.password.$errors.map(Boolean),
      });

      Button({
        text: "Log in",
        event: authForm.submit,
        size: "base",
        prevent: true,
        variant: "default",
      });

      ErrorHint($authError, $authError.map(Boolean));
    });
  });
};
