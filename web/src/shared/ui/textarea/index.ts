import { Event, Store, createEvent, createStore, sample } from "effector";
import { DOMProperty, h, spec } from "forest";
import { ErrorHint, Popover } from "@/shared/ui";

type InputType = "text" | "number" | "date" | "datetime-local" | "checkbox" | "email" | "password";

export const TextArea = <T extends DOMProperty>({
  value,
  label,
  required,
  rows,
  autoHeight,
  inputChanged,
  visible,
  disabled,
  errorText,
  hint,
}: {
  value?: Store<T>;
  type: InputType;
  label: Store<string> | string;
  required?: boolean;
  autoHeight?: boolean;
  rows?: number;
  inputChanged?: Event<T>;
  visible?: Store<boolean>;
  disabled?: Store<boolean>;
  errorText?: Store<string>;
  hint?: Store<string> | string;
}) => {
  h("div", () => {
    spec({
      classList: ["relative", "mb-6"],
      visible: visible,
    });

    h("div", () => {
      spec({ classList: ["flex", "items-center", "mb-2"] });

      h("label", () => {
        spec({
          classList: ["block", "text-sm", "font-medium", "text-gray-900", "dark:text-white"],
          text: label,
          attr: { for: label },
        });
      });

      h("div", () => {
        spec({
          classList: ["pl-1", "relative"],
          visible: createStore(hint ?? "").map(Boolean),
        });

        Popover({ text: hint });
      });
    });

    h("textarea", () => {
      const localInputChanged = createEvent<any>();
      sample({
        source: localInputChanged,
        fn: (event) => event.target.value,
        target: [inputChanged],
      });

      const $hasError = createStore(false);

      sample({
        source: localInputChanged,
        fn: (event) => {
          try {
            JSON.parse(event.target.value);
            return false;
          } catch {
            return true;
          }
        },
        target: $hasError,
      });

      const height = createStore<string>("");
      sample({
        source: localInputChanged,
        filter: () => Boolean(autoHeight),
        fn: (event) => `${event.target.scrollHeight + 3}px`,
        target: height,
      });

      spec({
        classList: {
          "resize-none": true,
          "overflow-hidden": true,
          "bg-gray-50": true,
          border: true,
          "border-gray-300": $hasError.map((err) => !err),
          "border-red-600": $hasError,
          "text-gray-900": true,
          "text-sm": true,
          "rounded-lg": true,
          "focus:ring-blue-500": $hasError.map((err) => !err),
          "focus:border-blue-500": $hasError.map((err) => !err),
          "focus:border-red-600": $hasError,
          "focus:outline": true,
          "border-2": true,
          "outline-0": true,
          block: true,
          "w-full": true,
          "p-2.5": true,
          "dark:bg-squid-ink": true,
          "dark:border-slate-gray": $hasError.map((err) => !err),
          "dark:border-red-400": $hasError,
          "dark:placeholder-gray-400": true,
          "dark:text-white": true,
          "dark:focus:ring-blue-500": $hasError.map((err) => !err),
          "dark:focus:border-blue-500": $hasError.map((err) => !err),
          "dark:focus:border-red-400": $hasError,
        },
        attr: {
          rows: rows || 5,
          disabled: disabled ?? createStore(false),
          id: label,
          required: Boolean(required),
          value: value ?? createStore<T | null>(null),
          checked: value ?? createStore<T | null>(null),
        },
        style: { height: height },
        handler: { on: { input: localInputChanged } },
      });
    });

    ErrorHint(errorText, errorText?.map(Boolean));
  });
};
