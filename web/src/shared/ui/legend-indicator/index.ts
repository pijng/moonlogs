import { createEvent, createStore, sample, Store } from "effector";
import { ClassListArray, h, node, spec } from "forest";

const baseClasses = ["flex", "w-2.5", "h-2.5", "rounded-full", "me-1.5", "shrink-0"];

const indicatorClass = (color: string): string => {
  const bgColor = [`bg-${color}`];

  return baseClasses.concat(bgColor).join(" ");
};

export const LegendIndicator = ({ text, color }: { text: string | Store<string>; color: Store<string> }) => {
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

    const touch = createEvent();
    const $classes = createStore("");

    sample({
      source: color,
      clock: touch,
      fn: (clr) => indicatorClass(clr),
      target: $classes,
    });

    h("span", {
      classList: [$classes] as ClassListArray,
    });

    spec({ text });

    node(() => {
      touch();
    });
  });
};
