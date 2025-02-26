import { RouteInstance, redirect } from "atomic-router";
import { Store, createEvent, createStore, sample } from "effector";
import { DOMElement, h, node, spec } from "forest";
import {
  GridPlusIcon,
  LayersIcon,
  LockOpenIcon,
  LogoIcon,
  PermissionGate,
  TagIcon,
  UserIcon,
  UsersIcon,
  BullhornIcon,
  ExclamationMarkIcon,
  LightBulbIcon,
} from "@/shared/ui";
import { i18n } from "@/shared/lib/i18n";
import {
  Link,
  actionsRoute,
  alertingRulesRoute,
  notificationProfileRoute,
  apiTokensRoute,
  homeRoute,
  membersRoute,
  profileRoute,
  tagsRoute,
  insightsRoute,
} from "@/shared/routing";
import { bindLinkNavigation } from "@/shared/lib";

export const sidebarClosed = createEvent();
const sidebarTriggered = createEvent<MouseEvent>();
const $sidebarVisible = createStore(false);
const outsideClicked = createEvent<{ node: DOMElement; event: any }>();

sample({
  clock: sidebarClosed,
  fn: () => false,
  target: $sidebarVisible,
});

sample({
  source: $sidebarVisible,
  clock: sidebarTriggered,
  fn: (v) => !v,
  target: $sidebarVisible,
});

sample({
  source: $sidebarVisible,
  clock: outsideClicked,
  filter: (visible, { node, event }) => !node.contains(event.target),
  fn: () => false,
  target: $sidebarVisible,
});

const SidebarButton = () => {
  h("button", () => {
    spec({
      handler: { on: { click: sidebarTriggered } },
      attr: {
        "data-drawer-target": "default-sidebar",
        "data-drawer-toggle": "default-sidebar",
        "aria-controls": "default-sidebar",
        type: "button",
      },
      classList: [
        "inline-flex",
        "items-center",
        "pt-3.5",
        "pl-3.5",
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
      text: i18n("components.sidebar.open"),
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
  h("div", () => {
    SidebarButton();

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
          "-translate-x-full": $sidebarVisible.map((v) => !v),
          "sm:translate-x-0": $sidebarVisible.map((v) => !v),
        },
      });

      h("div", () => {
        spec({
          classList: ["h-full", "px-3", "py-4", "overflow-y-auto", "bg-gray-50", "dark:bg-raisin-black"],
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

          h("div", () => {
            spec({
              classList: ["mr-2", "leading-7", "text-2xl", "w-8"],
              attr: { alt: "Moonlogs logo" },
            });

            LogoIcon();
          });

          h("span", {
            classList: ["self-center", "text-xl", "font-semibold", "whitespace-nowrap", "dark:text-white"],
            text: i18n("miscellaneous.brand"),
          });
        });

        h("ul", () => {
          spec({
            classList: ["space-y-2", "font-medium"],
          });

          SidebarItem({ text: i18n("profile.label"), route: profileRoute, icon: UserIcon });

          SidebarItem({ text: i18n("log_groups.label"), route: homeRoute, icon: LayersIcon });

          SidebarItem({ text: i18n("insights.label"), route: insightsRoute, icon: LightBulbIcon });

          PermissionGate("Admin", () => {
            SidebarItem({ text: i18n("members.label"), route: membersRoute, icon: UsersIcon });
          });

          PermissionGate("Admin", () => {
            SidebarItem({ text: i18n("tags.label"), route: tagsRoute, icon: TagIcon });
          });

          PermissionGate("Admin", () => {
            SidebarItem({ text: i18n("actions.label"), route: actionsRoute, icon: GridPlusIcon });
          });

          PermissionGate("Admin", () => {
            SidebarItem({ text: i18n("alerting_rules.label"), route: alertingRulesRoute, icon: ExclamationMarkIcon });
          });

          PermissionGate("Admin", () => {
            SidebarItem({ text: i18n("notification_profiles.label"), route: notificationProfileRoute, icon: BullhornIcon });
          });

          PermissionGate("Admin", () => {
            SidebarItem({ text: i18n("api_tokens.label"), route: apiTokensRoute, icon: LockOpenIcon });
          });
        });
      });
    });

    node((node) => {
      document.addEventListener("click", (event) => {
        outsideClicked({ node, event });
      });
    });
  });
};

export const SidebarItem = ({
  text,
  route,
  icon,
}: {
  text: Store<string> | string;
  route: RouteInstance<Record<string, any>>;
  icon: () => void;
}) => {
  const click = bindLinkNavigation({ route, link: Link });

  h("li", () => {
    Link(route, {
      text: text,
      handler: {
        config: { prevent: true, capture: true, stop: true },
        on: { click },
      },
      classList: {
        flex: true,
        "items-center": true,
        "p-2": true,
        "text-gray-900": true,
        "rounded-lg": true,
        "dark:text-white": true,
        "hover:bg-gray-200": true,
        "bg-gray-200": route.$isOpened,
        "dark:bg-squid-ink": route.$isOpened,
        "dark:hover:bg-squid-ink": true,
        group: true,
      },
      fn: () => {
        h("div", () => {
          spec({ classList: ["me-1.5"] });

          icon();
        });
      },
    });
  });
};
