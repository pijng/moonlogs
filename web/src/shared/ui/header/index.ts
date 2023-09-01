import { Store } from "effector";
import { h } from "forest";

export const Header = (text: string | Store<string>) => {
  h("h1", {
    classList: [
      "inline-block",
      "text-2xl",
      "sm:text-3xl",
      "font-extrabold",
      "text-slate-900",
      "tracking-tight",
      "dark:text-slate-200",
      "pb-3",
    ],
    text: text,
  });
};
