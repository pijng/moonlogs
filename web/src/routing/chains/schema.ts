import { chainRoute } from "atomic-router";
import { schemaModel } from "@/entities/schema";
import { chainAuthorized, chainRole, homeRoute, schemaCreateRoute, schemaEditRoute } from "@/routing/shared";
import { createEffect } from "effector";

chainRoute({
  route: chainAuthorized(homeRoute),
  beforeOpen: {
    effect: schemaModel.effects.getSchemasFx,
    mapParams: () => ({}),
  },
});

chainRoute({
  route: chainRole("Admin", chainAuthorized(schemaCreateRoute)),
  beforeOpen: {
    effect: createEffect(),
    mapParams: () => ({}),
  },
});

chainRoute({
  route: chainRole("Admin", chainAuthorized(schemaEditRoute)),
  beforeOpen: {
    effect: schemaModel.effects.getSchemaFx,
    mapParams: ({ params }) => params.id,
  },
});
