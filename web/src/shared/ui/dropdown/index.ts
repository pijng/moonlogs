import { Event, Store, createEvent, sample } from "effector";
import { h, list, spec } from "forest";
import { FloatingInput, Select } from "@/shared/ui";
import { FilterItem, KindItem } from "@/features";

export const Dropdown = ({
  items,
  itemChanged,
  currentKind,
  kinds,
  kindChanged,
}: {
  items: Store<FilterItem[]>;
  itemChanged: Event<Record<string, any>>;
  currentKind: Store<string>;
  kinds: Store<KindItem[]>;
  kindChanged: Event<string>;
}) => {
  h("div", () => {
    spec({
      classList: [
        "w-max",
        "top-0",
        "left-0",
        "w-5/6",
        "absolute",
        "left-0",
        "top-9",
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
        classList: ["flex", "flex-col", "flex-wrap", "py-2", "text-sm", "text-gray-700", "dark:text-gray-200"],
      });

      h("li", () => {
        spec({
          visible: kinds.map((kinds) => kinds?.length > 0),
          classList: ["block", "px-4", "py-2", "flex-auto", "shrink-0"],
        });

        Select({ value: currentKind, text: "Kind", options: kinds, optionSelected: kindChanged });
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
