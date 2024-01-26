import { Event, Store, createEvent, createStore, sample } from "effector";
import { DOMProperty, h, spec } from "forest";
import { ErrorHint, Popover } from "@/shared/ui";

type InputType = "text" | "number" | "date" | "checkbox" | "email" | "password";

export const Input = <T extends DOMProperty>({
  value,
  type,
  label,
  required,
  inputChanged,
  visible,
  disabled,
  errorText,
  hint,
}: {
  value?: Store<T>;
  type: InputType;
  label: string;
  required?: boolean;
  inputChanged?: Event<T>;
  visible?: Store<boolean>;
  disabled?: Store<boolean>;
  errorText?: Store<string>;
  hint?: string;
}) => {
  h("div", () => {
    spec({
      classList: { relative: true, "mb-6": true, flex: type === "checkbox", "items-center": type === "checkbox" },
      visible: visible,
    });

    h("div", () => {
      spec({ classList: { flex: true, "items-center": true, "mb-2": type !== "checkbox" } });

      h("label", () => {
        spec({
          classList: {
            block: true,
            "mr-2": type === "checkbox",
            "text-sm": true,
            "font-medium": true,
            "text-gray-900": true,
            "dark:text-white": true,
          },
          text: label,
          attr: { for: label },
        });
      });

      h("div", () => {
        spec({
          classList: ["pl-1"],
          visible: createStore(hint ?? "").map(Boolean),
        });

        Popover({ text: hint });
      });
    });

    h("input", () => {
      const localInputChanged = createEvent<any>();
      sample({
        source: localInputChanged,
        fn: (event) => {
          if (type === "number") {
            try {
              return parseInt(event.target.value, 10);
            } catch {
              return 0;
            }
          }
          if (type === "checkbox") {
            return Boolean(event.target.checked);
          }

          return event.target.value;
        },
        target: [inputChanged],
      });

      spec({
        classList: {
          "bg-gray-50": true,
          border: true,
          "border-gray-300": true,
          "text-gray-900": true,
          "text-sm": true,
          "rounded-lg": true,
          "focus:ring-blue-500": true,
          "focus:border-blue-500": true,
          block: true,
          "w-full": type !== "checkbox",
          "w-4": type === "checkbox",
          "h-4": type === "checkbox",
          "p-2.5": true,
          "dark:bg-gray-700": true,
          "dark:border-gray-600": true,
          "dark:placeholder-gray-400": true,
          "dark:text-white": true,
          "dark:focus:ring-blue-500": true,
          "dark:focus:border-blue-500": true,
        },
        attr: {
          disabled: disabled ?? createStore(false),
          id: label,
          type: type,
          required: Boolean(required),
          value: value ?? createStore<T | null>(null),
          checked: value ?? createStore<T | null>(null),
        },
        handler: { on: { input: localInputChanged } },
      });
    });

    ErrorHint(errorText, errorText?.map(Boolean));
  });
};
