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
};

export interface LogsResponse extends BaseResponse {
  data: Log[];
}

export const getLogs = ({
  schema_name,
  text,
  query,
  kind,
  page,
}: {
  schema_name: string;
  text?: string;
  query?: Record<string, string | number | boolean>;
  kind?: string;
  page?: number;
}): Promise<LogsResponse> => {
  return post({
    url: `/api/logs/search?page=${page}`,
    body: JSON.stringify({ schema_name: schema_name, text: text, query: query, kind: kind }),
  });
};

export const getLogGroup = ({ schema_name, hash }: { schema_name: string; hash: string }): Promise<LogsResponse> => {
  return get({
    url: `/api/logs/group/${schema_name}/${hash}`,
  });
};
