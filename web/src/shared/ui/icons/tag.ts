import { h, spec } from "forest";

export const TagIcon = () => {
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
        d: "M15.2 6H4a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h11.2a1 1 0 0 0 .747-.334l4.46-5a1 1 0 0 0 0-1.332l-4.46-5A1 1 0 0 0 15.2 6Z",
      },
    });
  });
};
