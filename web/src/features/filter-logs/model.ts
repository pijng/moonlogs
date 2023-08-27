import { layoutClicked } from "@/shared/ui";
import { createEvent, createStore, restore, sample } from "effector";

export const filterClicked = createEvent<any>();
export const $filterIsOpened = createStore(false);

sample({
  source: $filterIsOpened,
  clock: filterClicked,
  fn: (state) => !state,
  target: $filterIsOpened,
});

sample({
  source: [$filterIsOpened, restore(filterClicked, null)],
  clock: layoutClicked,
  filter: ([isOpened, clicked], layoutClicked) => {
    const path = layoutClicked.composedPath();

    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore
    return !path.includes(clicked?.target?.parentNode) && isOpened;
  },
  target: filterClicked,
});
