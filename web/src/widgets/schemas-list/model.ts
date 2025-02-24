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

const $taggedSchemas = $filteredSchemas.map((schemas) => schemas.filter((s) => Boolean(s.tag_id)));
export const $generalSchemas = $filteredSchemas.map((schemas) => schemas.filter((s) => !Boolean(s.tag_id)));

export const $groupedTaggedSchemas = $taggedSchemas.map((schemas) => {
  return schemas.reduce((acc: Record<string, Schema[]>, schema) => {
    const tagId = schema.tag_id!;
    acc[tagId] = acc[tagId] || [];
    acc[tagId].push(schema);

    return acc;
  }, {});
});

export const $sortedTags = tagModel.$tags.map((tags) => {
  return tags.sort((t1, t2) => Number(t1.view_order) - Number(t2.view_order));
});
