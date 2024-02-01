import { Link, apiTokensRoute, homeRoute, membersRoute, profileRoute, tagsRoute } from "@/routing/shared";
import { RouteInstance, redirect } from "atomic-router";
import { createEvent, createStore, restore, sample } from "effector";
import { h, spec } from "forest";
import { layoutClicked } from "..";
import { PermissionGate } from "@/shared/lib";

export const sidebarToggled = createEvent<any>();
export const sidebarClosed = createEvent<any>();
const $isOpened = createStore(false);

sample({
  source: $isOpened,
  clock: sidebarToggled,
  fn: (state) => !state,
  target: $isOpened,
});

sample({
  source: $isOpened,
  clock: sidebarClosed,
  filter: (isOpened) => isOpened,
  fn: (state) => !state,
  target: $isOpened,
});

sample({
  source: [$isOpened, restore(sidebarToggled, null)],
  clock: layoutClicked,
  filter: ([isOpened, clicked], layoutClicked) => {
    const path = layoutClicked.composedPath();

    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore
    return !path.includes(clicked?.currentTarget) && !layoutClicked.target.closest("aside") && isOpened;
  },
  target: sidebarToggled,
});

export const SidebarButton = () => {
  h("button", () => {
    spec({
      handler: { on: { click: sidebarToggled } },
      attr: {
        "data-drawer-target": "default-sidebar",
        "data-drawer-toggle": "default-sidebar",
        "aria-controls": "default-sidebar",
        type: "button",
      },
      classList: [
        "inline-flex",
        "items-center",
        "p-2",
        "mt-2",
        "ml-3",
        "text-sm",
        "text-gray-500",
        "rounded-lg",
        "sm:hidden",
        "hover:bg-gray-100",
        "focus:outline-none",
        "focus:ring-2",
        "focus:ring-gray-200",
        "dark:text-gray-400",
        "dark:hover:bg-gray-200",
        "dark:focus:ring-gray-600",
      ],
    });

    h("span", {
      classList: ["sr-only"],
      text: "Open sidebar",
    });

    h("svg", () => {
      spec({
        classList: ["w-6 h-6"],
        attr: { aria_hidden: true, fill: "currentColor", viewBox: "0 0 20 20", xmlns: "http://www.w3.org/2000/svg" },
      });

      h("path", {
        attr: {
          "clip-rule": "evenodd",
          "fill-rule": "evenodd",
          d: "M2 4.75A.75.75 0 012.75 4h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 4.75zm0 10.5a.75.75 0 01.75-.75h7.5a.75.75 0 010 1.5h-7.5a.75.75 0 01-.75-.75zM2 10a.75.75 0 01.75-.75h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 10z",
        },
      });
    });
  });
};

export const Sidebar = () => {
  h("aside", () => {
    spec({
      attr: { id: "default-sidebar", "aria-label": "Sidebar" },
      classList: {
        fixed: true,
        "top-0": true,
        "left-0": true,
        "z-40": true,
        "w-64": true,
        "h-screen": true,
        "transition-transform": true,
        "-translate-x-full": $isOpened.map((s) => !s),
        "sm:translate-x-0": $isOpened.map((s) => !s),
      },
    });

    h("div", () => {
      spec({
        classList: ["h-full", "px-3", "py-4", "overflow-y-auto", "bg-gray-50", "dark:bg-gray-800"],
      });

      h("a", () => {
        const homeClicked = createEvent<MouseEvent>();

        redirect({
          clock: homeClicked,
          route: homeRoute,
        });

        spec({
          classList: ["flex", "items-center", "pl-2.5", "mb-5"],
          attr: { href: "" },
          handler: {
            on: {
              click: homeClicked,
            },
            config: { prevent: true },
          },
        });

        h("span", {
          classList: ["mr-3", "leading-7", "text-2xl"],
          attr: { alt: "Moonlogs logo" },
          text: "🌘",
        });

        h("span", {
          classList: ["self-center", "text-xl", "font-semibold", "whitespace-nowrap", "dark:text-white"],
          text: "Moonlogs",
        });
      });

      h("ul", () => {
        spec({
          classList: ["space-y-2", "font-medium"],
        });

        SidebarItem("Profile", profileRoute);

        SidebarItem("Log groups", homeRoute);

        PermissionGate("Admin", () => {
          SidebarItem("Members", membersRoute);
        });

        PermissionGate("Admin", () => {
          SidebarItem("Tags", tagsRoute);
        });

        PermissionGate("Admin", () => {
          SidebarItem("API tokens", apiTokensRoute);
        });
      });
    });
  });
};

export const SidebarItem = (text: string, route: RouteInstance<Record<string, any>>) => {
  h("li", () => {
    Link(route, {
      text: text,
      classList: {
        flex: true,
        "items-center": true,
        "p-2": true,
        "text-gray-900": true,
        "rounded-lg": true,
        "dark:text-white": true,
        "hover:bg-gray-200": true,
        "bg-gray-200": route.$isOpened,
        "dark:bg-gray-700": route.$isOpened,
        "dark:hover:bg-gray-700": true,
        group: true,
      },
    });
  });
};
