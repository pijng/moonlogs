import * as routes from '@/shared/routes';

import { HomePage } from './home';
import { LogsListPage } from './logs-list';

export const ROUTES = [
  { path: '/', route: routes.home },
  { path: '/logs', route: routes.logsList },
];

export function Pages() {
  HomePage();
  LogsListPage()
}