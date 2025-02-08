import "./chains";
import { userModel } from "@/entities/user";
import { notFoundTriggered, obtainSession } from "@/shared/session";
import { Link, history, router } from "@/shared/routing";
import { sidebarClosed } from "@/shared/ui";
import { linkRouter, onAppMount } from "atomic-router-forest";
import { createEvent, sample } from "effector";
import { h, spec } from "forest";
import { Pages } from "./pages";
import { loadClipboardFromStorageFx, loadThemeFromStorageFx } from "@/shared/lib";

export const Application = () => {
  h("body", () => {
    spec({
      classList: {
        "w-full": true,
        "min-h-screen": true,
        "bg-white": true,
        "text-slate-700": true,
        "dark:bg-eigengrau": true,
        "dark:text-slate-300": true,
      },
    });

    Pages();
    onAppMount(appMounted);
  });
};

// This event need to setup initial configuration. You can move it into src/shared
export const appMounted = createEvent();

sample({
  clock: router.routeNotFound,
  target: notFoundTriggered,
});

// Attach history for the router on the app start
sample({
  clock: appMounted,
  fn: () => history,
  target: router.setHistory,
});

// Add router into the Link instance to easily resolve routes paths
linkRouter({
  clock: appMounted,
  router,
  Link: Link,
});

sample({
  clock: appMounted,
  target: obtainSession,
});

sample({
  clock: appMounted,
  // TODO: move loadLocale to shared/lib/i18n
  target: [loadThemeFromStorageFx, loadClipboardFromStorageFx, userModel.effects.loadLocaleFromStorageFx],
});

sample({
  clock: router.$path,
  target: sidebarClosed,
});
