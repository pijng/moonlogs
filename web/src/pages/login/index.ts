import { h, spec } from "forest";
import { withRoute } from "atomic-router-forest";

import { loginRoute } from "@/routing/shared";
import { AuthForm } from "@/widgets";

export const LoginPage = () => {
  h("div", {
    fn() {
      // This allows to show/hide route if page is matched
      // It is required to call `withRoute` inside `h` call
      withRoute(loginRoute);

      h("div", () => {
        spec({
          classList: ["h-screen", "flex", "flex-col", "items-center", "justify-center", "px-7"],
        });

        AuthForm();
      });
    },
  });
};
