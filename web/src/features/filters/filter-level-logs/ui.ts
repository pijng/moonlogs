import { Level } from "@/shared/api";
import { i18n } from "@/shared/lib";
import { Button, ButtonVariant, DownIcon, FireIcon } from "@/shared/ui";
import { Event, Store, combine, createEvent, createStore, sample } from "effector";
import { DOMElement, h, list, node, spec } from "forest";

export const FilterLevel = ({
  applied,
  levelChanged,
  level,
  resetLevelFilter,
}: {
  applied: Store<boolean>;
  levelChanged: Event<Level>;
  level: Store<string>;
  resetLevelFilter: Event<any>;
}) => {
  const dropdownTriggered = createEvent<MouseEvent>();
  const $localVisible = createStore(false);
  const outsideClicked = createEvent<{ node: DOMElement; event: any }>();
  const blankOptionClicked = createEvent<MouseEvent>();

  sample({
    clock: blankOptionClicked,
    target: [dropdownTriggered, resetLevelFilter],
  });

  sample({
    source: $localVisible,
    clock: [dropdownTriggered, levelChanged],
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

  const $levels = createStore<Level[]>(["Info", "Error", "Warn", "Debug", "Trace", "Fatal"]);

  h("div", () => {
    spec({ classList: ["relative"] });

    Button({
      text: combine([i18n("log_groups.filters.level.label"), level], ([defaultLabel, level]) => level || defaultLabel),
      variant: applied.map<ButtonVariant>((state) => (state ? "default" : "alternative")),
      size: "small",
      event: dropdownTriggered,
      fullWidth: true,
      preIcon: FireIcon,
      postIcon: DownIcon,
    });

    h("div", () => {
      spec({
        visible: $localVisible,
      });

      h("div", () => {
        spec({
          classList: [
            "max-h-56",
            "overflow-auto",
            "dark:scrollbar",
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
            "dark:divide-gray-600",
            "shadow-2xl",
          ],
        });

        h("ul", () => {
          spec({
            classList: ["flex", "flex-col", "flex-wrap", "text-gray-700", "dark:text-gray-200"],
          });

          h("li", {
            classList: [
              "hover:bg-gray-100",
              "dark:hover:bg-slate-gray",
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
            handler: { on: { click: blankOptionClicked } },
          });

          list($levels, ({ store: level }) => {
            h("li", () => {
              const localLevelSelected = createEvent<any>();
              sample({
                source: level,
                clock: localLevelSelected,
                fn: (level) => level,
                target: levelChanged,
              });

              spec({
                classList: [
                  "hover:bg-gray-100",
                  "dark:hover:bg-slate-gray",
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
                text: level.map((l) => l),
                handler: { on: { click: localLevelSelected } },
              });
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
