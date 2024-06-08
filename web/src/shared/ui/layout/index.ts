import { h, spec } from "forest";
import { GeneralTooltip, Sidebar } from "@/shared/ui";
import { Store, createEvent } from "effector";

export const layoutClicked = createEvent<MouseEvent>();

export const Layout = ({ content, layoutVisible }: { content: () => void; layoutVisible: Store<boolean> }) => {
  h("div", () => {
    spec({
      classList: ["w-full", "min-h-screen", "bg-white", "text-slate-700", "dark:bg-eigengrau", "dark:text-slate-300"],
      handler: { on: { click: layoutClicked } },
    });

    GeneralTooltip();

    h("div", () => {
      spec({
        visible: layoutVisible,
      });

      Sidebar();
    });

    h("div", () => {
      spec({
        classList: {
          "py-2": layoutVisible,
          "px-4": layoutVisible,
          "sm:p-7": layoutVisible,
          "sm:ml-64": layoutVisible,
          "min-h-screen": layoutVisible,
          relative: layoutVisible,
        },
      });

      content();
    });
  });
};
