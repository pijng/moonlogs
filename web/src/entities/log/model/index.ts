import { Log, getLogs } from "@/shared/api";
import { getLogGroup } from "@/shared/api/logs";
import { DATEFORMAT_OPTIONS, getLocale } from "@/shared/lib";
import { createEffect, createEvent, createStore } from "effector";

type LogsGroup = { tags: string[]; schema_name: string; group_hash: string; logs: Log[] };

const getLogsFx = createEffect((schema_name: string) => {
  return getLogs({ schema_name: schema_name });
});

const queryLogsFx = createEffect(
  ({
    schema_name,
    text,
    query,
  }: {
    schema_name: string;
    text?: string;
    query?: Record<string, string | number | boolean>;
  }) => {
    return getLogs({ schema_name: schema_name, text: text, query: query });
  },
);

const getLogGroupFx = createEffect(({ schema_name, hash }: { schema_name: string; hash: string }) => {
  return getLogGroup({ schema_name: schema_name, hash: hash });
});

export const $logs = createStore<Log[]>([])
  .on(getLogsFx.doneData, (_, logs) => logs)
  .on(queryLogsFx.doneData, (_, logs) => logs);

export const $logsGroups = $logs.map((logs) => {
  const intl = Intl.DateTimeFormat(getLocale(), DATEFORMAT_OPTIONS);
  const groupedLogs = logs.reduce((acc: Record<string, LogsGroup>, log) => {
    const key = JSON.stringify(log.query);

    const logsGroup: LogsGroup = {
      tags: Object.entries(log.query).map((q) => `${q[0]}: ${q[1]}`),
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

export const $groupedLogs = createStore<Log[]>([]).on(getLogGroupFx.doneData, (_, logs) => logs);

export const $searchQuery = createStore("");
const queryChanged = createEvent<string>();
$searchQuery.on(queryChanged, (_, query) => query);

export const effects = {
  getLogsFx,
  queryLogsFx,
  getLogGroupFx,
};

export const events = {
  queryChanged,
};
