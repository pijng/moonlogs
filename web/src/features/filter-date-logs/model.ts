import { layoutClicked } from "@/shared/ui";
import { createEvent, createStore, restore, sample } from "effector";

export const filterDateClicked = createEvent<any>();
export const $filterDateIsOpened = createStore(false);

sample({
  source: $filterDateIsOpened,
  clock: filterDateClicked,
  fn: (state) => !state,
  target: $filterDateIsOpened,
});

sample({
  source: [$filterDateIsOpened, restore(filterDateClicked, null)],
  clock: layoutClicked,
  filter: ([isOpened, clicked], layoutClicked) => {
    const path = layoutClicked.composedPath();

    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore
    return !path.includes(clicked?.target?.parentNode) && isOpened;
  },
  target: filterDateClicked,
});
