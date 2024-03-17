import { h, list, spec, variant } from "forest";
import { NextIcon, PreviousIcon } from "../icons";
import { Event, Store, combine, createEvent, createStore, sample } from "effector";
import { i18n } from "@/shared/lib/i18n";

const BASE_CLASSES = [
  "flex",
  "items-center",
  "justify-center",
  "px-3",
  "h-8",
  "leading-tight",
  "text-gray-500",
  "bg-white",
  "border",
  "border-gray-300",
  "hover:bg-gray-100",
  "hover:text-gray-700",
  "dark:bg-raisin-black",
  "dark:border-shadow-gray",
  "dark:text-gray-400",
  "dark:hover:bg-squid-ink",
  "dark:hover:text-white",
];

const ACTIVE_CLASSES = [
  "z-10",
  "flex",
  "items-center",
  "justify-center",
  "px-3",
  "h-8",
  "leading-tight",
  "text-blue-600",
  "border",
  "border-blue-300",
  "bg-blue-50",
  "hover:bg-blue-100",
  "hover:text-blue-700",
  "dark:border-shadow-gray",
  "dark:bg-squid-ink",
  "dark:text-white",
];

// hack to avoid multiple page selection bug
// https://github.com/effector/effector/issues/964
const $localCurrentPage = createStore("1");

export const Pagination = (pages: Store<number>, currentPage: Store<string>, pageChanged: Event<string>) => {
  sample({
    source: currentPage,
    target: $localCurrentPage,
  });

  h("nav", () => {
    spec({
      visible: pages.map((pages) => pages > 0),
      classList: ["flex", "pb-2"],
      attr: { "aria-label": "Page navigation" },
    });

    h("ul", () => {
      spec({ classList: ["grid", "grid-cols-9", "items-center", "-space-x-px", "h-8", "text-sm"] });

      const $pagesList = pages.map((count) => new Array(count).fill(null).map((_, i) => i + 1));
      const $truncatedPages = truncatedPages($pagesList, currentPage);

      PreviousPages(pages, pageChanged);
      FirstPage(pages, pageChanged);
      PagesList($truncatedPages, pageChanged);
      LastPage($truncatedPages, $pagesList, pageChanged);
      NextPage(pages, pageChanged);
    });
  });
};

const PagesList = (truncatedPages: Store<number[]>, pageChanged: Event<any>) => {
  list(truncatedPages, ({ store: page }) => {
    h("li", () => {
      const $isActive = combine([$localCurrentPage, page], ([currentPage, page]) => {
        return currentPage === `${page}`;
      });

      const localPageClicked = createEvent<MouseEvent>();

      sample({
        source: page,
        clock: localPageClicked,
        fn: (page) => `${page}`,
        target: pageChanged,
      });

      const $state = $isActive.map<{ state: "active" | "not_active" }>((isActive) => ({
        state: isActive ? "active" : "not_active",
      }));

      variant({
        source: $state,
        key: "state",
        cases: {
          active: () => {
            h("a", {
              attr: { href: "" },
              handler: { on: { click: localPageClicked }, config: { prevent: true } },
              classList: ACTIVE_CLASSES,
              text: page,
            });
          },
          not_active: () => {
            h("a", {
              attr: { href: "" },
              handler: { on: { click: localPageClicked }, config: { prevent: true } },
              classList: BASE_CLASSES,
              text: page,
            });
          },
        },
      });
    });
  });
};

const FirstPage = (pages: Store<number>, pageChanged: Event<any>) => {
  h("li", () => {
    spec({
      visible: pages.map((pages) => pages >= 7),
    });

    const $isActive = combine([$localCurrentPage, createStore("1")], ([currentPage, pastPage]) => {
      return currentPage === pastPage;
    });

    const localPageClicked = createEvent<MouseEvent>();

    sample({
      source: createStore("1"),
      clock: localPageClicked,
      target: pageChanged,
    });

    const $state = $isActive.map<{ state: "active" | "not_active" }>((isActive) => ({
      state: isActive ? "active" : "not_active",
    }));

    variant({
      source: $state,
      key: "state",
      cases: {
        active: () => {
          h("a", {
            attr: { href: "" },
            handler: { on: { click: localPageClicked }, config: { prevent: true } },
            classList: ACTIVE_CLASSES,
            text: i18n("pagination.first_page"),
          });
        },
        not_active: () => {
          h("a", {
            attr: { href: "" },
            handler: { on: { click: localPageClicked }, config: { prevent: true } },
            classList: BASE_CLASSES,
            text: i18n("pagination.first_page"),
          });
        },
      },
    });
  });
};

const LastPage = (truncatedPages: Store<number[]>, pagesList: Store<number[]>, pageChanged: Event<any>) => {
  const $showLastPage = combine(
    [truncatedPages, pagesList],
    ([truncatedPages, pagesList]) => pagesList.length > truncatedPages.length,
  );

  h("li", () => {
    spec({
      visible: $showLastPage,
    });

    const $lastPageNumber = pagesList.map((pagesList) => pagesList[pagesList.length - 1]);

    const $isActive = combine([$localCurrentPage, $lastPageNumber], ([currentPage, pastPage]) => {
      return parseInt(currentPage) === pastPage;
    });

    const localPageClicked = createEvent<MouseEvent>();

    sample({
      source: $lastPageNumber,
      clock: localPageClicked,
      fn: (page) => `${page}`,
      target: pageChanged,
    });

    const $state = $isActive.map<{ state: "active" | "not_active" }>((isActive) => ({
      state: isActive ? "active" : "not_active",
    }));

    variant({
      source: $state,
      key: "state",
      cases: {
        active: () => {
          h("a", {
            attr: { href: "" },
            handler: { on: { click: localPageClicked }, config: { prevent: true } },
            classList: ACTIVE_CLASSES,
            text: $lastPageNumber,
          });
        },
        not_active: () => {
          h("a", {
            attr: { href: "" },
            handler: { on: { click: localPageClicked }, config: { prevent: true } },
            classList: BASE_CLASSES,
            text: $lastPageNumber,
          });
        },
      },
    });
  });
};

const PreviousPages = (pages: Store<number>, pageChanged: Event<any>) => {
  h("li", () => {
    const previousClicked = createEvent<MouseEvent>();
    sample({
      source: [$localCurrentPage, pages] as const,
      clock: previousClicked,
      filter: ([cur]) => parseInt(cur) > 1,
      fn: ([cur]) => {
        const previousPage = parseInt(cur) - 1;
        return `${previousPage}`;
      },
      target: pageChanged,
    });

    h("a", () => {
      spec({
        attr: { href: "" },
        handler: { on: { click: previousClicked }, config: { prevent: true } },
        classList: [
          "flex",
          "items-center",
          "justify-center",
          "px-3",
          "h-8",
          "ml-0",
          "leading-tight",
          "text-gray-500",
          "bg-white",
          "border",
          "border-gray-300",
          "rounded-l-lg",
          "hover:bg-gray-100",
          "hover:text-gray-700",
          "dark:bg-raisin-black",
          "dark:border-shadow-gray",
          "dark:text-gray-400",
          "dark:hover:bg-squid-ink",
          "dark:hover:text-white",
        ],
      });

      h("span", {
        classList: ["sr-only"],
        text: i18n("buttons.previous"),
      });

      PreviousIcon();
    });
  });
};

const NextPage = (pages: Store<number>, pageChanged: Event<any>) => {
  h("li", () => {
    const nextClicked = createEvent<MouseEvent>();
    sample({
      source: [$localCurrentPage, pages] as const,
      clock: nextClicked,
      filter: ([cur, pages]) => parseInt(cur) !== pages,
      fn: ([cur]) => {
        const previousPage = parseInt(cur) + 1;
        return `${previousPage}`;
      },
      target: pageChanged,
    });

    h("a", () => {
      spec({
        attr: { href: "" },
        handler: { on: { click: nextClicked }, config: { prevent: true } },
        classList: [
          "flex",
          "items-center",
          "justify-center",
          "px-3",
          "h-8",
          "leading-tight",
          "text-gray-500",
          "bg-white",
          "border",
          "border-gray-300",
          "rounded-r-lg",
          "hover:bg-gray-100",
          "hover:text-gray-700",
          "dark:bg-raisin-black",
          "dark:border-shadow-gray",
          "dark:text-gray-400",
          "dark:hover:bg-squid-ink",
          "dark:hover:text-white",
        ],
      });

      h("span", {
        classList: ["sr-only"],
        text: i18n("buttons.next"),
      });

      NextIcon();
    });
  });
};

const truncatedPages = (pagesList: Store<number[]>, currentPage: Store<string>) => {
  return combine([pagesList, currentPage], ([pagesList, currentPage]) => {
    if (pagesList.length <= 5) return pagesList;

    const cur = parseInt(currentPage);

    const minPage = Math.max(Math.min(cur - 3, pagesList.length - 6), 1);
    const maxPage = Math.max(cur + 5, minPage);
    const pages = pagesList.slice(minPage, maxPage);

    const sliced = pages.slice(1, 5);
    const nextPageIsLast = sliced[sliced.length - 1] === pagesList.length;

    const lastIndex = nextPageIsLast ? 6 : 5;

    return pages.filter((p) => p !== pagesList.length).slice(0, lastIndex);
  });
};
