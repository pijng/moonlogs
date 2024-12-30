import { Button, ButtonVariant, DownIcon, Dropdown, FilterIcon, FilterItem, KindItem } from "@/shared/ui";
import { Event, Store, createEvent, createStore, sample } from "effector";
import { DOMElement, h, node, spec } from "forest";
import { i18n } from "@/shared/lib";

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
  const dropdownTriggered = createEvent<MouseEvent>();
  const $localVisible = createStore(false);
  const outsideClicked = createEvent<{ node: DOMElement; event: any }>();

  sample({
    source: $localVisible,
    clock: dropdownTriggered,
    fn: (v) => !v,
    target: $localVisible,
  });

  sample({
    source: $localVisible,
    clock: outsideClicked,
    filter: (visible, { node, event }) => !node.contains(event.target) && visible,
    fn: () => false,
    target: $localVisible,
  });

  h("div", () => {
    spec({ classList: ["relative"] });

    Button({
      text: i18n("log_groups.filters.query.label"),
      variant: applied.map<ButtonVariant>((state) => (state ? "default" : "alternative")),
      size: "small",
      event: dropdownTriggered,
      preIcon: FilterIcon,
      postIcon: DownIcon,
    });

    h("div", () => {
      spec({
        visible: $localVisible,
      });

      Dropdown({
        items: filterItems,
        itemChanged: filterChanged,
        currentKind: currentKind,
        kinds: kindItems,
        kindChanged: kindChanged,
      });
    });

    node((node) => {
      document.addEventListener("click", (event) => {
        outsideClicked({ node, event });
      });
    });
  });
};
