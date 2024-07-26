import { forbiddenRoute, homeRoute, loginRoute, notFoundRoute, registerAdminRoute } from "@/shared/routing";
import {
  $isAuthorized,
  createInitialAdmin,
  notAllowedTriggered,
  notFoundTriggered,
  unauthorizedTriggered,
} from "@/shared/session";
import { redirect } from "atomic-router";
import { sample } from "effector";

redirect({
  clock: unauthorizedTriggered,
  route: loginRoute,
});

redirect({
  clock: notAllowedTriggered,
  route: forbiddenRoute,
});

redirect({
  clock: notFoundTriggered,
  route: notFoundRoute,
});

redirect({
  clock: createInitialAdmin,
  route: registerAdminRoute,
});

sample({
  source: [loginRoute.$isOpened, $isAuthorized],
  filter: ([loginRouteOpened, isAuthorized]) => loginRouteOpened && isAuthorized,
  target: homeRoute.open,
});
