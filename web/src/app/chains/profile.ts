import { chainRoute } from "atomic-router";
import { chainAuthorized, profileRoute } from "@/shared/routing";
import { tagModel } from "@/entities/tag";

chainRoute({
  route: chainAuthorized(profileRoute),
  beforeOpen: {
    effect: tagModel.effects.getTagsFx,
    mapParams: () => ({}),
  },
});
