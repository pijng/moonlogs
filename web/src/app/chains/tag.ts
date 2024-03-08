import { chainRoute } from "atomic-router";
import { chainAuthorized, chainRole, homeRoute, tagCreateRoute, tagEditRoute, tagsRoute } from "@/shared/routing";
import { createEffect } from "effector";
import { tagModel } from "@/entities/tag";

chainRoute({
  route: chainAuthorized(homeRoute),
  beforeOpen: {
    effect: tagModel.effects.getTagsFx,
    mapParams: () => ({}),
  },
});

chainRoute({
  route: chainAuthorized(tagsRoute),
  beforeOpen: {
    effect: tagModel.effects.getTagsFx,
    mapParams: ({ params }) => params,
  },
});

chainRoute({
  route: chainRole("Admin", chainAuthorized(tagCreateRoute)),
  beforeOpen: {
    effect: createEffect(),
    mapParams: () => ({}),
  },
});

chainRoute({
  route: chainRole("Admin", chainAuthorized(tagEditRoute)),
  beforeOpen: {
    effect: tagModel.effects.getTagFx,
    mapParams: ({ params }) => params.id,
  },
});
