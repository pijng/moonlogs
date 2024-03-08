import { createHistoryRouter, createRouterControls } from "atomic-router";
import { createLink } from "atomic-router-forest";
import { createBrowserHistory } from "history";
import {
  apiTokenCreateRoute,
  apiTokenEditRoute,
  apiTokensRoute,
  forbiddenRoute,
  homeRoute,
  loginRoute,
  logsRoute,
  memberCreateRoute,
  memberEditRoute,
  membersRoute,
  notFoundRoute,
  profileRoute,
  registerAdminRoute,
  schemaCreateRoute,
  schemaEditRoute,
  showLogRoute,
  tagCreateRoute,
  tagEditRoute,
  tagsRoute,
} from "./routes";
export * from "./routes";
export * from "./middlewares";

export const Link = createLink();

const ROUTES = [
  { path: "/login", route: loginRoute },
  { path: "/register", route: registerAdminRoute },
  { path: "/forbidden", route: forbiddenRoute },
  { path: "/not_found", route: notFoundRoute },
  { path: "/profile", route: profileRoute },
  { path: "/", route: homeRoute },
  { path: "/schemas/create", route: schemaCreateRoute },
  { path: "/schemas/:id/edit", route: schemaEditRoute },
  { path: "/logs/:schemaName", route: logsRoute },
  { path: "/logs/:schemaName/:hash", route: showLogRoute },
  { path: "/members", route: membersRoute },
  { path: "/members/create", route: memberCreateRoute },
  { path: "/members/:id/edit", route: memberEditRoute },
  { path: "/api_tokens", route: apiTokensRoute },
  { path: "/api_tokens/create", route: apiTokenCreateRoute },
  { path: "/api_tokens/:id/edit", route: apiTokenEditRoute },
  { path: "/tags", route: tagsRoute },
  { path: "/tags/create", route: tagCreateRoute },
  { path: "/tags/:id/edit", route: tagEditRoute },
];

export const controls = createRouterControls();
export const history = createBrowserHistory();
export const router = createHistoryRouter({ routes: ROUTES, controls: controls });
