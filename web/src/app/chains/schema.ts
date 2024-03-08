import { chainRoute } from "atomic-router";
import { schemaModel } from "@/entities/schema";
import { chainAuthorized, chainRole, homeRoute, schemaCreateRoute, schemaEditRoute } from "@/shared/routing";
import { tagModel } from "@/entities/tag";

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
    effect: tagModel.effects.getTagsFx,
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

chainRoute({
  route: chainRole("Admin", chainAuthorized(schemaEditRoute)),
  beforeOpen: {
    effect: tagModel.effects.getTagsFx,
    mapParams: ({ params }) => params.id,
  },
});
