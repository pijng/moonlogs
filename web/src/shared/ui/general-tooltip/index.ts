import { Tooltip } from "@/shared/ui";
import { Store, createEvent, createStore, is, sample } from "effector";

export const triggerTooltip = createEvent<{ text: Store<string> | string; timeout?: Store<number> | number }>();

const $text = createStore<string>("");
const $timeout = createStore<number>(1000);

sample({
  source: triggerTooltip,
  //                                       BRUH
  fn: ({ text }) => (is.store(text) ? text.getState() : text),
  target: $text,
});

sample({
  source: triggerTooltip,
  //                                                BRUH
  fn: ({ timeout }) => (is.store(timeout) ? timeout.getState() : timeout),
  target: $timeout,
});

export const GeneralTooltip = () => {
  Tooltip({ text: $text, event: triggerTooltip, timeout: $timeout });
};
