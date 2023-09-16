import { h, list, spec } from "forest";
import { withRoute } from "atomic-router-forest";

import { SchemaCard, schemaModel } from "@/entities/schema";
import { homeRoute, logsRoute } from "@/routing/shared";
import { Search } from "@/shared/ui";
import { HeaderWithCreation } from "@/widgets";

export const HomePage = () => {
  h("div", {
    fn() {
      // This allows to show/hide route if page is matched
      // It is required to call `withRoute` inside `h` call
      withRoute(homeRoute);

      HeaderWithCreation("Categories", homeRoute);

      h("div", () => {
        spec({
          classList: ["pt-3"],
        });

        Search(schemaModel.events.queryChanged, schemaModel.$searchQuery);

        h("div", () => {
          spec({
            classList: ["grid", "gap-4", "grid-cols-2", "md:grid-cols-3", "lg:grid-cols-4", "xl:grid-cols-5"],
          });

          list({
            source: schemaModel.$filteredSchemas,
            key: "name",
            fields: ["title", "description", "name"],
            fn({ fields: [title, description, name] }) {
              SchemaCard({ title, description, route: logsRoute, link: name.map((l) => `logs/${l}`) });
            },
          });
        });
      });
    },
  });
};
