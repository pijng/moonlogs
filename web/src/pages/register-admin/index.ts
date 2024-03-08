import { h, spec } from "forest";
import { withRoute } from "atomic-router-forest";

import { registerAdminRoute } from "@/shared/routing";
import { RegisterAdminForm } from "@/widgets";

export const RegisterAdminPage = () => {
  h("div", {
    fn() {
      // This allows to show/hide route if page is matched
      // It is required to call `withRoute` inside `h` call
      withRoute(registerAdminRoute);

      h("div", () => {
        spec({
          classList: ["h-screen", "flex", "flex-col", "items-center", "justify-center"],
        });

        RegisterAdminForm();
      });
    },
  });
};
