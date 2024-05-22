import { Store } from "effector";
import { h } from "forest";

export const Text = ({ text }: { text?: string | Store<string> }) => {
  h("p", {
    classList: ["block", "mb-2", "text-sm", "font-medium", "text-gray-900", "dark:text-white"],
    text: text,
  });
};
