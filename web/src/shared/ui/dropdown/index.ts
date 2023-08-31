import { Event, Store, createEvent, sample } from "effector";
import { h, list, spec } from "forest";
import { FloatingInput } from "@/shared/ui";
import { FilterItem } from "@/features";

export const Dropdown = (items: Store<FilterItem[]>, itemChanged: Event<Record<string, any>>) => {
  h("div", () => {
    spec({
      classList: [
        "absolute",
        "border",
        "border-gray-200",
        "dark:bg-gray-800",
        "dark:border-gray-700",
        "z-50",
        "bg-white",
        "divide-y",
        "divide-gray-100",
        "rounded-lg",
        "shadow",
        "dark:bg-gray-700",
        "dark:divide-gray-600",
        "shadow-2xl",
      ],
    });

    h("ul", () => {
      spec({
        classList: ["flex", "flex-wrap", "py-2", "text-sm", "text-gray-700", "dark:text-gray-200"],
      });

      list(items, ({ store: item }) => {
        h("li", () => {
          spec({
            classList: ["block", "px-4", "py-2", "flex-auto", "shrink-0"],
          });

          const inputChanged = createEvent<InputEvent>();

          sample({
            source: item,
            clock: inputChanged,
            fn: (item, event) => {
              // eslint-disable-next-line @typescript-eslint/ban-ts-comment
              // @ts-ignore
              return { [item.key]: event.target.value };
            },
            target: itemChanged,
          });

          FloatingInput(item, inputChanged);
        });
      });
    });
  });
};
