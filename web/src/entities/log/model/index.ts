import { Log, getLogs } from "@/shared/api";
import { getLogGroup } from "@/shared/api/logs";
import { DATEFORMAT_OPTIONS, getLocale } from "@/shared/lib";
import { createEffect, createEvent, createStore, sample } from "effector";

type LogsGroup = { tags: string[]; schema_name: string; group_hash: string; logs: Log[]; formattedLogs: string[][] };

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

const getLogsFx = createEffect((schema_name: string) => {
  return getLogs({ schema_name: schema_name });
});

const queryLogsFx = createEffect(({ schema_name, text, query }: { schema_name: string; text?: string; query?: string }) => {
  const objectQuery = JSON.parse(query || "{}");
  const formattedQuery = Object.entries(objectQuery).reduce((acc, [k, v]) => (v ? { ...acc, [k]: v } : acc), {});
  return getLogs({ schema_name: schema_name, text: text, query: formattedQuery });
});

const getLogGroupFx = createEffect(({ schema_name, hash }: { schema_name: string; hash: string }) => {
  return getLogGroup({ schema_name: schema_name, hash: hash });
});

export const $logs = createStore<Log[]>([])
  .on(getLogsFx.doneData, (_, logs) => logs)
  .on(queryLogsFx.doneData, (_, logs) => logs)
  .reset(reset);

export const $groupedLogs = createStore<LogsGroup>({
  tags: [],
  schema_name: "",
  group_hash: "",
  logs: [],
  formattedLogs: [],
});

sample({
  source: getLogGroupFx.doneData,
  fn: (newLogs) => {
    const intl = Intl.DateTimeFormat(getLocale(), DATEFORMAT_OPTIONS);
    const logsGroup: LogsGroup = {
      tags: Object.entries(newLogs[0]?.query).map((q) => `${q[0]}: ${q[1]}`),
      schema_name: newLogs[0]?.schema_name,
      group_hash: newLogs[0]?.group_hash,
      logs: [],
      formattedLogs: [],
    };

    logsGroup.logs = newLogs.map((log) => {
      return { ...log, created_at: intl.format(new Date(log.created_at)) };
    });

    logsGroup.formattedLogs = logsGroup.logs.map((l) => [l.created_at, l.text]);

    return logsGroup;
  },
  target: $groupedLogs,
});

export const $logsGroups = $logs.map((logs) => {
  const intl = Intl.DateTimeFormat(getLocale(), DATEFORMAT_OPTIONS);
  const groupedLogs = logs.reduce((acc: Record<string, LogsGroup>, log) => {
    const key = JSON.stringify(log.query);

    const logsGroup: LogsGroup = {
      tags: Object.entries(log.query).map((q) => `${q[0]}: ${q[1]}`),
      schema_name: log.schema_name,
      group_hash: log.group_hash,
      logs: [],
      formattedLogs: [],
    };

    const formattedLog = { ...log, created_at: intl.format(new Date(log.created_at)) };

    acc[key] = acc[key] || logsGroup;
    acc[key].logs.push(formattedLog);
    acc[key].formattedLogs.push([formattedLog.created_at, formattedLog.text]);

    return acc;
  }, {});

  return Object.values(groupedLogs);
});

sample({
  clock: reset,
  target: [resetFilter, resetSearch],
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
  resetFilter,
  resetSearch,
};
