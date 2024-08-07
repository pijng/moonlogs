import { i18n } from "@/shared/lib";
import { Button, ButtonVariant, ClockIcon, DownIcon, Input } from "@/shared/ui";
import { Event, Store, createEvent, createStore, sample } from "effector";
import { DOMElement, h, node, spec } from "forest";

export const FilterDate = ({
  applied,
  fromTime,
  fromTimeChanged,
  toTime,
  toTimeChanged,
}: {
  applied: Store<boolean>;
  fromTime: Store<any>;
  fromTimeChanged: Event<any>;
  toTime: Store<any>;
  toTimeChanged: Event<any>;
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

  h("div", () => {
    spec({ classList: ["relative"] });

    Button({
      text: i18n("log_groups.filters.time.label"),
      variant: applied.map<ButtonVariant>((state) => (state ? "default" : "alternative")),
      size: "small",
      event: dropdownTriggered,
      preIcon: ClockIcon,
      postIcon: DownIcon,
    });

    h("div", () => {
      spec({
        visible: $localVisible,
      });

      h("div", () => {
        spec({
          classList: [
            "w-fit",
            "top-0",
            "left-0",
            "w-5/6",
            "absolute",
            "left-0",
            "top-9",
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
            classList: ["flex", "flex-col", "flex-wrap", "py-2", "text-sm", "text-gray-700", "dark:text-gray-200"],
          });

          h("li", () => {
            spec({
              classList: ["block", "px-4", "py-2", "flex-auto", "shrink-0"],
            });

            Input({
              type: "datetime-local",
              value: fromTime,
              label: i18n("log_groups.filters.time.from"),
              inputChanged: fromTimeChanged,
            });
          });

          h("li", () => {
            spec({
              classList: ["block", "px-4", "py-2", "flex-auto", "shrink-0"],
            });

            Input({
              type: "datetime-local",
              value: toTime,
              label: i18n("log_groups.filters.time.to"),
              inputChanged: toTimeChanged,
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
