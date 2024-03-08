import { Rule } from "effector-forms";
import { i18n } from "@/shared/lib/i18n";

export const rules = {
  required: (): Rule<any> => ({
    name: "required",
    source: i18n("validations.required"),
    validator: (value, _, errorText) => ({
      isValid: Boolean(value),
      errorText: errorText,
    }),
  }),
  email: (): Rule<string> => ({
    name: "email",
    source: i18n("validations.invalid_email"),
    validator: (value, _, errorText) => ({
      isValid: /\S+@\S+\.\S+/.test(value),
      errorText: errorText,
    }),
  }),
  password: (): Rule<string> => ({
    name: "password-equal",
    source: i18n("validations.passwords_dont_match"),
    validator: (confirm, { password }, errorText) => ({
      isValid: password === confirm,
      errorText: errorText,
    }),
  }),
};
