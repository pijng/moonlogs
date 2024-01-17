import { Event, Store, createEvent, is, sample } from "effector";

export const trigger = ({ source, target }: { source: Store<any>; target: Event<any> }) => {
  const result = createEvent();
  if (is.unit(source)) {
    sample({
      source,
      clock: result,
      target,
    });
  } else {
    sample({
      source: result,
      target,
      fn: () => source,
    });
  }
  return result;
};
