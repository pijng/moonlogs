import { Store } from "effector";
import { h } from "forest";

export const Subheader = (text: string | Store<string>) => {
  h("h5", {
    classList: ["text-xl", "font-bold", "dark:text-white"],
    text: text,
  });
};
