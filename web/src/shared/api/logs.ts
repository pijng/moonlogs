import { get, post } from "./base";

export type Log = {
  id: string;
  text: string;
  created_at: string;
  schema_name: string;
  group_hash: string;
  schema_id: number;
  query: Record<string, any>;
};

export type LogsResponse = {
  data: Log[];
  meta: {
    page: number;
    count: number;
    pages: number;
  };
};

export const getLogs = ({
  schema_name,
  text,
  query,
  page,
}: {
  schema_name: string;
  text?: string;
  query?: Record<string, string | number | boolean>;
  page?: number;
}): Promise<LogsResponse> => {
  return post({
    url: `/api/logs/search?page=${page}`,
    body: JSON.stringify({ schema_name: schema_name, text: text, query: query }),
  });
};

export const getLogGroup = ({ schema_name, hash }: { schema_name: string; hash: string }): Promise<LogsResponse> => {
  return get({
    url: `/api/logs/group/${schema_name}/${hash}`,
  });
};
