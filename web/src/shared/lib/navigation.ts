import { RouteInstance, RouteParams } from "atomic-router";
import { createLink } from "atomic-router-forest";
import { createEffect, createEvent, createStore, is, sample } from "effector";
import { condition } from "patronum";

export const bindLinkNavigation = <T extends RouteParams>({
  params,
  route,
  link,
}: {
  params?: T;
  route: RouteInstance<T>;
  link: ReturnType<typeof createLink>;
}) => {
  const $params = createStore<T>(params || ({} as T));
  const $route = createStore(route);
  const click = createEvent<MouseEvent>();
  const navigate = createEvent<any>();
  const openNewTab = createEvent();
  const openNewTabFx = createEffect((path: string) => {
    window.open(path, "_blank");
  });

  condition({
    source: sample({
      clock: click,
      fn: (evt) => evt.metaKey || evt.altKey || evt.ctrlKey || evt.shiftKey,
    }),
    if: Boolean,
    then: openNewTab,
    else: navigate,
  });

  sample({
    clock: navigate,
    source: $params,
    fn: (params) => ({ params: buildParams(params), query: {} }),
    target: route.navigate,
  });

  sample({
    clock: openNewTab,
    source: [$params, $route, link.$routes] as const,
    fn: ([params, route, routes]) => {
      const rawPath = routes.find((r) => r.route === route)?.path || "";
      const newParams = buildParams(params);
      const path = rawPath.replace(/:([\w]+)/g, (_, key) => newParams[key] ?? `:${key}`);

      return path;
    },
    target: openNewTabFx,
  });

  const buildParams = (params: T): T => {
    const newParams: RouteParams = {};
    for (const [key, value] of Object.entries(params)) {
      newParams[key] = is.store(value) ? value.getState() : value;
    }

    return newParams as T;
  };

  return click;
};
