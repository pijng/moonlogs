import { h, spec } from "forest";
import { Sidebar, SidebarButton } from "@/shared/ui";
import { createEvent } from "effector";

export const layoutClicked = createEvent<MouseEvent>();

export const Layout = (content: () => void) => {
  h("div", () => {
    spec({
      classList: ["w-full", "min-h-screen", "bg-white", "text-slate-700", "dark:bg-slate-900", "dark:text-slate-300"],
      handler: { on: { click: layoutClicked } },
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
