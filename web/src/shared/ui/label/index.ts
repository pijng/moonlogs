import { Store } from "effector";
import { h, spec } from "forest";
import { Popover } from "../popover";

export const Label = ({ text, hint }: { text: string | Store<string>; hint: string }) => {
  h("div", () => {
    spec({ classList: ["flex", "items-center"] });

    h("p", {
      classList: ["block", "text-sm", "font-medium", "text-gray-900", "dark:text-white"],
      text: text,
    });

    h("div", () => {
      spec({ classList: ["pl-1"] });

      Popover({ text: hint });
    });
  });
};
