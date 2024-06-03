import { Store, createEvent, createStore, sample } from "effector";
import { DOMElement, h, node, spec } from "forest";
import { Button } from "../buttons";

export const Popup = ({ text, icon, content }: { text?: string | Store<string>; icon?: () => void; content: () => void }) => {
  const clicked = createEvent();
  const $visible = createStore(false);
  const outsideClicked = createEvent<{ node: DOMElement; event: any }>();

  sample({
    source: $visible,
    clock: clicked,
    fn: (v) => !v,
    target: $visible,
  });

  sample({
    source: $visible,
    clock: outsideClicked,
    filter: (visible, { node, event }) => !node.contains(event.target) && visible,
    fn: () => false,
    target: $visible,
  });

  h("div", () => {
    spec({ classList: ["relative"] });
    Button({
      text: text,
      preIcon: icon,
      variant: "alternative",
      size: "extra_small",
      event: clicked,
    });
    h("div", () => {
      spec({
        visible: $visible,
        classList: [
          "w-max",
          "top-0",
          "left-0",
          "w-5/6",
          "absolute",
          "left-0",
          "top-9",
          "border",
          "border-gray-200",
          "dark:bg-raisin-black",
          "dark:border-shadow-gray",
          "z-50",
          "bg-white",
          "divide-y",
          "divide-gray-100",
          "rounded-lg",
          "shadow",
          "dark:divide-gray-600",
          "shadow-2xl",
        ],
      });

      h("ul", () => {
        spec({
          classList: ["flex", "flex-col", "flex-wrap", "py-2", "text-sm", "text-gray-700", "dark:text-gray-200"],
        });

        content();
      });
    });

    node((node) => {
      document.addEventListener("click", (event) => {
        outsideClicked({ node, event });
      });
    });
  });
};
