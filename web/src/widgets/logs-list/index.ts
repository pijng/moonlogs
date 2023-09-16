import { logModel } from "@/entities/log";
import { CardHeaded, LogsTable } from "@/shared/ui";
import { h, list, spec } from "forest";

export const LogsList = () => {
  h("div", () => {
    spec({
      classList: ["flex", "flex-col", "space-y-6", "mt-2"],
    });

    list(logModel.$logsGroups, ({ store: group }) => {
      CardHeaded({
        tags: group.map((g) => g.tags),
        href: group.map((g) => `${g.schema_name}/${g.group_hash}`),
        content: () => {
          LogsTable(group.map((g) => g.logs));
        },
        withMore: true,
      });
    });
  });
};
