import { Event, Store, createEvent, createStore, sample } from "effector";
import { h, spec } from "forest";

export const KBD = ({
  text,
  event,
  visible,
}: {
  text: Store<string | null> | Store<string> | string;
  event?: Event<any>;
  visible?: Store<boolean>;
}) => {
  const $localVisible = visible ?? createStore(true);
  const localEvent = event ?? createEvent();
  const kbdClicked = createEvent<MouseEvent>();

  sample({
    clock: kbdClicked,
    target: localEvent,
  });

  h("li", () => {
    spec({
      visible: $localVisible,
      handler: { on: { click: kbdClicked } },
    });

    h("kbd", {
      classList: {
        block: true,
        "px-2": true,
        "py-1.5": true,
        "text-xs": true,
        "font-semibold": true,
        "text-gray-800": true,
        "cursor-pointer": Boolean(event),
        "bg-gray-100": true,
        border: true,
        "border-gray-200": true,
        "rounded-lg": true,
        "dark:bg-slate-gray": true,
        "dark:text-gray-200": true,
        "dark:border-none": true,
      },
      text: text,
    });
  });
};
