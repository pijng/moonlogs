import { Event, Store, createEvent, createStore, sample } from "effector";
import { h, list, spec } from "forest";

export const Select = ({
  id,
  value,
  text,
  options,
  optionSelected,
}: {
  id: string;
  value?: Store<string>;
  text: string;
  options: Store<any[]>;
  optionSelected: Event<any>;
}) => {
  h("div", () => {
    spec({
      classList: ["mb-6"],
    });

    h("label", {
      classList: ["block", "mb-2", "text-sm", "font-medium", "text-gray-900", "dark:text-white"],
      attr: { for: id.toLowerCase() },
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
        attr: { id: id, value: value ?? createStore("") },
        handler: { change: localOptionSelected },
      });

      list(options, ({ store: option }) => {
        h("option", {
          attr: { value: option },
          text: option,
        });
      });
    });
  });
};
