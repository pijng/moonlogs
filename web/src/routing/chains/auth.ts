import { homeRoute, loginRoute } from "@/routing/shared";
import { $isAuthorized, $token, unauthorizedTriggered } from "@/shared/auth";
import { sample } from "effector";

sample({
  clock: unauthorizedTriggered,
  target: loginRoute.open,
});

sample({
  source: $isAuthorized,
  clock: [loginRoute.opened, $token],
  filter: (isAuthorized) => isAuthorized,
  target: homeRoute.open,
});
