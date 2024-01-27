import { Button, ErrorHint, Input, Multiselect, Select } from "@/shared/ui";
import { h, spec } from "forest";
import { $editError, $tagsDropwdownIsOpened, deleteUserClicked, deleteUserFx, editUserFx, events, memberForm } from "./model";
import { createStore } from "effector";
import { UserRole } from "@/shared/api/users";
import { tagModel } from "@/entities/tag";

export const EditMemberForm = () => {
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
        value: memberForm.fields.role.$value.map(String),
        options: createStore<UserRole[]>(["Member", "Admin"]),
        optionSelected: memberForm.fields.role.changed,
        clear: [editUserFx.done, deleteUserFx.done],
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
        event: events.tagSelectionClicked,
        visible: $tagsDropwdownIsOpened,
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
      value: memberForm.fields.passwordConfirmation.$value,
      type: "password",
      label: "Confirm password",
      required: true,
      inputChanged: memberForm.fields.passwordConfirmation.changed,
      errorText: memberForm.fields.passwordConfirmation.$errorText,
    });

    h("div", () => {
      spec({ classList: ["flex", "justify-start", "space-x-2"] });

      Button({
        text: "Save",
        event: memberForm.submit,
        size: "base",
        prevent: true,
        variant: "default",
      });

      Button({
        text: "Delete",
        event: deleteUserClicked,
        size: "base",
        prevent: true,
        variant: "delete",
      });
    });

    ErrorHint($editError, $editError.map(Boolean));
  });
};
