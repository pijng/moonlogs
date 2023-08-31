import { Event, Store, createEvent, restore, sample } from "effector";
import { ClassListArray, h, node, spec } from "forest";

export type ButtonVariant = "default" | "alternative" | "light";
type Size = "base" | "small" | "extra_small";

const BASE_CLASSES = ["inline-flex", "items-center"];

const VARIANTS: Record<ButtonVariant, string[]> = {
  default: [
    "block",
    "text-white",
    "bg-blue-700",
    "border",
    "dark:border-blue-600",
    "hover:bg-blue-800",
    "font-medium",
    "rounded-lg",
    "dark:bg-blue-600",
    "dark:hover:bg-blue-700",
    "focus:outline-none",
  ],
  alternative: [
    "block",
    "font-medium",
    "text-gray-900",
    "bg-white",
    "rounded-lg",
    "border",
    "border-gray-200",
    "hover:bg-gray-100",
    "hover:text-blue-700",
    "dark:bg-gray-800",
    "dark:text-gray-400",
    "dark:border-gray-600",
    "dark:hover:text-white",
    "dark:hover:bg-gray-700",
  ],
  light: ["block", "font-medium", "text-gray-900", "hover:text-blue-700", "dark:text-gray-400", "dark:hover:text-white"],
};

const SIZES: Record<Size, string[]> = {
  base: ["px-5", "py-2.5", "text-sm"],
  small: ["px-3", "py-2", "text-sm"],
  extra_small: ["px-3", "py-2", "text-xs"],
};

const buttonClass = (variant: Store<ButtonVariant>, size: Size): Store<string> => {
  const $currentVariant = variant.map((variant) => VARIANTS[variant]);
  const $sumClasses = $currentVariant.map((variant) => BASE_CLASSES.concat(variant).concat(SIZES[size]));

  return $sumClasses.map((classes) => classes.join(" "));
};

export const Button = ({
  text,
  event,
  variant,
  size,
  preIcon,
  postIcon,
}: {
  text: string;
  event?: Event<any>;
  variant: Store<ButtonVariant>;
  size: Size;
  preIcon?: () => void;
  postIcon?: () => void;
}) => {
  const localPreIcon =
    preIcon ||
    function () {
      return;
    };

  const localPostIcon =
    postIcon ||
    function () {
      return;
    };

  h("button", () => {
    localPreIcon();

    const touchClasses = createEvent<string>();
    const $classes = restore(touchClasses, "");

    spec({
      handler: { click: event ?? createEvent() },
      classList: [$classes] as ClassListArray,
      text: text,
    });

    localPostIcon();

    const touch = createEvent();

    sample({
      clock: [touch, variant],
      source: buttonClass(variant, size),
      target: touchClasses,
    });

    // Hack to apply classes from store
    // https://github.com/effector/effector/issues/964
    node(() => {
      touch();
    });
  });
};
