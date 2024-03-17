import { h, list, remap, spec } from "forest";
import { Event, Store, createEvent, sample } from "effector";
import { User } from "@/shared/api";
import { Button } from "@/shared/ui";
import { i18n } from "@/shared/lib/i18n";

export const UsersTable = (users: Store<User[]>, editClicked: Event<number>) => {
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
              "dark:bg-squid-ink",
              "dark:text-gray-200",
            ],
          });

          h("tr", () => {
            h("th", {
              attr: { scope: "col" },
              classList: ["px-6", "py-3", "w-48", "lg:w-72"],
              text: i18n("tables.members.email"),
            });

            h("th", {
              attr: { scope: "col" },
              classList: ["px-6", "py-3", "w-48", "lg:w-72"],
              text: i18n("tables.members.name"),
            });

            h("th", {
              attr: { scope: "col" },
              classList: ["px-6", "py-3"],
              text: i18n("tables.members.role"),
            });

            h("th", {
              attr: { scope: "col" },
              classList: ["px-6", "py-3"],
              text: i18n("tables.members.revoked"),
            });

            h("th", {
              attr: { scope: "col" },
              classList: ["px-6", "py-3"],
              text: i18n("tables.members.actions"),
            });
          });
        });

        h("tbody", () => {
          list(users, ({ store: user }) => {
            h("tr", () => {
              spec({
                classList: {
                  "border-t": true,
                  "w-full": true,
                  "dark:border-shadow-gray": true,
                  "hover:bg-gray-50": true,
                  "dark:hover:bg-slate-gray": true,
                  "text-gray-900": true,
                  "dark:text-gray-200": true,
                  "opacity-50": user.map((u) => u.is_revoked),
                },
              });

              h("td", () => {
                spec({
                  classList: ["px-6", "py-4"],
                });
                h("div", {
                  text: remap(user, "email"),
                });
              });

              h("td", () => {
                spec({
                  classList: ["px-6", "py-4"],
                });
                h("div", {
                  text: remap(user, "name"),
                });
              });

              h("td", () => {
                spec({
                  classList: ["px-6", "py-4"],
                });
                h("div", {
                  text: remap(user, "role"),
                });
              });

              h("td", () => {
                spec({
                  classList: ["px-6", "py-4"],
                });
                h("div", {
                  text: user.map((u) => (u.is_revoked ? "Revoked" : "Active")),
                });
              });

              h("td", () => {
                spec({
                  classList: ["px-6", "py-4"],
                });

                const localEditClicked = createEvent();
                sample({
                  source: remap(user, "id"),
                  clock: localEditClicked,
                  target: editClicked,
                });

                Button({
                  text: i18n("buttons.edit"),
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
