import { Store, createStore } from "effector";
import { h } from "forest";

export const ErrorHint = (text: Store<string> | string | undefined, visible: Store<boolean> | undefined) => {
  h("p", {
    classList: ["mt-2", "text-sm", "text-red-600", "dark:text-red-400"],
    visible: visible || createStore(false),
    text: text || createStore(""),
  });
};
