import { chainRoute } from "atomic-router";
import { chainAuthorized, chainRole, memberCreateRoute, memberEditRoute, membersRoute } from "@/routing/shared";
import { userModel } from "@/entities/user";
import { tagModel } from "@/entities/tag";

chainRoute({
  route: chainRole("Admin", chainAuthorized(membersRoute)),
  beforeOpen: {
    effect: userModel.effects.getUsersFx,
    mapParams: ({ params }) => params,
  },
});

chainRoute({
  route: chainRole("Admin", chainAuthorized(memberCreateRoute)),
  beforeOpen: {
    effect: tagModel.effects.getTagsFx,
    mapParams: () => ({}),
  },
});

chainRoute({
  route: chainRole("Admin", chainAuthorized(memberEditRoute)),
  beforeOpen: {
    effect: userModel.effects.getUserFx,
    mapParams: ({ params }) => params.id,
  },
});

chainRoute({
  route: chainRole("Admin", chainAuthorized(memberEditRoute)),
  beforeOpen: {
    effect: tagModel.effects.getTagsFx,
    mapParams: () => ({}),
  },
});
