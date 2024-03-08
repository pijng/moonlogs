import * as routing from "@/shared/routing";
import { RouteInstance } from "atomic-router";
import { Store } from "effector";

export const Link = (route: RouteInstance<Record<string, any>>, text: Store<string> | string) => {
  routing.Link(route, {
    classList: ["font-medium", "text-blue-600", "dark:text-blue-500", "hover:underline"],
    text: text,
  });
};
