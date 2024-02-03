import { Button, ButtonVariant, ClockIcon, DownIcon, Input } from "@/shared/ui";
import { Event, Store } from "effector";
import { h, spec } from "forest";
import { $filterDateIsOpened, filterDateClicked } from "./model";

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
  Button({
    text: "Time",
    variant: applied.map<ButtonVariant>((state) => (state ? "default" : "alternative")),
    size: "small",
    event: filterDateClicked,
    preIcon: ClockIcon,
    postIcon: DownIcon,
  });

  h("div", () => {
    spec({
      visible: $filterDateIsOpened,
    });

    h("div", () => {
      spec({
        classList: [
          "md:w-64",
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
            classList: ["block", "px-4", "py-2", "flex-auto", "shrink-0"],
          });

          Input({
            type: "datetime-local",
            value: fromTime,
            label: "From",
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
            label: "To",
            inputChanged: toTimeChanged,
          });
        });
      });
    });
  });
};
