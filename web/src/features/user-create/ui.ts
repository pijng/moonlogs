import { Button, ErrorHint, Input, Multiselect, Select } from "@/shared/ui";
import { h, spec } from "forest";
import { $creationError, events, memberForm } from "./model";
import { createStore } from "effector";
import { UserRole } from "@/shared/api/users";
import { tagModel } from "@/entities/tag";

export const NewMemberForm = () => {
  h("form", () => {
    Input({
      type: "text",
      label: "Name",
      value: memberForm.fields.name.$value,
      inputChanged: memberForm.fields.name.changed,
      errorText: memberForm.fields.name.$errorText,
    });

    Input({
      type: "email",
      label: "Email",
      value: memberForm.fields.email.$value,
      inputChanged: memberForm.fields.email.changed,
      errorText: memberForm.fields.email.$errorText,
    });

    h("div", () => {
      spec({ classList: ["mb-6"] });

      Select({
        text: "Select a role",
        value: memberForm.fields.role.$value,
        options: createStore<UserRole[]>(["Member", "Admin"]),
        optionSelected: memberForm.fields.role.changed,
        withBlank: createStore(false),
      });
    });

    const $tagSelectionVisible = memberForm.fields.role.$value.map((role) => role !== "Admin");

    h("div", () => {
      spec({
        visible: $tagSelectionVisible,
      });

      Multiselect({
        text: "Select tags",
        options: tagModel.$tags,
        selectedOptions: memberForm.fields.tag_ids.$value,
        optionChecked: events.tagChecked,
        optionUnchecked: events.tagUnchecked,
      });
    });

    Input({
      type: "password",
      label: "Password",
      required: true,
      value: memberForm.fields.password.$value,
      inputChanged: memberForm.fields.password.changed,
      errorText: memberForm.fields.password.$errorText,
    });

    Input({
      type: "password",
      label: "Confirm password",
      required: true,
      value: memberForm.fields.passwordConfirmation.$value,
      inputChanged: memberForm.fields.passwordConfirmation.changed,
      errorText: memberForm.fields.passwordConfirmation.$errorText,
    });

    Button({
      text: "Create",
      event: memberForm.submit,
      size: "base",
      prevent: true,
      variant: "default",
    });

    ErrorHint($creationError, $creationError.map(Boolean));
  });
};
