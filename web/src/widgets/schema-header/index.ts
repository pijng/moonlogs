import { schemaModel } from "@/entities/schema";
import { logsRoute } from "@/routing";
import { combine } from "effector";
import { h } from "forest";

const $schema = combine([logsRoute.$params, schemaModel.$schemas], ([params, schemas]) => {
  return schemas.find((s) => s.name === params.schemaName) || null;
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
