import { userModel } from "@/entities/user";
import { memberEditRoute } from "@/routing/shared";
import { UsersTable } from "@/shared/ui";
import { createEvent, sample } from "effector";
import { h, spec } from "forest";

export const UsersList = () => {
  h("div", () => {
    spec({
      classList: ["flex", "flex-col", "space-y-6", "mt-3"],
    });

    const editUserClicked = createEvent<number>();
    sample({
      source: editUserClicked,
      fn: (id) => ({ id }),
      target: memberEditRoute.open,
    });

    UsersTable(userModel.$users, editUserClicked);
  });
};
