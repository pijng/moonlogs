import { Event, sample } from "effector";
import { Field } from "effector-forms";

export const bindFieldList = <T>({ field, added, removed }: { field: Field<T[]>; added: Event<T>; removed: Event<T> }) => {
  sample({
    source: field.$value,
    clock: added,
    fn: (fields, field) => [...fields, field],
    target: field.onChange,
  });

  sample({
    source: field.$value,
    clock: removed,
    fn: (fields, field) => fields.filter((f) => f !== field),
    target: field.onChange,
  });
};
