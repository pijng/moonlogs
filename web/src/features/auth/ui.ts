import { Button, Input } from "@/shared/ui";
import { emailChanged, logInSubmitted, passwordChanged } from "./model";
import { h, spec } from "forest";

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
      variant: "default",
    });
  });
};
