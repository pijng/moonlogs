import { UserRole } from "@/shared/api/users";
import { $currentAccount, $isAuthorized, getSessionFx, obtainSession, tokenReceived, unauthorizedTriggered } from "@/shared/auth";
import { RouteInstance, RouteParamsAndQuery, chainRoute, redirect } from "atomic-router";
import { createEvent, sample } from "effector";
import { condition } from "patronum";
import { homeRoute } from "./routes";

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
      source: $currentAccount,
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
