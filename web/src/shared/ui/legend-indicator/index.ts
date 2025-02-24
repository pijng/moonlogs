import { createEvent, createStore, is, restore, sample, Store } from "effector";
import { ClassListArray, h, node, spec } from "forest";

const baseClasses = ["flex", "w-2.5", "h-2.5", "rounded-full", "me-1.5", "shrink-0"];

const indicatorClass = (colors: Store<string> | string): Store<string> => {
  const $rawClrs = is.unit(colors) ? colors : createStore(colors);
  const $bgColors = $rawClrs.map((clr) => `bg-${clr}`);

  return $bgColors.map((colors) => baseClasses.concat(colors).join(" "));
};

export const LegendIndicator = ({ text, color }: { text: string | Store<string>; color: Store<string> | string }) => {
  h("span", () => {
    spec({
      classList: [
        "bg-gray-100",
        "dark:bg-slate-gray",
        "dark:text-gray-200",
        "py-1.5",
        "px-2.5",
        "flex",
        "rounded-lg",
        "items-center",
        "text-sm",
        "font-medium",
        "text-gray-900",
        "dark:text-white",
        "me3",
      ],
    });

    const touchClasses = createEvent<string>();
    const $classes = restore(touchClasses, "");

    h("span", {
      classList: [$classes] as ClassListArray,
    });

    const touch = createEvent();
    sample({
      clock: touch,
      source: indicatorClass(color),
      target: touchClasses,
    });

    spec({ text });

    node(() => {
      touch();
    });
  });
};
