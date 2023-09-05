import { Button, ButtonVariant, Input } from "@/shared/ui";
import { emailChanged, logInSubmitted, passwordChanged } from "./model";
import { h, spec } from "forest";
import { createStore } from "effector";

export const Auth = () => {
  h("div", () => {
    spec({
      classList: ["mt-10", "max-w-sm", "w-full"],
    });

    Input({ type: "email", label: "Email", inputChanged: emailChanged });
    Input({ type: "password", label: "Password", inputChanged: passwordChanged });

    Button({
      text: "Log in",
      event: logInSubmitted,
      size: "base",
      variant: createStore<ButtonVariant>("default"),
    });
  });
};
