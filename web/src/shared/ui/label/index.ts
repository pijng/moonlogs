import { Store, createStore, is } from "effector";
import { h, spec } from "forest";
import { Popover } from "../popover";

export const Label = ({ text, hint }: { text: string | Store<string>; hint: Store<string> | string | undefined }) => {
  h("div", () => {
    spec({ classList: ["flex", "items-center"] });

    h("p", {
      classList: ["block", "text-sm", "font-medium", "text-gray-900", "dark:text-white"],
      text: text,
    });

    h("div", () => {
      const $visible = is.store(hint) ? hint.map(Boolean) : createStore(Boolean(hint));

      spec({ classList: ["pl-1"], visible: $visible });

      Popover({ text: hint });
    });
  });
};
