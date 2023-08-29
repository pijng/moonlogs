import { Event, Store, createEvent, is, sample } from "effector";

export const trigger = <T>(source: Store<any> | any, target: Event<any>) => {
  const result = createEvent<T>();
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
