import { Link } from "@/shared/routing";
import { RouteInstance, RouteParams } from "atomic-router";
import { Store } from "effector";
import { h } from "forest";

export const CardLink = <T extends RouteParams>({
  title,
  description,
  route,
  params,
}: {
  title: Store<string>;
  description: Store<string>;
  route: RouteInstance<T>;
  params: T;
}) => {
  Link(route, {
    params: params,
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
      "dark:bg-raisin-black",
      "dark:border-shadow-gray",
      "dark:hover:bg-squid-ink",
    ],
    fn() {
      h("h5", {
        classList: ["mb-2", "truncate", "text-lg", "font-bold", "tracking-tight", "text-gray-900", "dark:text-white"],
        text: title,
      });

      h("p", {
        classList: ["font-normal", "line-clamp-3", "text-gray-700", "dark:text-gray-400"],
        text: description,
      });
    },
  });
};
