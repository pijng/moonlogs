import { Log, getLogs } from "@/shared/api";
import { Level, LogsGroup, getLogGroup } from "@/shared/api/logs";
import { DATEFORMAT_OPTIONS, objectToQueryString, queryStringToObject } from "@/shared/lib";
import { $preferredLanguage } from "@/shared/lib/i18n/locale";
import { combine, createEffect, createEvent, createStore, sample } from "effector";

const reset = createEffect();

const resetSearch = createEvent();
export const $searchQuery = createStore("").reset(resetSearch);

const queryChanged = createEvent<string>();
$searchQuery.on(queryChanged, (_, query) => query);

const resetFilter = createEvent();
const filterChanged = createEvent<Record<string, any>>();
export const $formattedSearchFilter = createStore<string>("").reset(resetFilter);

$formattedSearchFilter.on(filterChanged, (formattedFilter, changedFilter) => {
  const filter = queryStringToObject(formattedFilter);
  const newFilter = { ...filter, ...changedFilter };

  return objectToQueryString(newFilter);
});

const resetKindFilter = createEvent();
const kindChanged = createEvent<string>();
export const $currentKind = createStore("").reset(resetKindFilter, resetFilter);
$currentKind.on(kindChanged, (_, kind) => kind);

const resetLevelFilter = createEvent();
const levelChanged = createEvent<Level>();
export const $currentLevel = createStore<Level | string>("").reset(resetLevelFilter);
$currentLevel.on(levelChanged, (_, level) => level);

const resetTimeFilter = createEvent();
const fromTimeChanged = createEvent<any>();
export const $currentFromTime = createStore("").reset(resetTimeFilter);
$currentFromTime.on(fromTimeChanged, (_, from) => from);

const toTimeChanged = createEvent<any>();
export const $currentToTime = createStore("").reset(resetTimeFilter);
$currentToTime.on(toTimeChanged, (_, to) => to);

const resetPage = createEvent();
const pageChanged = createEvent<string>();

export const $currentPage = createStore("1")
  .on(pageChanged, (_, newPage) => newPage)
  .reset([queryChanged, filterChanged, kindChanged, fromTimeChanged, toTimeChanged, resetPage]);

const getLogsFx = createEffect((schema_name: string) => {
  return getLogs({ schema_name: schema_name });
});

const queryLogsFx = createEffect(
  ({
    schema_name,
    text,
    query,
    kind,
    level,
    from,
    to,
    page,
  }: {
    schema_name: string;
    text?: string;
    query?: string;
    kind?: string;
    level?: Level;
    from?: Date;
    to?: Date;
    page?: number;
  }) => {
    const objectQuery = queryStringToObject(query);

    const formattedQuery = Object.entries(objectQuery).reduce((acc, [k, v]) => {
      if (v.trim().length > 0) {
        return { ...acc, [k]: decodeURIComponent(v) };
      }

      return acc;
    }, {});

    return getLogs({
      schema_name: schema_name,
      text: text,
      query: formattedQuery,
      kind: kind,
      level: level,
      from: from,
      to: to,
      page: page,
    });
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
  kind: null,
  schema_name: "",
  group_hash: "",
  logs: [],
});

sample({
  source: $preferredLanguage,
  clock: getLogGroupFx.doneData,
  fn: (lang, logsResponse) => {
    // Extract to separate reducer
    const logs = logsResponse.data;

    const intl = Intl.DateTimeFormat(lang, DATEFORMAT_OPTIONS);
    const logsGroup: LogsGroup = {
      kind: logs[0].kind,
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

export const $logsGroups = combine([$logs, $preferredLanguage], ([logs, lang]) => {
  const intl = Intl.DateTimeFormat(lang, DATEFORMAT_OPTIONS);
  const groupedLogs = logs.reduce((acc: Record<string, LogsGroup>, log) => {
    const key = JSON.stringify(log.query);

    // if (acc[key]?.logs?.length === 5) return acc;

    const logsGroup: LogsGroup = {
      tags: Object.entries(log.query),
      kind: log.kind,
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

const requestURLClicked = createEvent<number>();
const openRequestURLFx = createEffect((id: number) => {
  window.open(`/api/logs/${id}/request`);
});

sample({
  source: requestURLClicked,
  target: openRequestURLFx,
});

const responseURLClicked = createEvent<number>();
const openResponseURLFx = createEffect((id: number) => {
  window.open(`/api/logs/${id}/response`);
});

sample({
  source: responseURLClicked,
  target: openResponseURLFx,
});

sample({
  clock: reset,
  target: [resetFilter, resetKindFilter, resetLevelFilter, resetSearch, resetPage],
});

export const effects = {
  getLogsFx,
  queryLogsFx,
  getLogGroupFx,
  reset,
};

export const events = {
  requestURLClicked,
  responseURLClicked,
  queryChanged,
  filterChanged,
  kindChanged,
  levelChanged,
  resetLevelFilter,
  resetFilter,
  fromTimeChanged,
  toTimeChanged,
  resetTimeFilter,
  resetSearch,
  pageChanged,
};
