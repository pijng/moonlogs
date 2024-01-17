import { PermissionGate } from "@/shared/lib";
import { Button, Header, PlusIcon } from "@/shared/ui";
import { RouteInstance, redirect } from "atomic-router";
import { Store, createEvent } from "effector";
import { h, spec } from "forest";

export const HeaderWithCreation = (text: string | Store<string>, route: RouteInstance<any>) => {
  const routeOpened = createEvent();
  redirect({
    clock: routeOpened,
    route: route,
  });

  h("div", () => {
    spec({
      classList: ["flex", "items-center"],
    });

    Header(text);

    PermissionGate("Admin", () => {
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
  });
};
