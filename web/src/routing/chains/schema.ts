import { chainRoute } from "atomic-router";
import { schemaModel } from "@/entities/schema";
import { homeRoute } from "@/routing/shared";

chainRoute({
  route: homeRoute,
  beforeOpen: {
    effect: schemaModel.effects.getSchemasFx,
    mapParams: ({ params }) => params,
  },
});
