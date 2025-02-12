import { h, spec } from "forest";

export const ChartLineUpIcon = () => {
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
        d: "M4 4.5V19a1 1 0 0 0 1 1h15M7 14l4-4 4 4 5-5m0 0h-3.207M20 9v3.207",
      },
    });
  });
};
