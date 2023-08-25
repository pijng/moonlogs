import { h, list, spec } from "forest";
import { withRoute } from "atomic-router-forest";

import { SchemaCard, schemaModel } from "@/entities/schema";
import { homeRoute, logsRoute } from "@/routing";
import { Search } from "@/shared/ui";

export const HomePage = () => {
  h("div", {
    fn() {
      // This allows to show/hide route if page is matched
      // It is required to call `withRoute` inside `h` call
      withRoute(homeRoute);

      Search(schemaModel.events.queryChanged, schemaModel.$searchQuery);

      h("div", () => {
        spec({
          classList: ["grid", "grid-cols-2", "gap-4", "md:grid-cols-3", "lg:grid-cols-4", "xl:grid-cols-5"],
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
    },
  });
};
