import { Button, ButtonVariant, DownIcon, Dropdown } from "@/shared/ui";
import { Event, Store } from "effector";
import { h, spec } from "forest";
import { $filterIsOpened, filterClicked } from "./model";
import { FilterIcon } from "@/shared/ui";

export type FilterItem = {
  name: string;
  key: string;
  value: string;
};

export type KindItem = {
  name: string;
  title: string;
};

export const Filter = ({
  filterItems,
  filterChanged,
  currentKind,
  kindItems,
  kindChanged,
  applied,
}: {
  filterItems: Store<FilterItem[]>;
  filterChanged: Event<Record<string, any>>;
  currentKind: Store<string>;
  kindItems: Store<KindItem[]>;
  kindChanged: Event<string>;
  applied: Store<boolean>;
}) => {
  Button({
    text: "Filter",
    variant: applied.map<ButtonVariant>((state) => (state ? "default" : "alternative")),
    size: "small",
    event: filterClicked,
    preIcon: FilterIcon,
    postIcon: DownIcon,
  });

  h("div", () => {
    spec({
      visible: $filterIsOpened,
    });

    Dropdown({
      items: filterItems,
      itemChanged: filterChanged,
      currentKind: currentKind,
      kinds: kindItems,
      kindChanged: kindChanged,
    });
  });
};
