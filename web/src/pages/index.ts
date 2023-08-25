import { Layout } from "@/shared/ui";
import { HomePage } from "./home";
import { LogsListPage } from "./logs-list";
import { ShowLogPage } from "./log";

export function Pages() {
  Layout(() => {
    HomePage();
    LogsListPage();
    ShowLogPage();
  });
}
