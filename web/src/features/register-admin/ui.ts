import { Button, ErrorHint, Input } from "@/shared/ui";
import { $registrationError, registerAdminForm } from "./model";
import { h, spec } from "forest";

export const RegisterAdmin = () => {
  h("div", () => {
    spec({
      classList: ["mt-10", "max-w-sm", "w-full"],
    });

    h("form", () => {
      Input({
        type: "text",
        label: "Name",
        autofocus: true,
        inputChanged: registerAdminForm.fields.name.changed,
        errorText: registerAdminForm.fields.name.$errorText,
      });

      Input({
        type: "email",
        label: "Email",
        inputChanged: registerAdminForm.fields.email.changed,
        errorText: registerAdminForm.fields.email.$errorText,
      });

      Input({
        type: "password",
        label: "Password",
        inputChanged: registerAdminForm.fields.password.changed,
        errorText: registerAdminForm.fields.password.$errorText,
      });

      Input({
        type: "password",
        label: "Confirm password",
        required: true,
        inputChanged: registerAdminForm.fields.passwordConfirmation.changed,
        errorText: registerAdminForm.fields.passwordConfirmation.$errorText,
      });

      Button({
        text: "Register",
        event: registerAdminForm.submit,
        size: "base",
        prevent: true,
        variant: "default",
      });

      ErrorHint($registrationError, $registrationError.map(Boolean));
    });
  });
};
