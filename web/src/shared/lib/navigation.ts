import { RouteInstance, RouteParams } from "atomic-router";
import { createEffect, createEvent, createStore, Event, is, sample } from "effector";
import { DOMElement } from "forest";
import { condition } from "patronum";

export const bindLinkNavigation = <T extends RouteParams>({
  params,
  route,
}: {
  params?: T;
  route: RouteInstance<T>;
}): { click: Event<MouseEvent>; mounted: Event<DOMElement> } => {
  const $params = createStore<T>(params || ({} as T));
  const mounted = createEvent<DOMElement>();
  const click = createEvent<MouseEvent>();
  const navigate = createEvent<any>();
  const openNewTab = createEvent();

  const $href = createStore<string>("");
  sample({
    clock: mounted,
    fn: (el) => {
      if (el instanceof HTMLAnchorElement) {
        return el.href;
      }
      return "";
    },
    target: $href,
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

  const openNewTabFx = createEffect((path: string) => {
    window.open(path, "_blank");
  });

  sample({
    clock: openNewTab,
    source: $href,
    target: openNewTabFx,
  });

  const buildParams = (params: T): T => {
    const newParams: RouteParams = {};
    for (const [key, value] of Object.entries(params)) {
      newParams[key] = is.store(value) ? value.getState() : value;
    }

    return newParams as T;
  };

  return { click, mounted };
};
