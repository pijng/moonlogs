import { logsRoute } from "@/routing/shared";
import { CardLink } from "@/shared/ui";
import { Store } from "effector";

export const SchemaCard = ({
  title,
  description,
  name,
}: {
  title: Store<string>;
  description: Store<string>;
  name: Store<string>;
}) => {
  CardLink({
    title: title,
    description: description,
    route: logsRoute,
    params: { schemaName: name },
  });
};
