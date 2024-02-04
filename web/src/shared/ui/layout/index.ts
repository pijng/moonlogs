import { h, spec } from "forest";
import { Sidebar } from "@/shared/ui";
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

      Sidebar();
    });

    h("div", () => {
      spec({
        classList: {
          "py-2": layoutVisible.map((visible) => visible),
          "px-4": layoutVisible.map((visible) => visible),
          "sm:p-7": layoutVisible.map((visible) => visible),
          "sm:ml-64": layoutVisible,
        },
      });

      content();
    });
  });
};
