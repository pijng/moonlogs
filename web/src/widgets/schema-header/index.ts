import { schemaModel } from "@/entities/schema";
import { router } from "@/routing";
import { Button, GearIcon, Header } from "@/shared/ui";
import { combine } from "effector";
import { h, spec } from "forest";

const $schema = combine([router.$activeRoutes, schemaModel.$schemas], ([activeRoutes, schemas]) => {
  const schemaName = activeRoutes[0]?.$params.getState().schemaName;
  return schemas.find((s) => s.name === schemaName) || null;
});

const $schemaTitle = $schema.map((s) => s?.title || "");

export const SchemaHeader = () => {
  h("div", () => {
    spec({
      classList: ["flex", "items-center"],
    });

    Header($schemaTitle);

    h("div", () => {
      spec({
        classList: ["ml-2.5"],
      });

      Button({
        text: "",
        variant: "light",
        style: "round",
        size: "small",
        preIcon: GearIcon,
      });
    });
  });
};
