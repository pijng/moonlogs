import { userModel } from "@/entities/user";
import { Table } from "@/shared/ui";
import { createStore } from "effector";
import { h, spec } from "forest";

export const UsersList = () => {
  h("div", () => {
    spec({
      classList: ["flex", "flex-col", "space-y-6", "mt-3"],
    });

    Table({
      columns: createStore(["Email", "Name", "Role"]),
      rows: userModel.$formattedUsers,
    });
  });
};
