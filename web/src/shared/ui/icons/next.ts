import { h, spec } from "forest";

export const NextIcon = () => {
  h("svg", () => {
    spec({
      classList: ["w-2.5", "h-2.5"],
      attr: { "aria-hidden": true, xmlns: "http://www.w3.org/2000/svg", fill: "none", viewBox: "0 0 6 10" },
    });

    h("path", {
      attr: {
        stroke: "currentColor",
        "stroke-linecap": "round",
        "stroke-linejoin": "round",
        "stroke-width": "2",
        d: "m1 9 4-4-4-4",
      },
    });
  });
};
