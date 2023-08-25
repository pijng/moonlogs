import { CardLink } from "@/shared/ui";
import { RouteInstance } from "atomic-router";
import { Store } from "effector";

export const SchemaCard = ({
  title,
  description,
  route,
  link,
}: {
  title: Store<string>;
  description: Store<string>;
  route: RouteInstance<{ schemaName: string }>;
  link: Store<string>;
}) => {
  CardLink({
    title: title,
    description: description,
    route: route,
    link: link,
  });
};
