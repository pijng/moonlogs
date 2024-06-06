import { logModel } from "@/entities/log";
import { schemaModel } from "@/entities/schema";
import { Level } from "@/shared/api";
import { queryStringToObject } from "@/shared/lib";
import { i18n } from "@/shared/lib/i18n";
import { logsRoute } from "@/shared/routing";
import { FilledFloatingInput, Input, Select, Text } from "@/shared/ui";
import { combine, createEvent, createStore, sample } from "effector";
import { h, list, spec } from "forest";

export const LogsFilter = () => {
  const $currentSchema = combine([logsRoute.$params, schemaModel.$schemas], ([params, schemas]) => {
    const currentSchema = schemas.find((s) => s.name === params.schemaName);
    if (currentSchema) return currentSchema;

    return { id: 0, title: "", description: "", name: "", fields: [], kinds: [] };
  });

  const $currentFilter = combine([$currentSchema, logModel.$formattedSearchFilter], ([schema, searchFilter]) => {
    const filter = queryStringToObject(searchFilter);
    const fields = schema?.fields || [];

    return fields.map((f) => {
      const value = filter[f.name] || "";
      return { name: f.title, key: f.name, value: value };
    });
  });

  h("div", () => {
    spec({
      classList: ["sticky", "top-5", "flex", "flex-col"],
    });

    // Filter by query and kind
    h("div", () => {
      spec({ classList: ["flex", "items-center", "w-full"] });

      h("ul", () => {
        spec({
          classList: ["w-full", "flex", "flex-col", "gap-y-2", "text-sm", "text-gray-700", "dark:text-gray-200"],
        });

        const $allKinds = $currentSchema.map((s) => s.kinds || []);

        h("li", () => {
          spec({
            visible: $allKinds.map((kinds) => kinds?.length > 0),
            classList: ["block", "pb-3", "flex-auto", "shrink-0"],
          });

          Select({
            value: logModel.$currentKind,
            text: i18n("log_groups.filters.query.kind"),
            options: $allKinds,
            optionSelected: logModel.events.kindChanged,
          });
        });

        h("li", () => {
          Text({ text: i18n("log_groups.filters.query.label") });
        });

        list($currentFilter, ({ store: item }) => {
          h("li", () => {
            spec({
              classList: ["block", "pb-2", "flex-auto", "shrink-0"],
            });

            const inputChanged = createEvent<InputEvent>();

            sample({
              source: item,
              clock: inputChanged,
              fn: (item, event) => {
                // eslint-disable-next-line @typescript-eslint/ban-ts-comment
                // @ts-ignore
                return { [item.key]: event.target.value };
              },
              target: logModel.events.filterChanged,
            });

            FilledFloatingInput(item, inputChanged);
          });
        });
      });
    });

    h("div", () => {
      spec({
        classList: ["flex", "flex-col", "mt-5", "w-full"],
      });

      Select({
        value: logModel.$currentLevel,
        text: i18n("log_groups.filters.level.label"),
        options: createStore<Level[]>(["Info", "Error", "Warn", "Debug", "Trace", "Fatal"]),
        optionSelected: logModel.events.levelChanged,
      });
    });

    h("div", () => {
      spec({
        classList: ["flex", "items-center", "mt-6", "w-full"],
      });

      h("ul", () => {
        spec({
          classList: ["w-full", "flex", "flex-col", "flex-wrap", "gap-y-3", "text-sm", "text-gray-700", "dark:text-gray-200"],
        });

        h("li", () => {
          spec({
            classList: ["block", "flex-auto", "w-full"],
          });

          Input({
            type: "datetime-local",
            value: logModel.$currentFromTime,
            label: i18n("log_groups.filters.time.from"),
            inputChanged: logModel.events.fromTimeChanged,
          });
        });

        h("li", () => {
          spec({
            classList: ["block", "flex-auto", "w-full"],
          });

          Input({
            type: "datetime-local",
            value: logModel.$currentToTime,
            label: i18n("log_groups.filters.time.to"),
            inputChanged: logModel.events.toTimeChanged,
          });
        });
      });
    });
  });
};
