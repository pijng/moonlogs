import { Event, Store, createEffect, createEvent, createStore, sample } from "effector";
import { h } from "forest";

export const Tooltip = ({ text, event, timeout }: { text: Store<string> | string; event: Event<any>; timeout?: number }) => {
  const hide = createEvent<any>();
  const $visible = createStore(false)
    .on(event, () => true)
    .on(hide, () => false);

  const scheduleHideFx = createEffect(() => {
    setTimeout(() => {
      hide();
    }, timeout || 1000);
  });

  sample({
    clock: event,
    target: scheduleHideFx,
  });

  h("div", {
    visible: $visible,
    attr: {
      role: "tooltip",
    },
    classList: [
      "absolute",
      "top-1/2",
      "z-10",
      "inline-block",
      "px-3",
      "py-2",
      "text-sm",
      "font-medium",
      "text-white",
      "bg-gray-900",
      "rounded-lg",
      "shadow-sm",
      "tooltip",
      "dark:bg-gray-700",
    ],
    text: text,
  });
};
