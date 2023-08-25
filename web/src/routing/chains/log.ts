import { logModel } from "@/entities/log";
import { schemaModel } from "@/entities/schema";
import { controls, logsRoute, showLogRoute } from "@/routing/shared";
import { chainRoute, querySync } from "atomic-router";
import { debounce } from "patronum";

export const logsLoadedRoute = chainRoute({
  route: logsRoute,
  beforeOpen: {
    effect: logModel.effects.queryLogsFx,
    mapParams: (route) => ({ schema_name: route.params.schemaName, text: route.query.q }),
  },
});

export const schemasLoadedRoute = chainRoute({
  route: logsRoute,
  beforeOpen: {
    effect: schemaModel.effects.getSchemasFx,
    mapParams: () => ({}),
  },
});

export const searchLogsRoute = querySync({
  source: { q: logModel.$searchQuery },
  clock: debounce({ source: logModel.$searchQuery, timeout: 300 }),
  route: logsRoute,
  controls: controls,
});

export const showLogLoadedRoute = chainRoute({
  route: showLogRoute,
  beforeOpen: {
    effect: logModel.effects.getLogGroupFx,
    mapParams: (route) => ({ schema_name: route.params.schemaName, hash: route.params.hash }),
  },
});
