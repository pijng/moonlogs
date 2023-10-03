import { $isAuthorized, getSessionFx, tokenReceived, unauthorizedTriggered } from "@/shared/auth";
import {
  RouteInstance,
  RouteParamsAndQuery,
  chainRoute,
  createHistoryRouter,
  createRoute,
  createRouterControls,
} from "atomic-router";
import { createLink, linkRouter } from "atomic-router-forest";
import { createEvent, sample } from "effector";
import { createBrowserHistory } from "history";

export const Link = createLink();

export const loginRoute = createRoute();
export const homeRoute = createRoute();
export const logsRoute = createRoute<{ schemaName: string }>();
export const showLogRoute = createRoute<{ schemaName: string; hash: string }>();
export const membersRoute = createRoute();
export const memberCreateRoute = createRoute();
export const memberEditRoute = createRoute<{ id: number }>();
export const schemaCreateRoute = createRoute();
export const schemaEditRoute = createRoute<{ id: number }>();

export const ROUTES = [
  { path: "/login", route: loginRoute },
  { path: "/", route: homeRoute },
  { path: "/schemas/create", route: schemaCreateRoute },
  { path: "/schemas/:id/edit", route: schemaEditRoute },
  { path: "/logs/:schemaName", route: logsRoute },
  { path: "/logs/:schemaName/:hash", route: showLogRoute },
  { path: "/members", route: membersRoute },
  { path: "/members/create", route: memberCreateRoute },
  { path: "/members/:id/edit", route: memberEditRoute },
];

const history = createBrowserHistory();
export const controls = createRouterControls();
export const router = createHistoryRouter({ routes: ROUTES, controls });

// This event need to setup initial configuration. You can move it into src/shared
export const appMounted = createEvent();

// Attach history for the router on the app start
sample({
  clock: appMounted,
  fn: () => history,
  target: router.setHistory,
});

// Add router into the Link instance to easily resolve routes paths
linkRouter({
  clock: appMounted,
  router,
  Link,
});

sample({
  clock: appMounted,
  target: getSessionFx,
});

export const chainAuthorized = (route: RouteInstance<any>) => {
  const sessionCheckStarted = createEvent<RouteParamsAndQuery<any>>();

  const alreadyAuthorized = sample({
    clock: sessionCheckStarted,
    filter: $isAuthorized,
  });

  const sessionCheck = sample({
    source: $isAuthorized,
    clock: sessionCheckStarted,
    filter: (isAuthorized) => !isAuthorized,
    target: getSessionFx,
  });

  sample({
    source: sessionCheck.doneData,
    filter: (sessionResponse) => !sessionResponse?.data?.token || !sessionResponse.success,
    target: unauthorizedTriggered,
  });

  return chainRoute({
    route,
    beforeOpen: sessionCheckStarted,
    openOn: [alreadyAuthorized, tokenReceived],
  });
};
