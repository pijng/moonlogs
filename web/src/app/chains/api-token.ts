import { chainRoute } from "atomic-router";

import { createEffect } from "effector";
import { apiTokenModel } from "@/entities/api-token";
import { apiTokenCreateRoute, apiTokenEditRoute, apiTokensRoute, chainAuthorized, chainRole } from "@/shared/routing";

chainRoute({
  route: chainRole("Admin", chainAuthorized(apiTokensRoute)),
  beforeOpen: {
    effect: apiTokenModel.effects.getApiTokensFx,
    mapParams: ({ params }) => params,
  },
});

chainRoute({
  route: chainRole("Admin", chainAuthorized(apiTokenCreateRoute)),
  beforeOpen: {
    effect: createEffect(),
    mapParams: () => ({}),
  },
});

chainRoute({
  route: chainRole("Admin", chainAuthorized(apiTokenEditRoute)),
  beforeOpen: {
    effect: apiTokenModel.effects.getApiTokenFx,
    mapParams: ({ params }) => params.id,
  },
});
