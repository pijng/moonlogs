import { userModel } from "@/entities/user";
import { UserRole } from "@/shared/api/users";
import { $isAuthorized, getSessionFx, obtainSession, tokenReceived, unauthorizedTriggered } from "@/shared/auth";
import { sidebarClosed } from "@/shared/ui";
import {
  RouteInstance,
  RouteParamsAndQuery,
  chainRoute,
  createHistoryRouter,
  createRoute,
  createRouterControls,
  redirect,
} from "atomic-router";
import { createLink, linkRouter } from "atomic-router-forest";
import { Store, createEvent, sample } from "effector";
import { createBrowserHistory } from "history";
import { condition } from "patronum";

export const Link = createLink();

export const loginRoute = createRoute();
export const registerAdminRoute = createRoute();
export const homeRoute = createRoute();
export const logsRoute = createRoute<{ schemaName: string | Store<string> }>();
export const showLogRoute = createRoute<{ schemaName: string; hash: string }>();
export const membersRoute = createRoute();
export const memberCreateRoute = createRoute();
export const memberEditRoute = createRoute<{ id: number }>();
export const schemaCreateRoute = createRoute();
export const schemaEditRoute = createRoute<{ id: number }>();
export const apiTokensRoute = createRoute();
export const apiTokenCreateRoute = createRoute();
export const apiTokenEditRoute = createRoute<{ id: number }>();
export const tagsRoute = createRoute();
export const tagCreateRoute = createRoute();
export const tagEditRoute = createRoute<{ id: number }>();

export const ROUTES = [
  { path: "/login", route: loginRoute },
  { path: "/register", route: registerAdminRoute },
  { path: "/", route: homeRoute },
  { path: "/schemas/create", route: schemaCreateRoute },
  { path: "/schemas/:id/edit", route: schemaEditRoute },
  { path: "/logs/:schemaName", route: logsRoute },
  { path: "/logs/:schemaName/:hash", route: showLogRoute },
  { path: "/members", route: membersRoute },
  { path: "/members/create", route: memberCreateRoute },
  { path: "/members/:id/edit", route: memberEditRoute },
  { path: "/api_tokens", route: apiTokensRoute },
  { path: "/api_tokens/create", route: apiTokenCreateRoute },
  { path: "/api_tokens/:id/edit", route: apiTokenEditRoute },
  { path: "/tags", route: tagsRoute },
  { path: "/tags/create", route: tagCreateRoute },
  { path: "/tags/:id/edit", route: tagEditRoute },
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
  target: obtainSession,
});

export const chainAuthorized = (route: RouteInstance<any>) => {
  const sessionCheckStarted = createEvent<RouteParamsAndQuery<any>>();

  const alreadyAuthorized = sample({
    clock: sessionCheckStarted,
    filter: $isAuthorized,
  });

  sample({
    source: $isAuthorized,
    clock: sessionCheckStarted,
    filter: (isAuthorized) => !isAuthorized,
    target: obtainSession,
  });

  sample({
    source: getSessionFx.doneData,
    filter: (sessionResponse) =>
      (!sessionResponse?.data?.token || !sessionResponse.success) && !sessionResponse.data.should_create_initial_admin,
    target: unauthorizedTriggered,
  });

  return chainRoute({
    route,
    beforeOpen: sessionCheckStarted,
    openOn: [alreadyAuthorized, tokenReceived],
  });
};

export const chainRole = (role: UserRole, route: RouteInstance<any>, fallback?: RouteInstance<any>) => {
  const checkRoleStarted = createEvent<any>();
  const succeed = createEvent();
  const failed = createEvent();

  condition({
    source: sample({
      clock: checkRoleStarted,
      source: userModel.$currentAccount,
      fn: (user) => user.role === role,
    }),
    if: Boolean,
    then: succeed,
    else: failed,
  });

  redirect({
    clock: failed,
    route: fallback || homeRoute,
  });

  return chainRoute({
    route,
    beforeOpen: checkRoleStarted,
    openOn: succeed,
    cancelOn: failed,
  });
};

sample({
  clock: router.$path,
  target: sidebarClosed,
});
