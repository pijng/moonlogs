import { FilterItem } from "@/features";
import { Event, Store } from "effector";
import { h, remap, spec } from "forest";

export const FilledFloatingInput = (item: Store<FilterItem>, inputChanged: Event<any>) => {
  h("div", () => {
    spec({
      classList: ["relative"],
    });

    h("input", {
      attr: {
        type: "text",
        id: item.map((i) => `input_${i.key}`),
        name: item.map((i) => `input_${i.key}`),
        placeholder: "",
        value: item.map((i) => decodeURIComponent(i.value)),
      },
      handler: { on: { input: inputChanged } },
      classList: [
        "block",
        "rounded-lg",
        "px-2.5",
        "pb-2.5",
        "pt-5",
        "w-full",
        "text-sm",
        "text-gray-900",
        "bg-gray-50",
        "dark:bg-squid-ink",
        "border-1",
        "border-gray-300",
        "appearance-none",
        "dark:text-white",
        "dark:border-gray-600",
        "dark:focus:border-blue-500",
        "focus:outline-none",
        "focus:ring-0",
        "focus:border-blue-600",
        "peer",
      ],
    });

    h("label", {
      attr: { for: item.map((i) => `input_${i.key}`) },
      classList: [
        "absolute",
        "text-sm",
        "text-gray-500",
        "dark:text-gray-400",
        "duration-300",
        "transform",
        "-translate-y-4",
        "scale-75",
        "top-4",
        "z-10",
        "origin-[0]",
        "start-2.5",
        "peer-focus:text-blue-600",
        "peer-focus:dark:text-blue-500",
        "peer-placeholder-shown:scale-100",
        "peer-placeholder-shown:translate-y-0",
        "peer-focus:scale-75",
        "peer-focus:-translate-y-4",
        "rtl:peer-focus:translate-x-1/4",
        "rtl:peer-focus:left-auto",
      ],
      text: remap(item, "name"),
    });
  });
};
