import { schemaModel } from "@/entities/schema";
import { router } from "@/routing";
import { combine } from "effector";
import { h } from "forest";

const $schema = combine([router.$activeRoutes, schemaModel.$schemas], ([activeRoutes, schemas]) => {
  const schemaName = activeRoutes[0]?.$params.getState().schemaName;
  return schemas.find((s) => s.name === schemaName) || null;
});

const $schemaTitle = $schema.map((s) => s?.title || "");

export const SchemaHeader = () => {
  h("h1", {
    classList: [
      "inline-block",
      "text-2xl",
      "sm:text-3xl",
      "font-extrabold",
      "text-slate-900",
      "tracking-tight",
      "dark:text-slate-200",
    ],
    text: $schemaTitle,
  });
};
