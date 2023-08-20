import * as routes from "@/shared/routes";

import { HomePage } from "./home";
import { LogsListPage } from "./logs-list";
import { Layout } from "@/shared/ui/layout";

export const ROUTES = [
  { path: "/", route: routes.home },
  { path: "/logs", route: routes.logsList },
];

export function Pages() {
  Layout(() => {
    HomePage();
    LogsListPage();
  });
}
