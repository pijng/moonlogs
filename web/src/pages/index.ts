import { combine } from "effector";
import { Layout } from "@/shared/ui";
import { HomePage } from "./home";
import { ShowLogPage } from "./log";
import { LogsListPage } from "./logs-list";
import { UsersListPage } from "./users-list";
import { LoginPage } from "./login";
import { loginRoute, registerAdminRoute } from "@/routing/shared";
import { UserCreatePage } from "./user-create";
import { UserEditPage } from "./user-edit";
import { SchemaCreatePage } from "./schema-create";
import { SchemaEditPage } from "./schema-edit";
import { ApiTokensListPage } from "./api-tokens-list";
import { ApiTokenCreatePage } from "./api-token-create";
import { ApiTokenEditPage } from "./api-token-edit";
import { RegisterAdminPage } from "./register-admin";
import { TagsListPage } from "./tags-list";
import { TagCreatePage } from "./tag-create";
import { TagEditPage } from "./tag-edit";

export function Pages() {
  const $layoutVisible = combine(
    [loginRoute.$isOpened, registerAdminRoute.$isOpened],
    ([loginOpened, registerOpened]) => !loginOpened && !registerOpened,
  );

  Layout({
    content: () => {
      LoginPage();
      RegisterAdminPage();
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
    },
    layoutVisible: $layoutVisible,
  });
}
