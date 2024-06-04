import { h, spec } from "forest";

export const LockOpenIcon = () => {
  h("svg", () => {
    spec({
      classList: ["w-6", "h-6"],
      attr: { "aria-hidden": "true", xmlhs: "http://www.w3.org/2000/svg", fill: "none", viewBox: "0 0 24 24" },
    });

    h("path", {
      attr: {
        stroke: "currentColor",
        "stroke-linecap": "round",
        "stroke-linejoin": "round",
        "stroke-width": "2",
        d: "M10 14v3m4-6V7a3 3 0 1 1 6 0v4M5 11h10a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-7a1 1 0 0 1 1-1Z",
      },
    });
  });
};
