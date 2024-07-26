import { Button, ErrorHint, Input } from "@/shared/ui";
import { $registrationError, registerAdminForm } from "./model";
import { h, spec } from "forest";
import { i18n } from "@/shared/lib";

export const RegisterAdmin = () => {
  h("div", () => {
    spec({
      classList: ["mt-10", "max-w-sm", "w-full"],
    });

    h("form", () => {
      Input({
        type: "text",
        label: i18n("members.form.name"),
        autofocus: true,
        inputChanged: registerAdminForm.fields.name.changed,
        errorText: registerAdminForm.fields.name.$errorText,
      });

      Input({
        type: "email",
        label: i18n("members.form.email"),
        inputChanged: registerAdminForm.fields.email.changed,
        errorText: registerAdminForm.fields.email.$errorText,
      });

      Input({
        type: "password",
        label: i18n("members.form.password"),
        inputChanged: registerAdminForm.fields.password.changed,
        errorText: registerAdminForm.fields.password.$errorText,
      });

      Input({
        type: "password",
        label: i18n("members.form.confirm_password"),
        required: true,
        inputChanged: registerAdminForm.fields.passwordConfirmation.changed,
        errorText: registerAdminForm.fields.passwordConfirmation.$errorText,
      });

      Button({
        text: i18n("buttons.register"),
        event: registerAdminForm.submit,
        size: "base",
        prevent: true,
        variant: "default",
      });

      ErrorHint($registrationError, $registrationError.map(Boolean));
    });
  });
};
