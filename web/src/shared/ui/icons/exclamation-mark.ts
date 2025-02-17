import { h, spec } from "forest";

export const ExclamationMarkIcon = () => {
  h("svg", () => {
    spec({
      classList: ["w-6", "h-6"],
      attr: {
        "aria-hidden": "true",
        xmlhs: "http://www.w3.org/2000/svg",
        width: 24,
        height: 24,
        fill: "none",
        viewBox: "0 0 24 24",
      },
    });

    h("path", {
      attr: {
        stroke: "currentColor",
        "stroke-linecap": "round",
        "stroke-linejoin": "round",
        "stroke-width": "2",
        d: "M12 13V8m0 8h.01M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z",
      },
    });
  });
};
