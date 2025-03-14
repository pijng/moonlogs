import { h, spec } from "forest";

export const BullhornIcon = () => {
  h("svg", () => {
    spec({
      classList: ["h-6", "w-6", "text-gray-800", "dark:text-white"],
      attr: { "aria-hidden": "true", xmlhs: "http://www.w3.org/2000/svg", fill: "none", viewBox: "0 0 22 22" },
    });

    h("path", {
      attr: {
        stroke: "currentColor",
        "stroke-linecap": "round",
        "stroke-linejoin": "round",
        "stroke-width": "2",
        d: "M11 9H5a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h6m0-6v6m0-6 5.419-3.87A1 1 0 0 1 18 5.942v12.114a1 1 0 0 1-1.581.814L11 15m7 0a3 3 0 0 0 0-6M6 15h3v5H6v-5Z",
      },
    });
  });
};
