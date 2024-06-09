import { Change, diffWords } from "diff";
import { Store, createEffect, createEvent, restore, sample } from "effector";
import { h, list, node, remap } from "forest";

const compareFx = createEffect(({ oldText, newText }: { oldText: string; newText: string }): Change[] => {
  return diffWords(oldText, newText);
});

const triggerDiff = createEvent();

export const DiffText = ({ oldText, newText }: { oldText: Store<string>; newText: Store<string> }) => {
  sample({
    source: [oldText, newText],
    clock: triggerDiff,
    fn: ([oldText, newText]) => ({ oldText, newText }),
    target: compareFx,
  });

  const $changes = restore(compareFx, []);

  list($changes, ({ store: change }) => {
    h("span", {
      text: remap(change, "value"),
      classList: {
        "whitespace-pre-wrap": true,
        "break-words": true,
        "cursor-pointer": true,
        "bg-green-200": change.map((c) => Boolean(c.added)),
        "dark:bg-green-800": change.map((c) => Boolean(c.added)),
        "bg-red-200": change.map((c) => Boolean(c.removed)),
        "dark:bg-red-800": change.map((c) => Boolean(c.removed)),
      },
    });
  });

  node(() => {
    triggerDiff();
  });
};
