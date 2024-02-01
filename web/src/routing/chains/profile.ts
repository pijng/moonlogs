import { chainRoute } from "atomic-router";
import { chainAuthorized, profileRoute } from "@/routing/shared";
import { tagModel } from "@/entities/tag";
import { userModel } from "@/entities/user";

chainRoute({
  route: chainAuthorized(profileRoute),
  beforeOpen: {
    effect: tagModel.effects.getTagsFx,
    mapParams: () => ({}),
  },
});

chainRoute({
  route: chainAuthorized(profileRoute),
  beforeOpen: {
    effect: userModel.effects.loadThemeFromStorageFx,
    mapParams: () => ({}),
  },
});
