import { InsightsBuilder } from "@/features/insights-builder";
import { h, spec } from "forest";

export const Insights = () => {
  h("div", () => {
    spec({
      classList: [],
    });

    InsightsBuilder();
  });
};
