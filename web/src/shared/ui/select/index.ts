import { Event, Store, combine, createEffect, createEvent, sample } from "effector";
import { DOMElement, h, list, node, spec } from "forest";

export const Select = ({
  value,
  text,
  options,
  optionSelected,
}: {
  value: Store<any>;
  text: string;
  options: Store<any[]>;
  optionSelected: Event<any>;
}) => {
  const $selected = combine([options, value], ([options, value]) => {
    return options.find((o) => o === value || o.id === value) ?? null;
  });

  h("div", () => {
    h("label", {
      classList: ["block", "mb-2", "text-sm", "font-medium", "text-gray-900", "dark:text-white"],
      text: text,
    });

    h("select", () => {
      const localOptionSelected = createEvent<any>();
      sample({
        source: localOptionSelected,
        fn: (option) => option.target.value,
        target: optionSelected,
      });

      spec({
        classList: [
          "bg-gray-50",
          "border",
          "border-gray-300",
          "text-gray-900",
          "text-sm",
          "rounded-lg",
          "focus:ring-blue-500",
          "focus:border-blue-500",
          "block",
          "w-full",
          "p-2.5",
          "dark:bg-gray-700",
          "dark:border-gray-600",
          "dark:placeholder-gray-400",
          "dark:text-white",
          "dark:focus:ring-blue-500",
          "dark:focus:border-blue-500",
        ],
        handler: { change: localOptionSelected },
      });

      h("option", {
        attr: { value: "", selected: true },
        text: "",
      });

      list(options, ({ store: option }) => {
        const localSelected = combine([$selected, option], ([selected, option]) => {
          return selected === option || selected === option.id;
        });

        h("option", {
          attr: {
            value: option.map((o) => o.name || o),
            selected: localSelected,
          },
          text: option.map((o) => o.title || o.name || o),
        });
      });

      // I hate myself
      node((ref: DOMElement) => {
        const selectElement = ref as HTMLSelectElement;
        const clearSelectionFx = createEffect(() => {
          selectElement.selectedIndex = -1;
        });

        sample({
          source: value,
          filter: (value) => !Boolean(value),
          target: clearSelectionFx,
        });
      });
    });
  });
};
