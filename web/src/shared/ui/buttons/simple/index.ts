import { Event, createEvent } from "effector";
import { h, spec } from "forest";

type Variant = "default" | "alternative";
type Size = "base" | "small" | "extra_small";

const BASE_CLASSES = ["inline-flex", "items-center"];

const VARIANTS: Record<Variant, string[]> = {
  default: [
    "block",
    "text-white",
    "bg-blue-700",
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
};

const SIZES: Record<Size, string[]> = {
  base: ["px-5", "py-2.5", "text-sm"],
  small: ["px-3", "py-2", "text-sm"],
  extra_small: ["px-3", "py-2", "text-xs"],
};

const buttonClass = (variant: Variant, size: Size): string[] => {
  return BASE_CLASSES.concat(VARIANTS[variant].concat(SIZES[size]));
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
  variant: Variant;
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

    spec({
      handler: { click: event ?? createEvent() },
      classList: buttonClass(variant, size),
      text: text,
    });

    localPostIcon();
  });
};
