import { h, spec } from "forest";
import { Sidebar, SidebarButton } from "@/shared/ui";
import { Store, createEvent } from "effector";

export const layoutClicked = createEvent<MouseEvent>();

export const Layout = ({ content, layoutVisible }: { content: () => void; layoutVisible: Store<boolean> }) => {
  h("div", () => {
    spec({
      classList: ["w-full", "min-h-screen", "bg-white", "text-slate-700", "dark:bg-slate-900", "dark:text-slate-300"],
      handler: { on: { click: layoutClicked } },
    });

    h("div", () => {
      spec({
        visible: layoutVisible,
      });

      SidebarButton();
      Sidebar();
    });

    h("div", () => {
      spec({
        classList: {
          "p-7": true,
          "sm:ml-64": layoutVisible,
        },
      });

      content();
    });
  });
};
