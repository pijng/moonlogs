import { RouteInstance } from "atomic-router";
import { createEvent, sample } from "effector";
import { h, spec } from "forest";

export const CardLink = ({
  title,
  description,
  route,
  link,
}: {
  title: string;
  description: string;
  route: RouteInstance<Record<string, any>>;
  link: string;
}) => {
  /*html*/
  `<a href="#" class="block max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700">
    <h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">Noteworthy technology acquisitions 2021</h5>
    <p class="font-normal text-gray-700 dark:text-gray-400">Here are the biggest enterprise technology acquisitions of 2021 so far, in reverse chronological order.</p>
  </a>`;

  h("a", () => {
    const linkClicked = createEvent<MouseEvent>("link clicked");

    sample({
      source: linkClicked,
      target: route.open,
    });

    spec({
      classList: [
        "block",
        "max-w-xs",
        "p-6",
        "bg-white",
        "border",
        "border-gray-200",
        "rounded-lg",
        "shadow",
        "hover:bg-gray-100",
        "dark:bg-gray-800",
        "dark:border-gray-700",
        "dark:hover:bg-gray-700",
      ],
      attr: { href: link },
      handler: {
        config: { prevent: true },
        on: { click: linkClicked },
      },
    });

    h("h5", {
      classList: ["mb-2", "text-xl", "font-bold", "tracking-tight", "text-gray-900", "dark:text-white"],
      text: title,
    });

    h("p", {
      classList: ["font-normal", "line-clamp-3", "text-gray-700", "dark:text-gray-400"],
      text: description,
    });
  });
};
