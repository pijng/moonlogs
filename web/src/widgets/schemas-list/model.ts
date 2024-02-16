import { schemaModel } from "@/entities/schema";
import { tagModel } from "@/entities/tag";
import { Schema } from "@/shared/api";
import { combine, createEvent, createStore } from "effector";

export const $searchQuery = createStore("");
export const queryChanged = createEvent<string>();
$searchQuery.on(queryChanged, (_, query) => query);

const $filteredSchemas = combine([schemaModel.$schemas, tagModel.$tags, $searchQuery], ([schemas, tags, searchQuery]) => {
  return schemas.filter((s) => {
    const titleMatched = s.title.toLowerCase().includes(searchQuery.toLocaleLowerCase());
    const descriptionMatched = s.description.toLocaleLowerCase().includes(searchQuery.toLowerCase());
    const tagMatched = tags
      .filter((t) => t.name.toLowerCase().includes(searchQuery.toLocaleLowerCase()))
      .map((t) => t.id)
      .includes(s.tag_id ?? 0);

    return titleMatched || descriptionMatched || tagMatched;
  });
});

export const $groupedFilteredSchemas = $filteredSchemas.map((schemas) => {
  return schemas.reduce((acc: Record<string, Schema[]>, schema) => {
    const tag = schema.tag_id || "General";

    acc[tag] = acc[tag] || [];
    acc[tag].push(schema);

    return acc;
  }, {});
});
