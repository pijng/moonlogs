import { FilterItem } from "@/features";
import { Event, Store } from "effector";
import { h, remap, spec } from "forest";

export const FloatingInput = (item: Store<FilterItem>, inputChanged: Event<any>) => {
  h("div", () => {
    spec({
      classList: ["flex", "flex-col", "z-0"],
    });

    h("input", {
      attr: {
        type: "text",
        id: item.map((i) => `input_${i.key}`),
        name: item.map((i) => `input_${i.key}`),
        placeholder: "_",
        value: item.map((i) => decodeURIComponent(i.value)),
      },
      handler: { on: { input: inputChanged } },
      classList: [
        "block",
        "py-0.5",
        "px-0",
        "w-full",
        "text-sm",
        "text-gray-900",
        "bg-transparent",
        "border-0",
        "border-b-2",
        "border-gray-300",
        "appearance-none",
        "dark:text-white",
        "dark:border-slate-gray",
        "dark:focus:border-blue-500",
        "focus:outline-none",
        "focus:ring-0",
        "focus:border-blue-600",
        "placeholder:opacity-0",
        "focus:outline",
        "outline-1",
        "peer",
      ],
    });

    h("label", {
      attr: { for: item.map((i) => `input_${i.key}`) },
      classList: [
        "order-first",
        "text-sm",
        "text-gray-500",
        "dark:text-gray-400",
        "duration-300",
        "transform",
        "scale-75",
        "top-3",
        "-z-10",
        "origin-[0]",
        "peer-focus:left-0",
        "peer-focus:text-blue-600",
        "peer-focus:dark:text-blue-500",
        "peer-placeholder-shown:scale-100",
        "peer-focus:scale-75",
        "peer-placeholder-shown:translate-y-4",
        "peer-focus:-translate-y-0",
      ],
      text: remap(item, "name"),
    });
  });
};
