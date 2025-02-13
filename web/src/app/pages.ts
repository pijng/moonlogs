import { ActionCreatePage } from "@/pages/action-create";
import { ActionEditPage } from "@/pages/action-edit";
import { ActionsListPage } from "@/pages/actions-list";
import { AlertingRuleCreatePage } from "@/pages/alerting-rule-create";
import { AlertingRuleEditPage } from "@/pages/alerting-rule-edit";
import { AlertingRulesListPage } from "@/pages/alerting-rule-list";
import { ApiTokenCreatePage } from "@/pages/api-token-create";
import { ApiTokenEditPage } from "@/pages/api-token-edit";
import { ApiTokensListPage } from "@/pages/api-tokens-list";
import { ForbiddenPage } from "@/pages/forbidden";
import { HomePage } from "@/pages/home";
import { ShowLogPage } from "@/pages/log";
import { LoginPage } from "@/pages/login";
import { LogsListPage } from "@/pages/logs-list";
import { NotFoundPage } from "@/pages/not-found";
import { ProfilePage } from "@/pages/profile";
import { RegisterAdminPage } from "@/pages/register-admin";
import { SchemaCreatePage } from "@/pages/schema-create";
import { SchemaEditPage } from "@/pages/schema-edit";
import { TagCreatePage } from "@/pages/tag-create";
import { TagEditPage } from "@/pages/tag-edit";
import { TagsListPage } from "@/pages/tags-list";
import { UserCreatePage } from "@/pages/user-create";
import { UserEditPage } from "@/pages/user-edit";
import { UsersListPage } from "@/pages/users-list";
import { forbiddenRoute, loginRoute, notFoundRoute, registerAdminRoute } from "@/shared/routing";
import { Layout } from "@/shared/ui";
import { combine } from "effector";
import { NotificationProfileCreatePage } from "@/pages/notification-profile-create";
import { NotificationProfleEditPage } from "@/pages/notification-profile-edit";
import { NotificationProfileListPage } from "@/pages/notification-profile-list";

export const Pages = () => {
  const $layoutVisible = combine(
    [loginRoute.$isOpened, registerAdminRoute.$isOpened, forbiddenRoute.$isOpened, notFoundRoute.$isOpened],
    ([loginOpened, registerOpened, forbiddenOpened, notFoundOpened]) => {
      return !loginOpened && !registerOpened && !forbiddenOpened && !notFoundOpened;
    },
  );

  Layout({
    content: () => {
      LoginPage();
      RegisterAdminPage();
      ForbiddenPage();
      NotFoundPage();
      ProfilePage();
      HomePage();
      LogsListPage();
      ShowLogPage();
      UsersListPage();
      UserCreatePage();
      UserEditPage();
      SchemaCreatePage();
      SchemaEditPage();
      ApiTokensListPage();
      ApiTokenCreatePage();
      ApiTokenEditPage();
      TagsListPage();
      TagCreatePage();
      TagEditPage();
      ActionsListPage();
      ActionCreatePage();
      ActionEditPage();
      AlertingRulesListPage();
      AlertingRuleCreatePage();
      AlertingRuleEditPage();
      NotificationProfileListPage();
      NotificationProfileCreatePage();
      NotificationProfleEditPage();
    },
    layoutVisible: $layoutVisible,
  });
};
