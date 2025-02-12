import { chainRoute } from "atomic-router";
import {
  alertingRuleCreateRoute,
  alertingRuleEditRoute,
  alertingRulesRoute,
  chainAuthorized,
  chainRole,
  homeRoute,
} from "@/shared/routing";
import { createEffect } from "effector";
import { alertingRuleModel } from "@/entities/alerting-rule";
import { schemaModel } from "@/entities/schema";

chainRoute({
  route: chainAuthorized(homeRoute),
  beforeOpen: {
    effect: alertingRuleModel.effects.getRulesFx,
    mapParams: () => ({}),
  },
});

chainRoute({
  route: chainAuthorized(alertingRulesRoute),
  beforeOpen: {
    effect: alertingRuleModel.effects.getRulesFx,
    mapParams: ({ params }) => params,
  },
});

chainRoute({
  route: chainRole("Admin", chainAuthorized(alertingRuleCreateRoute)),
  beforeOpen: {
    effect: createEffect(),
    mapParams: () => ({}),
  },
});

chainRoute({
  route: chainRole("Admin", chainAuthorized(alertingRuleCreateRoute)),
  beforeOpen: {
    effect: schemaModel.effects.getSchemasFx,
    mapParams: () => ({}),
  },
});

chainRoute({
  route: chainRole("Admin", chainAuthorized(alertingRuleEditRoute)),
  beforeOpen: {
    effect: alertingRuleModel.effects.getRuleFx,
    mapParams: ({ params }) => params.id,
  },
});

chainRoute({
  route: chainRole("Admin", chainAuthorized(alertingRuleEditRoute)),
  beforeOpen: {
    effect: schemaModel.effects.getSchemasFx,
    mapParams: () => ({}),
  },
});
