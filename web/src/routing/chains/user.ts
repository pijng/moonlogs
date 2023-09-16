import { chainRoute } from "atomic-router";
import { chainAuthorized, memberCreateRoute, memberEditRoute, membersRoute } from "@/routing/shared";
import { userModel } from "@/entities/user";
import { createEffect } from "effector";

chainRoute({
  route: chainAuthorized(membersRoute),
  beforeOpen: {
    effect: userModel.effects.getUsersFx,
    mapParams: ({ params }) => params,
  },
});

chainRoute({
  route: chainAuthorized(memberCreateRoute),
  beforeOpen: {
    effect: createEffect(),
    mapParams: () => ({}),
  },
});

chainRoute({
  route: chainAuthorized(memberEditRoute),
  beforeOpen: {
    effect: userModel.effects.getUserFx,
    mapParams: ({ params }) => params.id,
  },
});
