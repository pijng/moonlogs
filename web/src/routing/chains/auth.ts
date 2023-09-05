import { homeRoute, loginRoute } from "@/routing/shared";
import { $isAuthorized, unauthorizedTriggered } from "@/shared/auth";
import { sample } from "effector";

sample({
  clock: unauthorizedTriggered,
  target: loginRoute.open,
});

sample({
  source: $isAuthorized,
  clock: loginRoute.opened,
  filter: (isAuthorized) => Boolean(isAuthorized),
  target: homeRoute.open,
});
