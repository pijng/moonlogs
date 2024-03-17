import { Store } from "effector";
import { h, spec } from "forest";
import { QuestionIcon } from "../icons";

export const Popover = ({ text }: { text?: string | Store<string> }) => {
  h("div", () => {
    spec({ classList: ["group", "w-max"] });

    h("div", () => {
      spec({ classList: ["group-hover:underline", "cursor-pointer"] });
      QuestionIcon();
    });

    h("div", () => {
      spec({
        classList: [
          "absolute",
          "pointer-events-none",
          "bottom-1/3",
          "transform",
          "-translate-y-1/2",
          "ml-6",
          "bg-raisin-black",
          "text-white",
          "text-center",
          "px-2",
          "py-1",
          "rounded",
          "opacity-0",
          "group-hover:opacity-100",
          "transition-opacity",
          "duration-300",
        ],
      });

      h("div", () => {
        spec({ classList: ["text-left", "text-sm", "px-1", "py-0.5"] });

        h("p", {
          text: text,
        });
      });
    });
  });
};
