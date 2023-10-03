import { schemaModel } from "@/entities/schema";
import { router } from "@/routing";
import { logsRoute, schemaEditRoute } from "@/routing/shared";
import { Button, Header } from "@/shared/ui";
import { combine, createEvent, sample } from "effector";
import { h, spec } from "forest";

const $schema = combine([router.$activeRoutes, schemaModel.$schemas], ([activeRoutes, schemas]) => {
  const schemaName = activeRoutes[0]?.$params.getState().schemaName;
  return schemas.find((s) => s.name === schemaName) || null;
});

const $schemaTitle = $schema.map((s) => s?.title || "");

export const SchemaHeader = () => {
  const routeOpened = createEvent<number>();

  sample({
    source: [logsRoute.$params, schemaModel.$schemas] as const,
    clock: routeOpened,
    fn: ([params, schemas]) => {
      const schemaId = schemas.find((s) => s.name === params.schemaName)?.id;

      return { id: schemaId || 0 };
    },
    target: schemaEditRoute.open,
  });

  h("div", () => {
    spec({
      classList: ["flex", "items-center", "justify-between"],
    });

    Header($schemaTitle);

    h("div", () => {
      spec({
        classList: ["ml-2.5"],
      });

      Button({
        text: "Settings",
        variant: "default",
        size: "extra_small",
        event: routeOpened,
      });
    });
  });
};
