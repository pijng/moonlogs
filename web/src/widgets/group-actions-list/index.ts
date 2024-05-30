import { actionModel } from "@/entities/action";
import { logModel } from "@/entities/log";
import { LogGroupAction } from "@/features";
import { Log } from "@/shared/api";
import { Store } from "effector";
import { list } from "forest";

export const GroupActionsList = ({ logGroup, primary }: { logGroup?: Store<Log[]>; primary?: boolean }) => {
  const logs = logGroup || logModel.$groupedLogs.map((g) => g.logs);

  list(actionModel.$actions, ({ store: action }) => {
    LogGroupAction({ action: action, logGroup: logs, primary: primary });
  });
};
