import { homeRoute } from "@/routing";
import { RouteInstance } from "atomic-router";
import { createEvent, sample } from "effector";
import { h, spec } from "forest";

export const SidebarButton = () => {
  h("button", () => {
    spec({
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
        "dark:hover:bg-gray-700",
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
      classList: [
        "fixed",
        "top-0",
        "left-0",
        "z-40",
        "w-64",
        "h-screen",
        "transition-transform",
        "-translate-x-full",
        "sm:translate-x-0",
      ],
    });

    h("div", () => {
      spec({
        classList: ["h-full", "px-3", "py-4", "overflow-y-auto", "bg-gray-50", "dark:bg-gray-800"],
      });

      h("a", () => {
        const homeClicked = createEvent<MouseEvent>();

        sample({
          source: homeClicked,
          target: homeRoute.open,
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

        SidebarItem("Home", homeRoute);
        SidebarItem("Account", homeRoute);
        SidebarItem("Members", homeRoute);
      });
    });
  });
};

export const SidebarItem = (text: string, route: RouteInstance<Record<string, any>>) => {
  h("li", () => {
    const linkClicked = createEvent<MouseEvent>("link clicked");

    sample({
      clock: linkClicked,
      target: route!.open,
    });

    h("a", {
      attr: { href: "" },
      handler: { on: { click: linkClicked }, config: { prevent: true } },
      classList: [
        "flex",
        "items-center",
        "p-2",
        "text-gray-900",
        "rounded-lg",
        "dark:text-white",
        "hover:bg-gray-100",
        "dark:hover:bg-gray-700",
        "group",
      ],
      text: text,
    });
  });
};
