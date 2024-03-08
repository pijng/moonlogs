import { userModel } from "@/entities/user";
import { memberEditRoute } from "@/shared/routing";
import { UsersTable } from "@/shared/ui";
import { redirect } from "atomic-router";
import { createEvent } from "effector";
import { h, spec } from "forest";

export const UsersList = () => {
  h("div", () => {
    spec({
      classList: ["flex", "flex-col", "space-y-6", "mt-3"],
    });

    const editUserClicked = createEvent<number>();
    redirect({
      clock: editUserClicked,
      params: (id) => ({ id: id }),
      route: memberEditRoute,
    });

    UsersTable(userModel.$users, editUserClicked);
  });
};
