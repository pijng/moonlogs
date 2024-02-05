import { Store } from "effector";
import { h } from "forest";

export const Subheader = (text: string | Store<string>) => {
  h("h5", {
    classList: ["text-xl", "text-slate-800", "font-bold", "dark:text-slate-200"],
    text: text,
  });
};
