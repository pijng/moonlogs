import { h, spec } from "forest";

export const LightBulbIcon = () => {
  h("svg", () => {
    spec({
      classList: ["w-6", "h-6", "text-gray-800", "dark:text-white"],
      attr: {
        "aria-hidden": "true",
        xmlhs: "http://www.w3.org/2000/svg",
        fill: "none",
        width: 24,
        height: 24,
        viewBox: "0 0 24 24",
      },
    });

    h("path", {
      attr: {
        stroke: "currentColor",
        "stroke-linecap": "round",
        "stroke-linejoin": "round",
        "stroke-width": "2",
        d: "M9 9a3 3 0 0 1 3-3m-2 15h4m0-3c0-4.1 4-4.9 4-9A6 6 0 1 0 6 9c0 4 4 5 4 9h4Z",
      },
    });
  });
};
