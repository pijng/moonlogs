import { chainRoute } from "atomic-router";
import { schemaModel } from "@/entities/schema";
import { chainAuthorized, homeRoute } from "@/routing/shared";

chainRoute({
  route: chainAuthorized(homeRoute),
  beforeOpen: {
    effect: schemaModel.effects.getSchemasFx,
    mapParams: ({ params }) => params,
  },
});
