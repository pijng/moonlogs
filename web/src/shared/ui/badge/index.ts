import { Level } from "@/shared/api";
import { Store, createEvent, restore, sample } from "effector";
import { ClassListArray, h, node, spec } from "forest";

const BASE_CLASSES = ["text-sm", "font-medium", "mr-2", "px-2.5", "py-0.5", "rounded"];
const LEVEL_CLASSES: Record<string, string[]> = {
  trace: ["bg-green-100", "text-green-800", "dark:bg-green-900", "dark:text-green-300"],
  debug: ["bg-blue-100", "text-blue-800", "dark:bg-blue-900", "dark:text-blue-300"],
  info: ["bg-gray-100", "text-gray-800", "dark:bg-squid-ink", "dark:text-gray-300"],
  warn: ["bg-yellow-100", "text-yellow-800", "dark:bg-yellow-900", "dark:text-yellow-300"],
  error: ["bg-red-100", "text-red-800", "dark:bg-red-900", "dark:text-red-300"],
  fatal: ["bg-pink-100", "text-pink-800", "dark:bg-pink-900", "dark:text-pink-300"],
};

const levelCLasses = (text: Store<Level>) => {
  const classes = text.map((text) => LEVEL_CLASSES[text.toLowerCase()]);
  return classes.map((classes) => BASE_CLASSES.concat(classes).join(" "));
};

export const LevelBadge = (text: Store<Level>) => {
  h("span", () => {
    const touchClasses = createEvent<string>();
    const $classes = restore(touchClasses, "");

    spec({
      classList: [$classes] as ClassListArray,
      text: text,
    });

    const touch = createEvent();

    sample({
      clock: [touch, text],
      source: levelCLasses(text),
      target: touchClasses,
    });

    // Hack to apply classes from store
    // https://github.com/effector/effector/issues/964
    node(() => {
      touch();
    });
  });
};
