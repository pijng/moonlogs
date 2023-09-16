import { Button, Header, PlusIcon } from "@/shared/ui";
import { RouteInstance } from "atomic-router";
import { Store, createEvent, sample } from "effector";
import { h, spec } from "forest";

export const HeaderWithCreation = (text: string | Store<string>, route: RouteInstance<any>) => {
  const routeOpened = createEvent();
  sample({
    clock: routeOpened,
    target: route.open,
  });

  h("div", () => {
    spec({
      classList: ["flex", "items-center"],
    });
    Header(text);

    h("div", () => {
      spec({
        classList: ["ml-2.5"],
      });

      Button({
        text: "",
        variant: "default",
        style: "round",
        size: "extra_small",
        event: routeOpened,
        preIcon: PlusIcon,
      });
    });
  });
};
