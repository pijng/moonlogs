import { Event, Store, createEvent, sample } from "effector";
import { h, spec } from "forest";
import { Button } from "@/shared/ui";

export const Search = (inputChanged: Event<string>, searchQuery: Store<string>) => {
  const searchCleared = createEvent();
  sample({
    clock: searchCleared,
    fn: () => "",
    target: inputChanged,
  });

  h("div", () => {
    spec({
      classList: ["bg-white", "dark:bg-gray-900", "max-w-xl", "py-3"],
    });

    h("label", {
      attr: { for: "table-search" },
      classList: ["sr-only"],
      text: "Search",
    });

    h("div", () => {
      spec({
        classList: ["relative", "flex"],
      });

      h("div", () => {
        spec({
          classList: ["absolute", "inset-y-0", "left-0", "flex", "items-center", "pl-3", "pointer-events-none"],
        });

        h("svg", () => {
          spec({
            classList: ["w-4", "h-4", "text-gray-500", "dark:text-gray-400"],
            attr: { "aria-hidden": true, xmlns: "http://www.w3.org/2000/svg", fill: "none", viewBox: "0 0 20 20" },
          });

          h("path", {
            attr: {
              stroke: "currentColor",
              "stroke-linecap": "round",
              "stroke-linejoin": "round",
              "stroke-width": "2",
              d: "m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z",
            },
          });
        });
      });

      h("input", () => {
        spec({
          attr: {
            type: "text",
            id: "table-search",
            placeholder: "Search",
            value: searchQuery,
          },
          handler: {
            // eslint-disable-next-line @typescript-eslint/ban-ts-comment
            // @ts-ignore
            input: inputChanged.prepend((e) => e.target.value),
          },
          classList: [
            "block",
            "w-full",
            "p-2.5",
            "pl-10",
            "text-sm",
            "text-gray-900",
            "border",
            "border-gray-300",
            "rounded-lg",
            "bg-gray-50",
            "focus:ring-blue-500",
            "focus:border-blue-500",
            "dark:bg-gray-700",
            "dark:border-gray-600",
            "dark:placeholder-gray-400",
            "dark:text-white",
            "dark:focus:ring-blue-500",
            "dark:focus:border-blue-500",
          ],
        });
      });

      Button({
        text: "Clear",
        variant: "light",
        size: "small",
        event: searchCleared,
        visible: searchQuery.map((query) => query.length > 0),
      });
    });
  });
};
