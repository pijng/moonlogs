import { h, spec } from "forest";

export const PlusIcon = () => {
  h("svg", () => {
    spec({
      classList: ["w-3", "h-3"],
      attr: {
        "aria-hidden": true,
        xmlns: "http://www.w3.org/2000/svg",
        fill: "currentColor",
        viewBox: "0 0 24 24",
      },
    });

    h("path", {
      attr: {
        stroke: "currentColor",
        "stroke-linecap": "round",
        "stroke-linejoin": "round",
        "stroke-width": "1",
        d: "M24 10h-10v-10h-4v10h-10v4h10v10h4v-10h10z",
      },
    });
  });
};
