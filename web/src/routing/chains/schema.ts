import { chainRoute } from "atomic-router";
import { schemaModel } from "@/entities/schema";
import { homeRoute } from "@/routing/shared";

export const schemasLoadedRoute = chainRoute({
  route: homeRoute,
  beforeOpen: {
    effect: schemaModel.effects.getSchemasFx,
    mapParams: ({ params }) => params,
  },
});
