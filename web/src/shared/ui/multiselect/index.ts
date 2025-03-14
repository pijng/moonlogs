import { Event, Store, combine, createEvent, createStore, sample } from "effector";
import { DOMElement, h, list, node, remap, spec } from "forest";
import { condition } from "patronum";
import { ErrorHint } from "../error-hint";
import { Search } from "../search";

interface SelectItem<T extends string | number> {
  id: T;
  name: string;
}

export const Multiselect = <T extends string | number>({
  text,
  options,
  selectedOptions,
  hint,
  optionChecked,
  optionUnchecked,
  errorText,
}: {
  text: Store<string> | string;
  options: Store<SelectItem<T>[]>;
  selectedOptions?: Store<T[]>;
  hint: Store<string>;
  optionChecked: Event<any>;
  optionUnchecked: Event<any>;
  errorText?: Store<string>;
}) => {
  const $searchQuery = createStore("");
  const queryChanged = createEvent<string>();
  $searchQuery.on(queryChanged, (_, query) => query);

  const $filteredOptions = combine(options, $searchQuery, (opts, query) => {
    const lowerQuery = query.toLocaleLowerCase();
    return opts.filter((opt) => {
      const lowerName = opt.name.toLocaleLowerCase();
      const lowerId = opt.id.toString().toLocaleLowerCase();

      return lowerName.includes(lowerQuery) || lowerId.includes(lowerQuery);
    });
  });

  const dropdownTriggered = createEvent<MouseEvent>();
  const $localVisible = createStore(false);
  const $withSearch = options.map((opts) => opts.length > 8);
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

  const localSelectedOptions = selectedOptions || createStore<T[]>([]);

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
          "max-h-96",
          "overflow-auto",
          "dark:scrollbar",
          "left-0",
          "absolute",
          "left-0",
          "top-18",
          "w-full",
          "border",
          "border-gray-200",
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

      h("div", () => {
        spec({ visible: $withSearch, classList: ["bg-inherit", "top-0", "sticky", "px-3"] });

        Search(queryChanged, $searchQuery);
      });

      h("ul", () => {
        spec({
          classList: ["flex", "flex-col", "flex-wrap", "text-gray-700", "dark:text-gray-200"],
        });

        list($filteredOptions, ({ store: option }) => {
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

    ErrorHint(errorText, errorText?.map(Boolean));

    node((node) => {
      document.addEventListener("click", (event) => {
        outsideClicked({ node, event });
      });
    });
  });
};
