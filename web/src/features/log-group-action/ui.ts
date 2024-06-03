import { Action, Log } from "@/shared/api";
import { createLogGroupAction } from "./model";
import { Button } from "@/shared/ui";
import { h, remap, spec } from "forest";
import { Store, createEffect, createEvent, sample } from "effector";

type Style = "default" | "alternative" | "light";

export const LogGroupAction = ({
  action,
  logGroup,
  style = "default",
}: {
  action: Store<Action>;
  logGroup: Store<Log[]>;
  style?: Style;
}) => {
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
      variant: style,
      prevent: true,
      size: "extra_small",
      event: actionClicked,
    });
  });
};
