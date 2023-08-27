import { h, spec } from "forest";

// <svg xmlns="http://www.w3.org/2000/svg" aria-hidden="true" class="h-4 w-4 mr-2 text-gray-400" viewBox="0 0 20 20" fill="currentColor">
//     <path fill-rule="evenodd" d="M3 3a1 1 0 011-1h12a1 1 0 011 1v3a1 1 0 01-.293.707L12 11.414V15a1 1 0 01-.293.707l-2 2A1 1 0 018 17v-5.586L3.293 6.707A1 1 0 013 6V3z" clip-rule="evenodd"></path>
// </svg>

export const FilterIcon = () => {
  h("svg", () => {
    spec({
      classList: ["h-4", "w-4", "mr-2", "text-gray-400"],
      attr: { "aria-hidden": "true", xmlhs: "http://www.w3.org/2000/svg", fill: "currentColor", viewBox: "0 0 20 20" },
    });

    h("path", {
      attr: {
        "fill-rule": "evenodd",
        "clip-rule": "evenodd",
        d: "M3 3a1 1 0 011-1h12a1 1 0 011 1v3a1 1 0 01-.293.707L12 11.414V15a1 1 0 01-.293.707l-2 2A1 1 0 018 17v-5.586L3.293 6.707A1 1 0 013 6V3z",
      },
    });
  });
};
