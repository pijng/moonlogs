import { Event, createEvent, sample } from "effector";
import { h, spec } from "forest";

const forwardIcon = () => {
  h("svg", () => {
    spec({
      classList: ["w-5", "h-5"],
      attr: { "aria-hidden": true, xmlns: "http://www.w3.org/2000/svg", fill: "none", viewBox: "0 0 14 10" },
    });
    h("path", {
      attr: {
        stroke: "currentColor",
        "stroke-linecap": "round",
        "stroke-linejoin": "round",
        "stroke-width": "2",
        d: "M1 5h12m0 0L9 1m4 4L9 9",
      },
    });
  });
};

const backIcon = () => {
  h("svg", () => {
    spec({
      classList: ["w-5", "h-5"],
      attr: { "aria-hidden": true, xmlns: "http://www.w3.org/2000/svg", fill: "none", viewBox: "0 0 14 10" },
    });
    h("path", {
      attr: {
        stroke: "currentColor",
        "stroke-linecap": "round",
        "stroke-linejoin": "round",
        "stroke-width": "2",
        d: "M13 5H1m0 0 4 4M1 5l4-4",
      },
    });
  });
};

export const NavigationButton = (variant: "back" | "forward", event: Event<void>) => {
  const buttonClicked = createEvent<MouseEvent>();

  sample({
    clock: buttonClicked,
    target: event,
  });

  h("button", () => {
    spec({
      handler: {
        click: buttonClicked,
      },
      classList: [
        "text-white",
        "bg-blue-700",
        "hover:bg-blue-800",
        "focus:ring-4",
        "focus:outline-none",
        "focus:ring-blue-300",
        "font-medium",
        "rounded-lg",
        "text-sm",
        "p-2.5",
        "text-center",
        "inline-flex",
        "items-center",
        "mr-2",
        "dark:bg-blue-600",
        "dark:hover:bg-blue-700",
        "dark:focus:ring-blue-800",
      ],
      attr: { type: "button" },
    });

    variant === "back" ? backIcon() : forwardIcon();
  });
};
