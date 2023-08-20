import * as router from "@/shared/lib/router";
import { RouteInstance } from "atomic-router";

export const Link = (route: RouteInstance<Record<string, any>>, text: string) => {
  router.Link(route, {
    classList: ["font-medium", "text-blue-600", "dark:text-blue-500", "hover:underline"],
    text: text,
  });
};
