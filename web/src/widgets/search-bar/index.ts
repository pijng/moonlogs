import { logModel } from "@/entities/log";
import { schemaModel } from "@/entities/schema";
import { Filter } from "@/features";
import { logsRoute } from "@/routing/shared";
import { Button, Search } from "@/shared/ui";
import { combine, createEvent, sample } from "effector";
import { h, spec } from "forest";

const $currentSchema = combine([logsRoute.$params, schemaModel.$schemas], ([params, schemas]) => {
  return schemas.find((s) => s.name === params.schemaName) || null;
});

const $currentFilter = combine([$currentSchema, logModel.$formattedSearchFilter], ([schema, searchFilter]) => {
  const filter = JSON.parse(searchFilter);
  const fields = schema?.fields || [];

  return fields.map((f) => {
    const value = filter[f.name] || "";
    return { name: f.title, key: f.name, value: value };
  });
});

export const SearchBar = () => {
  h("div", () => {
    spec({
      classList: ["pb-4"],
    });

    h("div", () => {
      spec({
        classList: ["flex", "max-w-2xl", "w-full", "flex-row"],
      });

      h("div", () => {
        spec({
          classList: ["w-full"],
        });

        Search(logModel.events.queryChanged, logModel.$searchQuery);
      });
    });

    h("div", () => {
      spec({
        classList: ["flex", "relative"],
      });

      Filter($currentFilter, logModel.events.filterChanged);

      const filterCleared = createEvent<MouseEvent>();
      sample({
        clock: filterCleared,
        fn: () => ({}),
        target: logModel.events.resetFilter,
      });

      const $filtersApplied = $currentFilter.map((items) => items.filter((item) => item.value.trim().length > 0).length > 0);

      Button({
        text: "Clear",
        variant: "light",
        size: "small",
        event: filterCleared,
        visible: $filtersApplied,
      });
    });
  });
};
