import { h, list, remap, spec } from "forest";
import { Event, Store, createEvent, sample } from "effector";
import { ApiToken } from "@/shared/api";
import { Button } from "@/shared/ui";

export const ApiTokensTable = (apiTokens: Store<ApiToken[]>, editClicked: Event<number>) => {
  h("div", () => {
    spec({
      classList: ["antialiased"],
    });

    h("div", () => {
      spec({
        classList: ["relative", "overflow-x-auto", "sm:rounded-b-lg"],
      });

      h("table", () => {
        spec({
          classList: ["min-w-full", "w-max", "text-sm", "text-left", "table-fixed"],
        });

        h("thead", () => {
          spec({
            classList: [
              "text-xs",
              "text-gray-700",
              "uppercase",
              "bg-gray-50",
              "w-full",
              "dark:bg-gray-700",
              "dark:text-gray-200",
            ],
          });

          h("tr", () => {
            h("th", {
              attr: { scope: "col" },
              classList: ["px-6", "py-3"],
              text: "Name",
            });

            h("th", {
              attr: { scope: "col" },
              classList: ["px-6", "py-3", "w-48", "lg:w-72"],
              text: "Token",
            });

            h("th", {
              attr: { scope: "col" },
              classList: ["px-6", "py-3"],
              text: "Revoked",
            });

            h("th", {
              attr: { scope: "col" },
              classList: ["px-6", "py-3"],
              text: "Actions",
            });
          });
        });

        h("tbody", () => {
          list(apiTokens, ({ store: apiToken }) => {
            h("tr", () => {
              spec({
                classList: [
                  "border-t",
                  "w-full",
                  "dark:border-gray-700",
                  "hover:bg-gray-50",
                  "dark:hover:bg-gray-600",
                  "text-gray-900",
                  "dark:text-gray-200",
                ],
              });

              h("td", () => {
                spec({
                  classList: ["px-6", "py-4"],
                });
                h("div", {
                  text: remap(apiToken, "name"),
                });
              });

              h("td", () => {
                spec({
                  classList: ["px-6", "py-4"],
                });
                h("div", {
                  text: apiToken.map((t) => t.token || new Array(32).fill("*").join("")),
                });
              });

              h("td", () => {
                spec({
                  classList: ["px-6", "py-4"],
                });
                h("div", {
                  text: remap(apiToken, "is_revoked"),
                });
              });

              h("td", () => {
                spec({
                  classList: ["px-6", "py-4"],
                });

                const localEditClicked = createEvent();
                sample({
                  source: remap(apiToken, "id"),
                  clock: localEditClicked,
                  target: editClicked,
                });

                Button({
                  text: "Edit",
                  event: localEditClicked,
                  variant: "default",
                  size: "extra_small",
                });
              });
            });
          });
        });
      });
    });
  });
};
