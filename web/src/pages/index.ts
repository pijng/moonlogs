import { Layout } from "@/shared/ui";
import { HomePage } from "./home";
import { ShowLogPage } from "./log";
import { LogsListPage } from "./logs-list";

export function Pages() {
  Layout(() => {
    HomePage();
    LogsListPage();
    ShowLogPage();
  });
}
