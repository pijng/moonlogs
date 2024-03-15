import { Button, ErrorHint, Input, Multiselect, Select } from "@/shared/ui";
import { h, spec } from "forest";
import { $creationError, events, memberForm } from "./model";
import { createStore } from "effector";
import { UserRole } from "@/shared/api/users";
import { tagModel } from "@/entities/tag";
import { i18n } from "@/shared/lib/i18n";

export const NewMemberForm = () => {
  h("form", () => {
    Input({
      type: "text",
      label: i18n("members.form.name"),
      value: memberForm.fields.name.$value,
      inputChanged: memberForm.fields.name.changed,
      errorText: memberForm.fields.name.$errorText,
    });

    Input({
      type: "email",
      label: i18n("members.form.email"),
      value: memberForm.fields.email.$value,
      inputChanged: memberForm.fields.email.changed,
      errorText: memberForm.fields.email.$errorText,
    });

    h("div", () => {
      spec({ classList: ["mb-6"] });

      Select({
        text: i18n("members.form.role"),
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
        text: i18n("members.form.tag.label"),
        hint: i18n("members.form.tag.hint"),
        options: tagModel.$tags.map((tags) => tags.map((t) => ({ name: t.name, id: t.id }))),
        selectedOptions: memberForm.fields.tag_ids.$value,
        optionChecked: events.tagChecked,
        optionUnchecked: events.tagUnchecked,
      });
    });

    Input({
      type: "password",
      label: i18n("members.form.password"),
      required: true,
      value: memberForm.fields.password.$value,
      inputChanged: memberForm.fields.password.changed,
      errorText: memberForm.fields.password.$errorText,
    });

    Input({
      type: "password",
      label: i18n("members.form.confirm_password"),
      required: true,
      value: memberForm.fields.passwordConfirmation.$value,
      inputChanged: memberForm.fields.passwordConfirmation.changed,
      errorText: memberForm.fields.passwordConfirmation.$errorText,
    });

    Button({
      text: i18n("buttons.create"),
      event: memberForm.submit,
      size: "base",
      prevent: true,
      variant: "default",
    });

    ErrorHint($creationError, $creationError.map(Boolean));
  });
};
