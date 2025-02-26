import { Action, Log } from "@/shared/api";
import { Store, combine } from "effector";

export const createLogGroupAction = (action: Store<Action>, logGroup: Store<Log[]>) => {
  const log = logGroup.map((lg) => lg[0]);

  const $evaluated: Store<boolean> = combine(action, log, (action, log) => {
    if (!log) {
      return false;
    }

    if (action.schema_ids.length > 0 && !action.schema_ids.includes(log.schema_id)) {
      return false;
    }

    const conditions = action.conditions ?? [];
    return conditions.every((condition) => {
      let attrValue: string | null;
      if (condition.attribute in log.query) {
        attrValue = log.query[condition.attribute];
      } else {
        attrValue = log[condition.attribute as "kind" | "level"];
      }

      switch (condition.operation) {
        case "EXISTS":
          return !!attrValue;
        case "EMPTY":
          return !attrValue;
      }

      const result: boolean = eval(`"${attrValue}" ${condition.operation} "${condition.value}"`);

      return result;
    });
  });

  const $values: Store<Record<string, any>> = log.map((l) => {
    if (!l) {
      return {};
    }

    const values = l.query;
    values["kind"] = l.kind;
    values["level"] = l.level;

    return values;
  });

  const $link = combine($values, action, (values, action) => {
    return action.pattern.replace(/{{(.*?)}}/g, (match, key) => values[key.trim()] || match);
  });

  return { $evaluated, $link };
};
