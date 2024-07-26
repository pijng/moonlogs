import { Search, Subheader } from "@/shared/ui";
import { h, list, remap, spec } from "forest";
import { $generalSchemas, $groupedTaggedSchemas, $searchQuery, $sortedTags, queryChanged } from "./model";
import { combine } from "effector";
import { SchemaCard } from "@/entities/schema";
import { i18n } from "@/shared/lib";

export const SchemasList = () => {
  h("div", () => {
    spec({
      classList: ["pt-3"],
    });

    Search(queryChanged, $searchQuery);

    list($sortedTags, ({ store: tag }) => {
      h("div", () => {
        const $tagSchemas = combine(tag, $groupedTaggedSchemas, (t, schemas) => schemas[t.id] || []);

        spec({ classList: ["mt-2", "mb-9"], visible: $tagSchemas.map((s) => Boolean(s?.length)) });

        Subheader(remap(tag, "name"));

        h("div", () => {
          spec({
            classList: [
              "mt-3",
              "max-w-7xl",
              "grid",
              "gap-4",
              "grid-cols-2",
              "md:grid-cols-3",
              "lg:grid-cols-4",
              "xl:grid-cols-5",
            ],
          });

          list({
            source: $tagSchemas,
            key: "id",
            fields: ["title", "description", "name"],
            fn({ fields: [title, description, name] }) {
              SchemaCard({ title, description, name });
            },
          });
        });
      });
    });

    h("div", () => {
      spec({ classList: ["mt-2", "mb-9"], visible: $generalSchemas.map((s) => s?.length > 0) });

      Subheader(i18n("log_groups.general"));

      h("div", () => {
        spec({
          classList: ["mt-3", "max-w-7xl", "grid", "gap-4", "grid-cols-2", "md:grid-cols-3", "lg:grid-cols-4", "xl:grid-cols-5"],
        });

        list({
          source: $generalSchemas,
          key: "id",
          fields: ["title", "description", "name"],
          fn({ fields: [title, description, name] }) {
            SchemaCard({ title, description, name });
          },
        });
      });
    });
  });
};
