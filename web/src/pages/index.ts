import { Layout } from "@/shared/ui";
import { HomePage } from "./home";
import { ShowLogPage } from "./log";
import { LogsListPage } from "./logs-list";
import { UsersListPage } from "./users-list";
import { LoginPage } from "./login";
import { loginRoute } from "@/routing/shared";
import { UserCreatePage } from "./user-create";
import { UserEditPage } from "./user-edit";

export function Pages() {
  const $layoutVisible = loginRoute.$isOpened.map((state) => !state);

  Layout({
    content: () => {
      LoginPage();
      HomePage();
      LogsListPage();
      ShowLogPage();
      UsersListPage();
      UserCreatePage();
      UserEditPage();
    },
    layoutVisible: $layoutVisible,
  });
}
