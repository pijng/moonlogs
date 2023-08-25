import { Search } from "@/shared/ui";
import { Event, Store } from "effector";
import { h, spec } from "forest";

export const SearchBar = (inputChanged: Event<string>, $searchQuery: Store<string>) => {
  h("div", () => {
    spec({
      classList: ["flex", "max-w-2xl", "w-full", "flex-row"],
    });

    h("div", () => {
      spec({
        classList: ["w-full"],
      });

      Search(inputChanged, $searchQuery);
    });

    // h("div", () => {
    //   Button({ text: "Search", variant: "default", event: inputSubmitted });
    // });
  });
};
