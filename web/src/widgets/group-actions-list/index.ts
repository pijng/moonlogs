import { actionModel } from "@/entities/action";
import { logModel } from "@/entities/log";
import { LogGroupAction } from "@/features";
import { Log } from "@/shared/api";
import { Popup } from "@/shared/ui";
import { Store, createStore } from "effector";
import { h, list, spec, variant } from "forest";

type Style = "default" | "alternative" | "light";

export const GroupActionsList = ({
  logGroup,
  style,
  collapsed,
}: {
  logGroup?: Store<Log[]>;
  style?: Style;
  collapsed?: boolean;
}) => {
  const logs = logGroup || logModel.$groupedLogs.map((g) => g.logs);
  const $displayStyle = createStore<{ style: "collapsed" | "default" }>({ style: Boolean(collapsed) ? "collapsed" : "default" });

  variant({
    source: $displayStyle,
    key: "style",
    cases: {
      collapsed: () => {
        Popup({
          text: "•••",
          content: () => {
            list(actionModel.$actions, ({ store: action }) => {
              h("li", () => {
                spec({
                  classList: ["block", "flex-auto", "shrink-0"],
                });

                LogGroupAction({ action: action, logGroup: logs, style: "light" });
              });
            });
          },
        });
      },
      default: () => {
        list(actionModel.$actions, ({ store: action }) => {
          LogGroupAction({ action: action, logGroup: logs, style: style });
        });
      },
    },
  });
};
