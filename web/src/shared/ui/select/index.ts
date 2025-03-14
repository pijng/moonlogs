import { hashCode } from "@/shared/lib";
import { Event, Store, combine, createEvent, createStore, sample } from "effector";
import { DOMElement, h, list, node, spec } from "forest";
import { i18n } from "@/shared/lib/i18n";
import { Label, DownIcon, Search } from "@/shared/ui";

export const Select = ({
  value,
  text,
  hint,
  options,
  optionSelected,
  withBlank,
}: {
  value: Store<any>;
  text?: Store<string> | string;
  hint?: Store<string> | string;
  options: Store<any[]>;
  optionSelected: Event<any>;
  withBlank?: Store<boolean>;
}) => {
  const $searchQuery = createStore("");
  const queryChanged = createEvent<string>();
  $searchQuery.on(queryChanged, (_, query) => query);

  const $filteredOptions = combine(options, $searchQuery, (opts, query) => {
    const lowerQuery = query.toLocaleLowerCase();
    return opts.filter((opt) => {
      const optTitle = !!opt.title ? opt.title.toString() : "";
      const lowerOptTitle = optTitle.toLocaleLowerCase();
      const optName = !!opt.name ? opt.name.toString() : "";
      const lowerOptName = optName.toLocaleLowerCase();
      const optId = !!opt.id ? opt.id.toString() : "";
      const lowerOptId = optId.toLocaleLowerCase();

      return (
        (typeof opt === "string" && opt.toLocaleLowerCase().includes(lowerQuery)) ||
        lowerOptTitle.includes(lowerQuery) ||
        lowerOptName.includes(lowerQuery) ||
        lowerOptId.toString().includes(lowerQuery)
      );
    });
  });

  const $optionsHash = options.map((o) => hashCode(o.join("")));

  const dropdownTriggered = createEvent<MouseEvent>();
  const $withSearch = options.map((opts) => opts.length > 8);
  const $localVisible = createStore(false);
  const outsideClicked = createEvent<{ node: DOMElement; event: any }>();
  sample({
    source: $localVisible,
    clock: dropdownTriggered,
    fn: (v) => !v,
    target: $localVisible,
  });

  sample({
    clock: optionSelected,
    fn: () => false,
    target: $localVisible,
  });

  sample({
    source: { visible: $localVisible, hash: $optionsHash },
    clock: outsideClicked,
    filter: ({ visible, hash }, { node, event }) => {
      return !node.contains(event.target) && hash !== event.target.id && visible;
    },
    fn: () => false,
    target: $localVisible,
  });

  const $selected = combine([options, value], ([options, value]) => {
    return options.find((o) => o === value || o.id === value || o.name === value) || "";
  });

  h("div", () => {
    spec({
      classList: ["relative"],
    });

    h("div", () => {
      spec({ classList: ["mb-2"] });

      Label({ text: text ?? createStore(""), hint: hint });
    });

    // h("label", {
    //   classList: ["block", "mb-2", "text-sm", "font-medium", "text-gray-900", "dark:text-white"],
    //   text: text ?? createStore(""),
    // });

    h("div", () => {
      spec({ classList: ["relative"] });

      h("input", () => {
        spec({
          classList: [
            "cursor-default",
            "bg-gray-50",
            "focus:outline",
            "outline-1",
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
            "dark:bg-squid-ink",
            "dark:border-slate-gray",
            "dark:placeholder-gray-400",
            "dark:text-white",
            "dark:focus:ring-blue-500",
            "dark:focus:border-blue-500",
            "cursor-pointer",
          ],
          attr: {
            readonly: true,
            type: "select",
            value: $selected.map((s) => s?.title || s?.name || s),
            id: $optionsHash,
          },
          handler: { on: { click: dropdownTriggered } },
        });
      });

      h("div", () => {
        spec({ classList: ["absolute", "top-4", "right-2"] });

        DownIcon();
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

      h("div", () => {
        spec({
          visible: $withSearch,
          classList: ["bg-white", "dark:bg-squid-ink", "top-0", "sticky", "px-3"],
        });

        Search(queryChanged, $searchQuery);
      });

      h("ul", () => {
        spec({
          classList: ["flex", "flex-col", "flex-wrap", "text-gray-700", "dark:text-gray-200"],
        });

        h("li", {
          classList: [
            "hover:bg-gray-100",
            "dark:hover:bg-gray-600",
            "flex",
            "border-gray-300",
            "dark:border-shadow-gray",
            "items-center",
            "px-6",
            "py-2.5",
            "flex-auto",
            "w-full",
            "text-sm",
            "font-medium",
            "text-gray-900",
            "dark:text-white",
            "shrink-0",
            "block",
            "cursor-pointer",
          ],
          text: i18n("miscellaneous.blank_option"),
          visible: withBlank,
          handler: { on: { click: optionSelected.prepend(() => null) } },
        });

        list($filteredOptions, ({ store: option }) => {
          h("li", () => {
            const localOptionSelected = createEvent<any>();
            sample({
              clock: localOptionSelected,
              fn: () => false,
              target: $localVisible,
            });
            sample({
              source: option,
              clock: localOptionSelected,
              fn: (option) => option.name || option,
              target: optionSelected,
            });

            spec({
              classList: [
                "hover:bg-gray-100",
                "dark:hover:bg-gray-600",
                "flex",
                "border-gray-300",
                "dark:border-shadow-gray",
                "items-center",
                "px-6",
                "py-2.5",
                "flex-auto",
                "w-full",
                "text-sm",
                "font-medium",
                "text-gray-900",
                "dark:text-white",
                "shrink-0",
                "block",
                "cursor-pointer",
              ],
              text: option.map((o) => o.title || o.name || o),
              handler: { on: { click: localOptionSelected } },
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
