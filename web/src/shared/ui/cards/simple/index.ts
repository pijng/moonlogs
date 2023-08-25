import { RouteInstance } from "atomic-router";
import { Store, createEvent, sample } from "effector";
import { h, spec } from "forest";

export const CardLink = ({
  title,
  description,
  route,
  link,
}: {
  title: Store<string>;
  description: Store<string>;
  route: RouteInstance<any>;
  link: Store<string>;
}) => {
  h("a", () => {
    const linkClicked = createEvent<MouseEvent>();

    sample({
      clock: linkClicked,
      source: link,
      fn: (link) => ({ schemaName: link.split("/").at(-1) }),
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
      classList: ["mb-2", "truncate", "text-lg", "font-bold", "tracking-tight", "text-gray-900", "dark:text-white"],
      text: title,
    });

    h("p", {
      classList: ["font-normal", "line-clamp-3", "text-gray-700", "dark:text-gray-400"],
      text: description,
    });
  });
};
