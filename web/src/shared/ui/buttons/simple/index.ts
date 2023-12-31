import { Event, Store, createEvent, createStore, is, restore, sample } from "effector";
import { ClassListArray, h, node, spec } from "forest";

export type ButtonVariant = "default" | "alternative" | "light";
type Size = "base" | "small" | "extra_small";
type Style = "default" | "round";

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
    "dark:bg-blue-600",
    "dark:hover:bg-blue-700",
    "focus:outline-none",
  ],
  alternative: [
    "block",
    "font-medium",
    "text-gray-900",
    "bg-white",
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

const SIZES: Record<Size, Record<Style, string[]>> = {
  base: {
    default: ["px-5", "py-2.5", "text-sm", "rounded-lg"],
    round: ["p-2.5", "text-sm"],
  },
  small: {
    default: ["px-3", "py-2", "text-sm", "rounded-lg"],
    round: ["p-2", "text-sm"],
  },
  extra_small: {
    default: ["px-3", "py-2", "text-xs", "rounded-lg"],
    round: ["p-1.5", "text-xs"],
  },
};

const STYLES: Record<Style, string[]> = {
  round: ["rounded-full"],
  default: [""],
};

const buttonClass = (variant: Store<ButtonVariant>, size: Size, style: Style): Store<string> => {
  const $currentVariant = variant.map((variant) => VARIANTS[variant]);
  const $sumClasses = $currentVariant.map((variant) =>
    STYLES[style].concat(BASE_CLASSES).concat(variant).concat(SIZES[size][style]),
  );

  return $sumClasses.map((classes) => classes.join(" "));
};

export const Button = ({
  text,
  event,
  variant,
  size,
  style,
  visible,
  prevent,
  preIcon,
  postIcon,
}: {
  text: string;
  event?: Event<any>;
  variant: ButtonVariant | Store<ButtonVariant>;
  style?: Style;
  size: Size;
  visible?: Store<boolean>;
  prevent?: boolean;
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
      handler: { on: { click: event ?? createEvent() }, config: { prevent: prevent } },
      classList: [$classes] as ClassListArray,
      text: text,
      visible: visible,
    });

    localPostIcon();

    const touch = createEvent();
    const $localVariant = is.unit(variant) ? variant : createStore(variant);
    const localStyle = style ?? "default";

    sample({
      clock: [touch, $localVariant],
      source: buttonClass($localVariant, size, localStyle),
      target: touchClasses,
    });

    // Hack to apply classes from store
    // https://github.com/effector/effector/issues/964
    node(() => {
      touch();
    });
  });
};
