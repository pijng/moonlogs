import { tagModel } from "@/entities/tag";
import { tagEditRoute } from "@/routing/shared";
import { TagsTable } from "@/shared/ui";
import { createEvent, sample } from "effector";
import { h, spec } from "forest";

export const TagsList = () => {
  h("div", () => {
    spec({
      classList: ["flex", "flex-col", "space-y-6", "mt-3"],
    });

    const editTagClicked = createEvent<number>();
    sample({
      source: editTagClicked,
      fn: (id) => ({ id }),
      target: tagEditRoute.open,
    });

    TagsTable(tagModel.$tags, editTagClicked);
  });
};
