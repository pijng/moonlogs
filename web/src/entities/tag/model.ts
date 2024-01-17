import { Tag, getTag, getTags } from "@/shared/api";
import { createEffect, createStore } from "effector";

const getTagsFx = createEffect(() => {
  return getTags();
});

const getTagFx = createEffect((id: number) => {
  return getTag(id);
});

export const $tags = createStore<Tag[]>([]).on(getTagsFx.doneData, (_, tagsResponse) => tagsResponse.data);

export const $currentTag = createStore<Tag>({ id: 0, name: "" }).on(getTagFx.doneData, (_, tagResponse) => tagResponse.data);

export const effects = {
  getTagsFx,
  getTagFx,
};

export const events = {};
