import { Event, Store, combine, createEvent, createStore, sample } from "effector";
import { DOMElement, h, list, node, remap, spec } from "forest";
import { condition } from "patronum";

interface SelectItem {
  id: any;
  name: string;
}

export const Multiselect = ({
  text,
  options,
  selectedOptions,
  hint,
  optionChecked,
  optionUnchecked,
}: {
  text: Store<string> | string;
  options: Store<SelectItem[]>;
  selectedOptions?: Store<number[]>;
  hint: Store<string>;
  optionChecked: Event<any>;
  optionUnchecked: Event<any>;
}) => {
  const dropdownTriggered = createEvent<MouseEvent>();
  const $localVisible = createStore(false);
  const outsideClicked = createEvent<{ node: DOMElement; event: any }>();

  sample({
    source: $localVisible,
    clock: dropdownTriggered,
    fn: (v) => !v,
    target: $localVisible,
  });

  sample({
    source: $localVisible,
    clock: outsideClicked,
    filter: (visible, { node, event }) => !node.contains(event.target) && visible,
    fn: () => false,
    target: $localVisible,
  });

  const localSelectedOptions = selectedOptions || createStore<number[]>([]);

  h("div", () => {
    spec({
      classList: ["relative", "mb-6"],
    });

    h("label", {
      classList: ["block", "mb-2", "text-sm", "font-medium", "text-gray-900", "dark:text-white"],
      text: text,
    });

    h("input", () => {
      const selectedText = combine([options, localSelectedOptions, hint], ([opts, selectedOpts, hint]) => {
        if (!selectedOpts || selectedOpts.length === 0) return hint;

        return opts
          .filter((o) => selectedOpts?.includes(o.id))
          .map((o) => o.name)
          .join(", ");
      });

      spec({
        classList: [
          "cursor-default",
          "bg-gray-50",
          "border",
          "border-gray-300",
          "text-gray-900",
          "text-sm",
          "rounded-lg",
          "focus:outline",
          "outline-1",
          "focus:ring-blue-500",
          "focus:border-blue-500",
          "block",
          "w-full",
          "p-2.5",
          "dark:bg-squid-ink",
          "dark:placeholder-gray-400",
          "dark:text-white",
          "dark:border-slate-gray",
          "dark:focus:ring-blue-500",
          "dark:focus:border-blue-500",
          "cursor-pointer",
        ],
        attr: {
          readonly: true,
          type: "select",
          value: selectedText,
        },
        handler: { on: { click: dropdownTriggered } },
      });
    });

    h("div", () => {
      spec({
        visible: $localVisible,
        classList: [
          "max-h-56",
          "overflow-auto",
          "left-0",
          "absolute",
          "left-0",
          "top-18",
          "w-fit",
          "border",
          "border-gray-200",
          "dark:bg-raisin-black",
          "dark:border-shadow-gray",
          "z-50",
          "bg-white",
          "divide-y",
          "divide-gray-100",
          "rounded-lg",
          "shadow",
          "dark:bg-squid-ink",
          "dark:divide-gray-600",
          "shadow-2xl",
        ],
      });

      h("ul", () => {
        spec({
          classList: ["flex", "flex-col", "flex-wrap", "text-gray-700", "dark:text-gray-200"],
        });

        list(options, ({ store: option }) => {
          h("li", () => {
            spec({
              classList: [
                "hover:bg-gray-100",
                "dark:hover:bg-slate-gray",
                "flex",
                "border-gray-300",
                "dark:border-shadow-gray",
                "items-center",
                "px-5",
                "flex-auto",
                "shrink-0",
              ],
            });

            const localOptionSelected = createEvent<any>();

            const localChecked = createEvent();
            sample({
              source: option,
              clock: localChecked,
              fn: (option) => option.id,
              target: optionChecked,
            });

            const localUnchecked = createEvent();
            sample({
              source: option,
              clock: localUnchecked,
              fn: (option) => option.id,
              target: optionUnchecked,
            });

            condition({
              source: localOptionSelected,
              if: (event) => event.target.checked,
              then: localChecked,
              else: localUnchecked,
            });

            h("input", () => {
              const selected = combine(option, localSelectedOptions, (opt, selectedOpts) => {
                if (!selectedOpts) return false;
                return selectedOpts.includes(opt.id);
              });

              spec({
                classList: [
                  "bg-gray-50",
                  "border-gray-300",
                  "text-gray-900",
                  "rounded-lg",
                  "focus:ring-blue-500",
                  "focus:border-blue-500",
                  "block",
                  "w-4",
                  "h-4",
                  "p-2.5",
                  "dark:bg-squid-ink",
                  "dark:border-slate-gray",
                  "dark:placeholder-gray-400",
                  "dark:text-white",
                  "dark:focus:ring-blue-500",
                  "dark:focus:border-blue-500",
                  "cursor-pointer",
                ],
                attr: { type: "checkbox", id: remap(option, "id"), checked: selected },
                handler: { on: { change: localOptionSelected } },
              });
            });

            h("label", {
              classList: [
                "block",
                "ml-3",
                "w-full",
                "text-sm",
                "font-medium",
                "py-2.5",
                "text-gray-900",
                "dark:text-white",
                "cursor-pointer",
              ],
              text: remap(option, "name"),
              attr: { for: remap(option, "id") },
            });
          });
        });
      });
    });

    node((node) => {
      document.addEventListener("click", (event) => {
        outsideClicked({ node, event });
      });
    });
  });
};
