import { Rule } from "effector-forms";

export const rules = {
  required: (): Rule<any> => ({
    name: "required",
    validator: (value) => ({
      isValid: Boolean(value),
      errorText: "Required field",
    }),
  }),
  email: (): Rule<string> => ({
    name: "email",
    validator: (value) => ({
      isValid: /\S+@\S+\.\S+/.test(value),
      errorText: "Invalid email",
    }),
  }),
  password: (): Rule<string> => ({
    name: "password-equal",
    validator: (confirm, { password }) => ({
      isValid: password === confirm,
      errorText: "Passwords must match",
    }),
  }),
};
