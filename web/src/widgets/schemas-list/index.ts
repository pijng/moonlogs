import { Search, Subheader } from "@/shared/ui";
import { h, list, spec } from "forest";
import { $groupedFilteredSchemas, $searchQuery, queryChanged } from "./model";
import { combine } from "effector";
import { tagModel } from "@/entities/tag";
import { SchemaCard } from "@/entities/schema";
import { i18n } from "@/shared/lib/i18n";

export const SchemasList = () => {
  h("div", () => {
    spec({
      classList: ["pt-3"],
    });

    Search(queryChanged, $searchQuery);

    list(
      $groupedFilteredSchemas.map((groups) => Object.entries(groups)),
      ({ store: schemaGroups }) => {
        h("div", () => {
          spec({ classList: ["mt-2", "mb-9"] });

          const tagId = schemaGroups.map((g) => g[0]);
          const tagName = combine([tagModel.$tags, tagId, i18n("log_groups.general")], ([tags, tagId, general]) => {
            return tags.find((t) => String(t.id) === tagId)?.name || general;
          });

          Subheader(tagName);

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
              source: schemaGroups.map((g) => g[1]),
              key: "id",
              fields: ["title", "description", "name"],
              fn({ fields: [title, description, name] }) {
                SchemaCard({ title, description, name });
              },
            });
          });
        });
      },
    );
  });
};
