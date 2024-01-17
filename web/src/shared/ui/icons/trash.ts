import { h, spec } from "forest";

export const TrashIcon = () => {
  h("svg", () => {
    spec({
      classList: ["w-5", "h-5"],
      attr: {
        "aria-hidden": true,
        xmlns: "http://www.w3.org/2000/svg",
        fill: "none",
        viewBox: "0 0 18 20",
      },
    });

    h("path", {
      attr: {
        stroke: "currentColor",
        "stroke-linecap": "round",
        "stroke-linejoin": "round",
        "stroke-width": "2",
        d: "M1 5h16M7 8v8m4-8v8M7 1h4a1 1 0 0 1 1 1v3H6V2a1 1 0 0 1 1-1ZM3 5h12v13a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V5Z",
      },
    });
  });
};
