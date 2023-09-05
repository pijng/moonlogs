import { Layout } from "@/shared/ui";
import { HomePage } from "./home";
import { ShowLogPage } from "./log";
import { LogsListPage } from "./logs-list";
import { MembersListPage } from "./members-list";
import { LoginPage } from "./login";
import { loginRoute } from "@/routing";

export function Pages() {
  const $layoutVisible = loginRoute.$isOpened.map((state) => !state);

  Layout({
    content: () => {
      LoginPage();
      HomePage();
      LogsListPage();
      ShowLogPage();
      MembersListPage();
    },
    layoutVisible: $layoutVisible,
  });
}
