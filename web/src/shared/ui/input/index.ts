import { Event, Store, createEvent, createStore, sample } from "effector";
import { h, spec } from "forest";
import { ErrorHint } from "@/shared/ui";

export const Input = ({
  value,
  type,
  label,
  required,
  inputChanged,
  errorVisible,
  errorText,
}: {
  value?: Store<string>;
  type: string;
  label: string;
  required?: boolean;
  inputChanged: Event<any>;
  errorVisible?: Store<boolean>;
  errorText?: Store<string>;
}) => {
  h("div", () => {
    spec({
      classList: ["mb-6"],
    });

    h("label", () => {
      spec({
        classList: ["block", "mb-2", "text-sm", "font-medium", "text-gray-900", "dark:text-white"],
        text: label,
      });
    });

    h("input", () => {
      const localInputChanged = createEvent<any>();
      sample({
        source: localInputChanged,
        fn: (event) => event.target.value,
        target: inputChanged,
      });

      spec({
        classList: [
          "bg-gray-50",
          "border",
          "border-gray-300",
          "text-gray-900",
          "text-sm",
          "rounded-lg",
          "focus:ring-blue-500",
          "focus:border-blue-500",
          "block",
          "w-full",
          "p-2.5",
          "dark:bg-gray-700",
          "dark:border-gray-600",
          "dark:placeholder-gray-400",
          "dark:text-white",
          "dark:focus:ring-blue-500",
          "dark:focus:border-blue-500",
        ],
        attr: { type: type, required: Boolean(required), value: value ?? createStore("") },
        handler: { on: { input: localInputChanged } },
      });
    });

    ErrorHint(errorText, errorVisible);
  });
};
