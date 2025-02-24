import { chainRoute } from "atomic-router";
import { chainAuthorized, insightsRoute } from "@/shared/routing";
import { schemaModel } from "@/entities/schema";

chainRoute({
  route: chainAuthorized(insightsRoute),
  beforeOpen: {
    effect: schemaModel.effects.getSchemasFx,
    mapParams: () => ({}),
  },
});
