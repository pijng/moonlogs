import { Tooltip } from "@/shared/ui";
import { createEvent, createStore, sample } from "effector";

export const triggerTooltip = createEvent<{ text: string; timeout?: number }>();

const $text = createStore<string>("");
const $timeout = createStore<number>(1000);

sample({
  source: triggerTooltip,
  fn: ({ text }) => text,
  target: $text,
});

sample({
  source: triggerTooltip,
  fn: ({ timeout }) => timeout || 1000,
  target: $timeout,
});

export const GeneralTooltip = () => {
  Tooltip({ text: $text, event: triggerTooltip, timeout: $timeout });
};
