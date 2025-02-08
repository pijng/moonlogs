import { i18n } from "@/shared/lib/i18n";
import { Change, diffWords } from "diff";
import { Store, createEffect, createEvent, restore, sample } from "effector";
import { h, list, node, remap, spec } from "forest";
import { triggerTooltip } from "../general-tooltip";
import { $shouldCopyToClipboard } from "@/shared/lib";

type ComparePayload = { oldText: string | undefined; newText: string | undefined };

export const DiffText = ({ oldText, newText }: { oldText: Store<string>; newText: Store<string> }) => {
  h("div", () => {
    const compareFx = createEffect(({ oldText, newText }: ComparePayload): Change[] => {
      return diffWords(JSON.stringify(oldText), JSON.stringify(newText));
    });

    const triggerDiff = createEvent();

    const diffFx = sample({
      source: [oldText, newText],
      clock: triggerDiff,
      fn: ([oldText, newText]) => ({ oldText, newText }),
      target: compareFx,
    });

    const $before = restore(diffFx, []).map((parts) => parts.filter((p) => !p.added));
    const $after = restore(diffFx, []).map((parts) => parts.filter((p) => !p.removed));

    const beforeClicked = createEvent<MouseEvent>();
    const afterClicked = createEvent<MouseEvent>();
    const copyTextFx = createEffect((clickedText: string) => {
      return navigator.clipboard.writeText(clickedText);
    });

    sample({
      source: i18n("miscellaneous.copied_to_clipboard"),
      clock: copyTextFx.done,
      fn: (text) => ({ text: text }),
      target: triggerTooltip,
    });

    spec({ classList: ["grid", "grid-cols-2"] });

    h("div", () => {
      spec({
        classList: ["border-r-2", "pr-4", "border-dashed"],
        handler: { on: { click: beforeClicked } },
      });

      sample({
        source: [$before.map((parts) => parts.map((p) => p.value).join("")), $shouldCopyToClipboard] as const,
        clock: beforeClicked,
        filter: ([, shouldCopy]) => !Boolean(window.getSelection()?.toString()) && shouldCopy,
        fn: ([text]) => text,
        target: copyTextFx,
      });

      list($before, ({ store: before }) => {
        h("span", {
          text: remap(before, "value"),
          classList: {
            "whitespace-pre-wrap": true,
            "break-words": true,
            "cursor-pointer": $shouldCopyToClipboard,
            "bg-red-200": before.map((p) => Boolean(p.removed)),
            "dark:bg-red-800": before.map((p) => Boolean(p.removed)),
          },
        });
      });
    });

    h("div", () => {
      spec({
        classList: ["border-l-2", "pl-4", "border-dashed"],
        handler: { on: { click: afterClicked } },
      });

      sample({
        source: $after.map((parts) => parts.map((p) => p.value).join("")),
        clock: afterClicked,
        filter: () => !Boolean(window.getSelection()?.toString()),
        target: copyTextFx,
      });

      list($after, ({ store: after }) => {
        h("span", {
          text: remap(after, "value"),
          classList: {
            "whitespace-pre-wrap": true,
            "break-words": true,
            "cursor-pointer": $shouldCopyToClipboard,
            "bg-green-200": after.map((p) => Boolean(p.added)),
            "dark:bg-green-800": after.map((p) => Boolean(p.added)),
          },
        });
      });
    });
    node(() => {
      triggerDiff();
    });
  });
};
