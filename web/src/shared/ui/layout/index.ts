import { h, spec } from "forest";

export const Layout = (content: () => void) => {
  h("div", () => {
    spec({
      // outer most
      classList: ["w-full", "bg-white", "text-slate-700", "dark:bg-slate-900", "dark:text-slate-300"],
    });

    h("div", () => {
      spec({
        // main content
        classList: ["p-5"],
      });

      content();
    });
  });
};
