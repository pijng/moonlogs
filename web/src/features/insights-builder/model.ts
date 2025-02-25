import { schemaModel } from "@/entities/schema";
import { userModel } from "@/entities/user";
import { GeminiRequest, generateContent, getLogs, Log } from "@/shared/api";
import { $intl, $preferredLanguage, isObjectPresent } from "@/shared/lib";
import { attach, createEffect, createEvent, createStore, sample } from "effector";

type InsightFilter = {
  fieldName: string;
  fieldValue: string;
};

export const $filterList = createStore<InsightFilter[]>([{ fieldName: "", fieldValue: "" }]);

const buildInsight = createEvent();
const addFilter = createEvent();
const removeFilter = createEvent<number>();
const fieldNameChanged = createEvent<{ fieldName: string; idx: number }>();
const fieldValueChanged = createEvent<{ fieldValue: string; idx: number }>();

export const events = {
  buildInsight,
  addFilter,
  removeFilter,
  fieldNameChanged,
  fieldValueChanged,
};

sample({
  source: $filterList,
  clock: addFilter,
  fn: (filters) => {
    return [...filters, { fieldName: "", fieldValue: "" }];
  },
  target: $filterList,
});

sample({
  source: $filterList,
  clock: removeFilter,
  fn: (filters, idx) => [...filters.slice(0, idx), ...filters.slice(idx + 1)],
  target: $filterList,
});

sample({
  source: $filterList,
  clock: fieldNameChanged,
  fn: (filters, { fieldName, idx }) => {
    return filters.map((filter, index) => (index === idx ? { ...filter, fieldName: fieldName } : filter));
  },
  target: $filterList,
});

sample({
  source: $filterList,
  clock: fieldValueChanged,
  fn: (filters, { fieldValue, idx }) => {
    return filters.map((filter, index) => (index === idx ? { ...filter, fieldValue: fieldValue } : filter));
  },
  target: $filterList,
});

export const $fieldsOptions = schemaModel.$schemas.map((schemas) => {
  const allFields = schemas.flatMap((schema) => {
    const fields = schema.fields || [];
    return fields.map((f) => ({ title: f.title, name: f.name, id: f.name }));
  });

  return Array.from(new Map(allFields.map((f) => [f.id, f])).values());
});

const getLogsFx = attach({
  source: $filterList,
  effect: createEffect((filterList: InsightFilter[]) => {
    const query = filterList.reduce(
      (acc, filter) => {
        acc[filter.fieldName] = filter.fieldValue;
        return acc;
      },
      {} as Record<string, string>,
    );

    return getLogs({ schema_name: "", query, limit: 1000 });
  }),
});

sample({
  clock: buildInsight,
  target: getLogsFx,
});

export const $isLoadingLogs = getLogsFx.pending;

export const $insightLogs = createStore<Log[]>([]);

sample({
  clock: getLogsFx.pending,
  fn: () => [],
  target: $insightLogs,
});

sample({
  source: $intl,
  clock: getLogsFx.doneData,
  fn: (intl, logsResponse) => {
    const logs = logsResponse.data.reverse();
    const formattedLogs = logs.map((log) => {
      return { ...log, created_at: intl.format(new Date(log.created_at)) };
    });

    return formattedLogs;
  },
  target: $insightLogs,
});

// TODO: Oof, this one is huge, refactor.
const summarizeLogsAiFx = createEffect(({ token, lang, logs }: { token: string; lang: string; logs: Log[] }) => {
  const logsText = logs
    .map((log) => {
      let base = `[${log.created_at}] SCHEMA: "${log.schema_name}" LEVEL: "${log.level}" MSG: "${log.text}"`;
      if (isObjectPresent(log.changes)) {
        const changesText = Object.entries(log.changes)
          .map(([field, { old_value, new_value }]) => {
            return `Field "${field}" was changed from "${old_value}" to "${new_value}"`;
          })
          .join(". ");

        base += ` CHANGES_HISTORY: ${changesText}`;
      }
      return base;
    })
    .join("\n");

  const content: GeminiRequest = {
    contents: [
      {
        role: "user",
        parts: [
          {
            text: `Write short summary of these logs in '${lang}' language, try to highlight problems if there are any error level logs, use only new lines as markup:\n\n${logsText}`,
          },
        ],
      },
    ],
  };

  return generateContent({ token, content });
});

sample({
  source: { lang: $preferredLanguage, user: userModel.$currentAccount },
  clock: $insightLogs,
  filter: ({ user }, logs) => logs.length > 0 && !!user.gemini_token,
  fn: ({ lang, user }, logs) => {
    return {
      token: user.gemini_token!,
      lang,
      logs,
    };
  },
  target: summarizeLogsAiFx,
});

export const $aiSummary = createStore<string>("");
sample({
  source: summarizeLogsAiFx.doneData,
  filter: (resp) => !!resp && resp.candidates.length > 0,
  fn: (resp) => {
    return resp!.candidates[0].content.parts[0].text;
  },
  target: $aiSummary,
});

export type InsightSchema = {
  schemaTitle: string;
  schemaName: string;
  schemaId: number;
  schemaColor: string;
};

export const $insightsSchemas = createStore<InsightSchema[]>([]);

const COLORS: Record<number, string> = {
  0: "green-500",
  1: "blue-500",
  2: "gray-500",
  3: "yellow-500",
  4: "red-500",
  5: "pink-500",
  6: "purple-500",
  7: "lime-500",
  8: "cyan-500",
  9: "orange-500",
};

sample({
  source: { logs: $insightLogs, schemas: schemaModel.$schemas },
  fn: ({ logs, schemas }) => {
    const uniqueSchemaIDs = new Set(logs.map((l) => l.schema_id));
    const relevantSchemas = schemas.filter((s) => uniqueSchemaIDs.has(s.id));

    return relevantSchemas.map((s) => {
      return {
        schemaTitle: s.title,
        schemaName: s.name,
        schemaId: s.id,
        schemaColor: COLORS[Number(s.id.toString()[0])],
      };
    });
  },
  target: $insightsSchemas,
});
