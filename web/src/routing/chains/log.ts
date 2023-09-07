import { logModel } from "@/entities/log";
import { schemaModel } from "@/entities/schema";
import { chainAuthorized, controls, logsRoute, showLogRoute } from "@/routing/shared";
import { chainRoute, querySync } from "atomic-router";
import { combine, sample } from "effector";
import { debounce } from "patronum";

sample({
  clock: logsRoute.closed,
  target: logModel.effects.reset,
});

chainRoute({
  route: chainAuthorized(logsRoute),
  beforeOpen: {
    effect: logModel.effects.queryLogsFx,
    mapParams: (route) => ({
      schema_name: route.params.schemaName,
      text: route.query.q,
      query: route.query.f,
      page: route.query.p,
    }),
  },
});

chainRoute({
  route: chainAuthorized(logsRoute),
  beforeOpen: {
    effect: schemaModel.effects.getSchemasFx,
    mapParams: () => ({}),
  },
});

chainRoute({
  route: chainAuthorized(showLogRoute),
  beforeOpen: {
    effect: schemaModel.effects.getSchemasFx,
    mapParams: () => ({}),
  },
});

chainRoute({
  route: chainAuthorized(showLogRoute),
  beforeOpen: {
    effect: logModel.effects.getLogGroupFx,
    mapParams: (route) => ({ schema_name: route.params.schemaName, hash: route.params.hash }),
  },
});

querySync({
  source: { q: logModel.$searchQuery, f: logModel.$formattedSearchFilter, p: logModel.$currentPage },
  clock: debounce({
    source: combine(logModel.$searchQuery, logModel.$formattedSearchFilter, logModel.$currentPage),
    timeout: 200,
  }),
  route: logsRoute,
  controls: controls,
});
