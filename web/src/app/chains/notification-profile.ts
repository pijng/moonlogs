import { chainRoute } from "atomic-router";
import {
  chainAuthorized,
  chainRole,
  notificationProfileCreateRoute,
  notificationProfileEditRoute,
  notificationProfileRoute,
} from "@/shared/routing";
import { notificationProfileModel } from "@/entities/notification-profile";
import { alertingRuleModel } from "@/entities/alerting-rule";

chainRoute({
  route: chainAuthorized(notificationProfileRoute),
  beforeOpen: {
    effect: notificationProfileModel.effects.getNotificationProfilesFx,
    mapParams: ({ params }) => params,
  },
});

chainRoute({
  route: chainRole("Admin", chainAuthorized(notificationProfileCreateRoute)),
  beforeOpen: {
    effect: notificationProfileModel.effects.getNotificationProfilesFx,
    mapParams: () => ({}),
  },
});

chainRoute({
  route: chainRole("Admin", chainAuthorized(notificationProfileCreateRoute)),
  beforeOpen: {
    effect: alertingRuleModel.effects.getRulesFx,
    mapParams: () => ({}),
  },
});

chainRoute({
  route: chainRole("Admin", chainAuthorized(notificationProfileEditRoute)),
  beforeOpen: {
    effect: notificationProfileModel.effects.getNotificationProfileFx,
    mapParams: ({ params }) => params.id,
  },
});

chainRoute({
  route: chainRole("Admin", chainAuthorized(notificationProfileEditRoute)),
  beforeOpen: {
    effect: alertingRuleModel.effects.getRulesFx,
    mapParams: () => ({}),
  },
});
