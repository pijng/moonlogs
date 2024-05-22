import { actionModel } from "@/entities/action";
import { actionEditRoute } from "@/shared/routing";
import { ActionsTable } from "@/shared/ui";
import { createEvent, sample } from "effector";
import { h, spec } from "forest";

export const ActionsList = () => {
  h("div", () => {
    spec({
      classList: ["flex", "flex-col", "space-y-6", "mt-3"],
    });

    const editActionClicked = createEvent<number>();
    sample({
      source: editActionClicked,
      fn: (id) => ({ id }),
      target: actionEditRoute.open,
    });

    ActionsTable(actionModel.$actions, editActionClicked);
  });
};
