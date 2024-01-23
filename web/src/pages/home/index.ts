import { h, list, spec } from "forest";
import { withRoute } from "atomic-router-forest";

import { SchemaCard, schemaModel } from "@/entities/schema";
import { homeRoute, schemaCreateRoute } from "@/routing/shared";
import { Search, Subheader } from "@/shared/ui";
import { HeaderWithCreation } from "@/widgets";
import { tagModel } from "@/entities/tag";
import { combine } from "effector";

export const HomePage = () => {
  h("div", {
    fn() {
      // This allows to show/hide route if page is matched
      // It is required to call `withRoute` inside `h` call
      withRoute(homeRoute);

      HeaderWithCreation("Log groups", schemaCreateRoute);

      h("div", () => {
        spec({
          classList: ["pt-3"],
        });

        Search(schemaModel.events.queryChanged, schemaModel.$searchQuery);

        list(
          schemaModel.$groupedFilteredSchemas.map((groups) => Object.entries(groups)),
          ({ store: schemaGroups }) => {
            h("div", () => {
              spec({ classList: ["mt-2", "mb-9"] });

              const tagId = schemaGroups.map((g) => g[0]);
              const tagName = combine(tagModel.$tags, tagId, (tags, tagId) => {
                return tags.find((t) => String(t.id) === tagId)?.name || "General";
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
    },
  });
};
