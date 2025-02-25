import { createStore, is, Store } from "effector";
import { h, spec } from "forest";
import { Spinner } from "@/shared/ui";

// <div id="toast-success" class="flex items-center w-full max-w-xs p-4 mb-4 text-gray-500 bg-white rounded-lg shadow-sm dark:text-gray-400 dark:bg-gray-800" role="alert">
//     <div class="inline-flex items-center justify-center shrink-0 w-8 h-8 text-green-500 bg-green-100 rounded-lg dark:bg-green-800 dark:text-green-200">
//         <svg class="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20">
//             <path d="M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5Zm3.707 8.207-4 4a1 1 0 0 1-1.414 0l-2-2a1 1 0 0 1 1.414-1.414L9 10.586l3.293-3.293a1 1 0 0 1 1.414 1.414Z"/>
//         </svg>
//         <span class="sr-only">Check icon</span>
//     </div>
//     <div class="ms-3 text-sm font-normal">Item moved successfully.</div>
//     <button type="button" class="ms-auto -mx-1.5 -my-1.5 bg-white text-gray-400 hover:text-gray-900 rounded-lg focus:ring-2 focus:ring-gray-300 p-1.5 hover:bg-gray-100 inline-flex items-center justify-center h-8 w-8 dark:text-gray-500 dark:hover:text-white dark:bg-gray-800 dark:hover:bg-gray-700" data-dismiss-target="#toast-success" aria-label="Close">
//         <span class="sr-only">Close</span>
//         <svg class="w-3 h-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 14">
//             <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"/>
//         </svg>
//     </button>
// </div>

export const Infobox = ({ text, emoji }: { text: string | Store<string>; emoji: Store<string> | string }) => {
  const $text = is.unit(text) ? text : createStore(text);

  h("div", () => {
    spec({
      classList: [
        "flex",
        "items-center",
        "w-full",
        "p-4",
        "mb-4",
        "rounded-lg",
        "border",
        "shadow",
        "dark:border-shadow-gray",
        "shadow-sm",
        "dark:bg-slate-gray",
        "dark:text-gray-200",
        "text-gray-900",
      ],
    });

    h("div", () => {
      spec({
        classList: ["inline-flex", "items-center", "justify-center", "shrink-0", "w-8", "h-8", "text-2xl", "rounded-lg"],
        text: emoji,
      });
    });

    h("div", () => {
      spec({ classList: ["absolute", "left-1/2", "max-w-[2rem]"] });
      Spinner({ visible: $text.map((t) => t.length <= 0) });
    });

    h("div", () => {
      spec({
        visible: $text.map((t) => t.length > 0),
        classList: ["whitespace-pre-wrap", "ms-3", "text-base", "font-normal"],
        text: text,
      });
    });
  });
};
