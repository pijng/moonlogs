import { h, spec } from "forest";

export const CaretDownIcon = () => {
  h("svg", () => {
    spec({
      classList: ["h-4", "w-4"],
      attr: {
        "aria-hidden": "true",
        xmlhs: "http://www.w3.org/2000/svg",
        fill: "currentColor",
        viewBox: "0 0 24 24",
      },
    });

    h("path", {
      attr: {
        "fill-rule": "evenodd",
        "clip-rule": "evenodd",
        d: "M18.425 10.271C19.499 8.967 18.57 7 16.88 7H7.12c-1.69 0-2.618 1.967-1.544 3.271l4.881 5.927a2 2 0 0 0 3.088 0l4.88-5.927Z",
      },
    });
  });
};
