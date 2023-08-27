import { h, spec } from "forest";

export const DownIcon = () => {
  h("svg", () => {
    spec({
      classList: ["w-2.5", "h-2.5", "ml-2.5"],
      attr: { "aria-hidden": "true", xmlhs: "http://www.w3.org/2000/svg", fill: "none", viewBox: "0 0 10 6" },
    });

    h("path", {
      attr: {
        stroke: "currentColor",
        "stroke-linecap": "round",
        "stroke-linejoin": "round",
        "stroke-width": "2",
        d: "m1 1 4 4 4-4",
      },
    });
  });
};
