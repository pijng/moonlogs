import { h, spec } from "forest";

export const ForwardStepIcon = () => {
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
        d: "M16 6v12M8 6v12l8-6-8-6Z",
      },
    });
  });
};
