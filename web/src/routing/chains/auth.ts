import { homeRoute, loginRoute } from "@/routing/shared";
import { $isAuthorized, unauthorizedTriggered } from "@/shared/auth";
import { sample } from "effector";

sample({
  clock: unauthorizedTriggered,
  target: loginRoute.open,
});

sample({
  source: [loginRoute.$isOpened, $isAuthorized],
  filter: ([loginRouteOpened, isAuthorized]) => loginRouteOpened && isAuthorized,
  target: homeRoute.open,
});
