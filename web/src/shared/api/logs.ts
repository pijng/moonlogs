import { TIMEZONE } from "@/shared/lib";
import { BaseResponse, get, post } from "./base";

export type Level = "Trace" | "Debug" | "Info" | "Warn" | "Error" | "Fatal";

export type Log = {
  id: string;
  text: string;
  created_at: string;
  level: Level;
  schema_name: string;
  group_hash: string;
  kind: string | null;
  schema_id: number;
  query: Record<string, any>;
  request: Record<string, any>;
  response: Record<string, any>;
  old_value: string;
  new_value: string;
};

export type LogsGroup = { tags: Array<[string, any]>; schema_name: string; kind: string | null; group_hash: string; logs: Log[] };

export interface LogsResponse extends BaseResponse {
  data: Log[];
}

export const getLogs = ({
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
  query?: Record<string, string | number | boolean>;
  kind?: string;
  level?: Level;
  from?: Date;
  to?: Date;
  page?: number;
}): Promise<LogsResponse> => {
  return post({
    url: `/api/logs/search?from=${from ?? ""}&to=${to ?? ""}&tz=${TIMEZONE}&page=${page ?? 1}`,
    body: JSON.stringify({ schema_name, text, query, kind, level }),
  });
};

export const getLogGroup = ({ schema_name, hash }: { schema_name: string; hash: string }): Promise<LogsResponse> => {
  return get({
    url: `/api/logs/group/${schema_name}/${hash}`,
  });
};
