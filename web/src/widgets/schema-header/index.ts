import { schemaModel } from "@/entities/schema";
import { router } from "@/routing";
import { Header } from "@/shared/ui";
import { combine } from "effector";

const $schema = combine([router.$activeRoutes, schemaModel.$schemas], ([activeRoutes, schemas]) => {
  const schemaName = activeRoutes[0]?.$params.getState().schemaName;
  return schemas.find((s) => s.name === schemaName) || null;
});

const $schemaTitle = $schema.map((s) => s?.title || "");

export const SchemaHeader = () => {
  Header($schemaTitle);
};
