import { Action, Log } from "@/shared/api";
import { createLogGroupAction } from "./model";
import { Button } from "@/shared/ui";
import { h, remap, spec } from "forest";
import { Store, createEffect, createEvent, sample } from "effector";

export const LogGroupAction = ({ action, logGroup }: { action: Store<Action>; logGroup: Store<Log[]> }) => {
  const actionModel = createLogGroupAction(action, logGroup);

  const actionClicked = createEvent();
  const openActionLink = createEffect((actionLink: string) => {
    window.open(actionLink, "_blank");
  });

  sample({
    source: actionModel.$link,
    clock: actionClicked,
    target: openActionLink,
  });

  h("div", () => {
    spec({ visible: actionModel.$evaluated });

    Button({
      text: remap(action, "name"),
      variant: "default",
      prevent: true,
      size: "extra_small",
      event: actionClicked,
    });
  });
};
