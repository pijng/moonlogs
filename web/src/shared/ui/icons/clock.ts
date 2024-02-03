import { h, spec } from "forest";

export const ClockIcon = () => {
  h("svg", () => {
    spec({
      classList: ["h-4", "w-4", "mr-2"],
      attr: { "aria-hidden": "true", xmlhs: "http://www.w3.org/2000/svg", fill: "currentColor", viewBox: "0 0 24 24" },
    });

    h("path", {
      attr: {
        "fill-rule": "evenodd",
        stroke: "currentColor",
        "clip-rule": "evenodd",
        d: "M2 12a10 10 0 1 1 20 0 10 10 0 0 1-20 0Zm11-4a1 1 0 1 0-2 0v4c0 .3.1.5.3.7l3 3a1 1 0 0 0 1.4-1.4L13 11.6V8Z",
      },
    });
  });
};
