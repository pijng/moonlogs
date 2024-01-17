import { h, spec } from "forest";

export const CrossIcon = () => {
  h("svg", () => {
    spec({
      classList: ["w-5", "h-5", "ml-2.5"],
      attr: { xmlhs: "http://www.w3.org/2000/svg", fill: "none", viewBox: "0 0 20 20" },
    });

    h("path", {
      attr: {
        stroke: "currentColor",
        fill: "none",
        "fill-rule": "evenodd",
        "stroke-linecap": "round",
        "stroke-linejoin": "round",
        d: "M10 10l5.09-5.09L10 10l5.09 5.09L10 10zm0 0L4.91 4.91 10 10l-5.09 5.09L10 10z",
      },
    });
  });
};
