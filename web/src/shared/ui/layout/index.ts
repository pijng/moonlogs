import { h, spec } from "forest";
import { Sidebar, SidebarButton } from "@/shared/ui";

export const Layout = (content: () => void) => {
  h("div", () => {
    spec({
      // outer most
      classList: ["w-full", "bg-white", "text-slate-700", "dark:bg-slate-900", "dark:text-slate-300"],
    });

    SidebarButton();
    Sidebar();
    h("div", () => {
      spec({
        classList: ["p-4", "sm:ml-64"],
      });

      content();
    });
  });
};
