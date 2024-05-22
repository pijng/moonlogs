import { chainRoute } from "atomic-router";
import { actionCreateRoute, actionEditRoute, actionsRoute, chainAuthorized, chainRole, homeRoute } from "@/shared/routing";
import { actionModel } from "@/entities/action";
import { schemaModel } from "@/entities/schema";

chainRoute({
  route: chainAuthorized(homeRoute),
  beforeOpen: {
    effect: actionModel.effects.getActionsFx,
    mapParams: () => ({}),
  },
});

chainRoute({
  route: chainAuthorized(actionsRoute),
  beforeOpen: {
    effect: actionModel.effects.getActionsFx,
    mapParams: ({ params }) => params,
  },
});

chainRoute({
  route: chainRole("Admin", chainAuthorized(actionCreateRoute)),
  beforeOpen: {
    effect: schemaModel.effects.getSchemasFx,
    mapParams: () => ({}),
  },
});

chainRoute({
  route: chainRole("Admin", chainAuthorized(actionEditRoute)),
  beforeOpen: {
    effect: actionModel.effects.getActionFx,
    mapParams: ({ params }) => params.id,
  },
});

chainRoute({
  route: chainRole("Admin", chainAuthorized(actionEditRoute)),
  beforeOpen: {
    effect: schemaModel.effects.getSchemasFx,
    mapParams: () => ({}),
  },
});
