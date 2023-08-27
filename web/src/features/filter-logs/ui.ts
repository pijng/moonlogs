import { Button, DownIcon, Dropdown } from "@/shared/ui";
import { Event, Store } from "effector";
import { h, spec } from "forest";
import { $filterIsOpened, filterClicked } from "./model";
import { FilterIcon } from "@/shared/ui/icons";

export type FilterItem = {
  name: string;
  key: string;
  value: string;
};

export const Filter = (items: Store<FilterItem[]>, filterChanged: Event<Record<string, any>>) => {
  h("div", () => {
    Button({
      text: "Filter",
      variant: "alternative",
      size: "small",
      event: filterClicked,
      preIcon: FilterIcon,
      postIcon: DownIcon,
    });

    h("div", () => {
      spec({
        visible: $filterIsOpened,
      });

      Dropdown(items, filterChanged);
    });
  });
};
