import { chainRoute } from "atomic-router";
import { chainAuthorized, membersRoute } from "@/routing/shared";
import { userModel } from "@/entities/user";

chainRoute({
  route: chainAuthorized(membersRoute),
  beforeOpen: {
    effect: userModel.effects.getUsersFx,
    mapParams: ({ params }) => params,
  },
});
