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

export const getLogs = ({
  schema_name,
  text,
  query,
}: {
  schema_name: string;
  text?: string;
  query?: Record<string, string | number | boolean>;
}): Promise<Log[]> => {
  return post({
    url: `/api/logs/search`,
    body: JSON.stringify({ schema_name: schema_name, text: text, query: query }),
  });
};

export const getLogGroup = ({ schema_name, hash }: { schema_name: string; hash: string }): Promise<Log[]> => {
  return get({
    url: `/api/logs/group/${schema_name}/${hash}`,
  });
};
