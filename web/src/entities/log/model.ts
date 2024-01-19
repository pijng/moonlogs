import { Log, getLogs } from "@/shared/api";
import { getLogGroup } from "@/shared/api/logs";
import { DATEFORMAT_OPTIONS, getLocale } from "@/shared/lib";
import { createEffect, createEvent, createStore, sample } from "effector";

type LogsGroup = { tags: Array<[string, any]>; schema_name: string; group_hash: string; logs: Log[] };

const reset = createEffect();

const resetSearch = createEvent();
export const $searchQuery = createStore("").reset(resetSearch);

const queryChanged = createEvent<string>();
$searchQuery.on(queryChanged, (_, query) => query);

const resetFilter = createEvent();

const filterChanged = createEvent<Record<string, any>>();
export const $formattedSearchFilter = createStore<string>("{}").reset(resetFilter);

$formattedSearchFilter.on(filterChanged, (formattedFilter, changedFilter) => {
  const filter = JSON.parse(formattedFilter);
  const newFilter = { ...filter, ...changedFilter };

  return JSON.stringify(newFilter);
});

const kindChanged = createEvent<string>();
export const $currentKind = createStore("").reset(resetFilter);
$currentKind.on(kindChanged, (_, kind) => kind);

const resetPage = createEvent();
const pageChanged = createEvent<string>();

export const $currentPage = createStore("1")
  .on(pageChanged, (_, newPage) => newPage)
  .reset([queryChanged, filterChanged, kindChanged, resetPage]);

const getLogsFx = createEffect((schema_name: string) => {
  return getLogs({ schema_name: schema_name });
});

const queryLogsFx = createEffect(
  ({
    schema_name,
    text,
    query,
    kind,
    page,
  }: {
    schema_name: string;
    text?: string;
    query?: string;
    kind?: string;
    page?: number;
  }) => {
    const objectQuery = JSON.parse(query || "{}") as Record<string, string>;

    const formattedQuery = Object.entries(objectQuery).reduce((acc, [k, v]) => {
      if (v.trim().length > 0) {
        return { ...acc, [k]: v };
      }

      return acc;
    }, {});

    return getLogs({ schema_name: schema_name, text: text, query: formattedQuery, kind: kind, page: page });
  },
);

const getLogGroupFx = createEffect(({ schema_name, hash }: { schema_name: string; hash: string }) => {
  return getLogGroup({ schema_name: schema_name, hash: hash });
});

export const $logs = createStore<Log[]>([])
  .on(getLogsFx.doneData, (_, logsResponse) => logsResponse.data)
  .on(queryLogsFx.doneData, (_, logsResponse) => logsResponse.data)
  .reset(reset);

export const $pages = createStore(0)
  .on(getLogsFx.doneData, (_, logsResponse) => logsResponse.meta.pages)
  .on(queryLogsFx.doneData, (_, logsResponse) => logsResponse.meta.pages);

export const $groupedLogs = createStore<LogsGroup>({
  tags: [],
  schema_name: "",
  group_hash: "",
  logs: [],
});

sample({
  source: getLogGroupFx.doneData,
  fn: (logsResponse) => {
    // Extract to separate reducer
    const logs = logsResponse.data;

    const intl = Intl.DateTimeFormat(getLocale(), DATEFORMAT_OPTIONS);
    const logsGroup: LogsGroup = {
      tags: Object.entries(logs[0]?.query),
      schema_name: logs[0]?.schema_name,
      group_hash: logs[0]?.group_hash,
      logs: [],
    };

    logsGroup.logs = logs.map((log) => {
      return { ...log, created_at: intl.format(new Date(log.created_at)) };
    });

    return logsGroup;
  },
  target: $groupedLogs,
});

export const $logsGroups = $logs.map((logs) => {
  const intl = Intl.DateTimeFormat(getLocale(), DATEFORMAT_OPTIONS);
  const groupedLogs = logs.reduce((acc: Record<string, LogsGroup>, log) => {
    const key = JSON.stringify(log.query);

    if (acc[key]?.logs?.length === 5) return acc;

    const logsGroup: LogsGroup = {
      tags: Object.entries(log.query),
      schema_name: log.schema_name,
      group_hash: log.group_hash,
      logs: [],
    };

    const formattedLog = { ...log, created_at: intl.format(new Date(log.created_at)) };

    acc[key] = acc[key] || logsGroup;
    acc[key].logs.push(formattedLog);

    return acc;
  }, {});

  return Object.values(groupedLogs);
});

sample({
  clock: reset,
  target: [resetFilter, resetSearch, resetPage],
});

export const effects = {
  getLogsFx,
  queryLogsFx,
  getLogGroupFx,
  reset,
};

export const events = {
  queryChanged,
  filterChanged,
  kindChanged,
  resetFilter,
  resetSearch,
  pageChanged,
};
