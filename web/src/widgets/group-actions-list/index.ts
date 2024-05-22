import { actionModel } from "@/entities/action";
import { logModel } from "@/entities/log";
import { LogGroupAction } from "@/features";
import { h, list, spec } from "forest";

export const GroupActionsList = () => {
  h("div", () => {
    spec({
      classList: ["inline-flex", "space-x-3", "pt-6", "pb-3"],
    });

    list(actionModel.$actions, ({ store: action }) => {
      LogGroupAction({ action: action, logGroup: logModel.$groupedLogs.map((g) => g.logs) });
    });
  });
};
