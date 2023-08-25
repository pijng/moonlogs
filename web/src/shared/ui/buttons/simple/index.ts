import { Event, createEvent } from "effector";
import { h } from "forest";

type Variant = "default" | "alternative";
type Size = "base" | "small" | "extra_small";

const VARIANTS: Record<Variant, string[]> = {
  default: [
    "m-auto",
    "block",
    "text-white",
    "bg-blue-700",
    "hover:bg-blue-800",
    "focus:ring-4",
    "focus:ring-blue-300",
    "font-medium",
    "rounded-lg",
    "dark:bg-blue-600",
    "dark:hover:bg-blue-700",
    "focus:outline-none",
    "dark:focus:ring-blue-800",
  ],
  alternative: [
    "m-auto",
    "block",
    "font-medium",
    "text-gray-900",
    "focus:outline-none",
    "bg-white",
    "rounded-lg",
    "border",
    "border-gray-200",
    "hover:bg-gray-100",
    "hover:text-blue-700",
    "focus:z-10",
    "focus:ring-4",
    "focus:ring-gray-200",
    "dark:focus:ring-gray-700",
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
  return VARIANTS[variant].concat(SIZES[size]);
};

export const Button = ({
  text,
  event,
  variant,
  size,
}: {
  text: string;
  event?: Event<any>;
  variant: Variant;
  size: Size;
}) => {
  h("button", {
    handler: { click: event ?? createEvent() },
    classList: buttonClass(variant, size),
    text: text,
  });
};
