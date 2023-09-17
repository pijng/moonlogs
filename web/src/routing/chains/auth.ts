import { loginRoute } from "@/routing/shared";
import { unauthorizedTriggered } from "@/shared/auth";
import { sample } from "effector";

sample({
  clock: unauthorizedTriggered,
  target: loginRoute.open,
});
