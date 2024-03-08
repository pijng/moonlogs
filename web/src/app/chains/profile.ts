import { chainRoute } from "atomic-router";
import { chainAuthorized, profileRoute } from "@/shared/routing";
import { tagModel } from "@/entities/tag";
import { userModel } from "@/entities/user";
import { createEffect } from "effector";

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
    effect: createEffect(() => {
      userModel.effects.loadThemeFromStorageFx();
      userModel.effects.loadLocaleFromStorageFx();
    }),
    mapParams: () => ({}),
  },
});
