import { homeRoute, loginRoute } from "@/routing/shared";
import { getSessionFx, unauthorizedTriggered } from "@/shared/auth";
import { sample } from "effector";

sample({
  clock: unauthorizedTriggered,
  target: loginRoute.open,
});

sample({
  source: getSessionFx.doneData,
  filter: (sessionResponse) => !!sessionResponse.data.token && sessionResponse.success,
  target: homeRoute.open,
});
