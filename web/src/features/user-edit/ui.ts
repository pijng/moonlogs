import { Button, ErrorHint, Input, Select } from "@/shared/ui";
import { h } from "forest";
import { $editError, memberForm } from "./model";
import { createStore } from "effector";
import { UserRole } from "@/shared/api/users";

export const EditMemberForm = () => {
  h("form", () => {
    Input({
      value: memberForm.fields.name.$value,
      type: "text",
      label: "Name",
      inputChanged: memberForm.fields.name.changed,
      errorText: memberForm.fields.name.$errorText,
      errorVisible: memberForm.fields.name.$errors.map(Boolean),
    });

    Input({
      value: memberForm.fields.email.$value,
      type: "email",
      label: "Email",
      inputChanged: memberForm.fields.email.changed,
      errorText: memberForm.fields.email.$errorText,
      errorVisible: memberForm.fields.email.$errors.map(Boolean),
    });

    Select({
      value: memberForm.fields.role.$value.map(String),
      id: "role",
      text: "Select a role",
      options: createStore<UserRole[]>(["Member", "Admin"]),
      optionSelected: memberForm.fields.role.changed,
    });

    Input({
      value: memberForm.fields.password.$value,
      type: "password",
      label: "Password",
      required: true,
      inputChanged: memberForm.fields.password.changed,
      errorText: memberForm.fields.password.$errorText,
      errorVisible: memberForm.fields.password.$errors.map(Boolean),
    });

    Input({
      value: memberForm.fields.passwordConfirmation.$value,
      type: "password",
      label: "Confirm password",
      required: true,
      inputChanged: memberForm.fields.passwordConfirmation.changed,
      errorText: memberForm.fields.passwordConfirmation.$errorText,
      errorVisible: memberForm.fields.passwordConfirmation.$errors.map(Boolean),
    });

    Button({
      text: "Save",
      event: memberForm.submit,
      size: "base",
      prevent: true,
      variant: "default",
    });

    ErrorHint($editError, $editError.map(Boolean));
  });
};
