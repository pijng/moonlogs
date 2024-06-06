import { Action, Log } from "@/shared/api";
import { createLogGroupAction } from "./model";
import { Button, ButtonSize } from "@/shared/ui";
import { h, remap, spec } from "forest";
import { Store, createEffect, createEvent, sample } from "effector";

type Style = "default" | "alternative" | "light";

const sizes: Record<Style, ButtonSize> = {
  default: "extra_small",
  alternative: "extra_small",
  light: "small",
};

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

    h("a", () => {
      spec({
        attr: { href: actionModel.$link },
      });

      Button({
        text: remap(action, "name"),
        variant: style,
        prevent: true,
        size: sizes[style],
        event: actionClicked,
      });
    });
  });
};
