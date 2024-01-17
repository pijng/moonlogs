import { apiTokenModel } from "@/entities/api-token";
import { apiTokenEditRoute } from "@/routing/shared";
import { ApiTokensTable } from "@/shared/ui";
import { createEvent, sample } from "effector";
import { h, spec } from "forest";

export const ApiTokensList = () => {
  h("div", () => {
    spec({
      classList: ["flex", "flex-col", "space-y-6", "mt-3"],
    });

    const editApiTokenClicked = createEvent<number>();
    sample({
      source: editApiTokenClicked,
      fn: (id) => ({ id }),
      target: apiTokenEditRoute.open,
    });

    ApiTokensTable(apiTokenModel.$apiTokens, editApiTokenClicked);
  });
};
