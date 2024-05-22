import { Button, ErrorHint, Input, Label, PlusIcon, Select, Text, TrashIcon } from "@/shared/ui";
import { h, list, remap, spec } from "forest";
import { $creationError, events, actionForm, $currentSchema } from "./model";
import { combine, createEvent, createStore, sample } from "effector";
import { trigger } from "@/shared/lib";
import { i18n } from "@/shared/lib/i18n";
import { schemaModel } from "@/entities/schema";
import { ActionToCreate, Condition } from "@/shared/api";

export const NewActionForm = () => {
  h("form", () => {
    Input({
      type: "text",
      label: i18n("actions.form.name.label"),
      required: true,
      value: actionForm.fields.name.$value,
      inputChanged: actionForm.fields.name.changed,
      errorText: actionForm.fields.name.$errorText,
      hint: i18n("actions.form.name.hint"),
    });

    h("div", () => {
      spec({ classList: ["mb-6"] });

      Select({
        text: i18n("actions.form.schema_name.label"),
        value: actionForm.fields.schema_id.$value,
        options: schemaModel.$schemas,
        optionSelected: events.schemaSelected,
      });
    });

    h("div", () => {
      spec({ visible: actionForm.fields.schema_id.$value.map(Boolean) });

      const $schemaQueries = $currentSchema.map((s) => s.fields.map((f) => f.name));
      const $attributeList = combine($schemaQueries, (queries) => queries.concat(["kind", "level"]));

      h("div", () => {
        const attributesText = $attributeList.map((attr) => attr.map((a) => `{{${a}}}`).join(", "));
        const $label = combine(i18n("actions.form.pattern.variables"), attributesText, (label, attrs) => `${label}: ${attrs}`);

        Text({ text: $label });
      });

      Input({
        type: "text",
        label: i18n("actions.form.pattern.label"),
        required: true,
        value: actionForm.fields.pattern.$value,
        inputChanged: actionForm.fields.pattern.changed,
        errorText: actionForm.fields.pattern.$errorText,
        hint: i18n("actions.form.pattern.hint"),
      });

      h("div", () => {
        spec({ classList: ["mb-6"] });

        Select({
          text: i18n("actions.form.method.label"),
          value: actionForm.fields.method.$value,
          options: createStore<ActionToCreate["method"][]>(["GET"]),
          optionSelected: events.schemaSelected,
          withBlank: createStore(false),
        });
      });

      Input({
        type: "checkbox",
        label: i18n("actions.form.disabled.label"),
        value: actionForm.fields.disabled.$value,
        inputChanged: actionForm.fields.disabled.changed,
        errorText: actionForm.fields.disabled.$errorText,
      });

      h("div", () => {
        spec({ classList: ["relative", "flex", "items-center", "mb-4", "pt-2"] });

        Label({
          text: i18n("actions.form.conditions.label"),
          hint: i18n("actions.form.conditions.hint"),
        });

        h("div", () => {
          spec({ classList: ["ml-1"] });

          Button({
            variant: "default",
            prevent: true,
            style: "round",
            size: "extra_small",
            event: events.addCondition,
            preIcon: PlusIcon,
          });
        });
      });

      h("div", () => {
        list(actionForm.fields.conditions.$value, ({ store: conditionField, key: idx }) => {
          const attributeChanged = createEvent<string>();
          const operationChanged = createEvent<Condition["operation"]>();
          const valueChanged = createEvent<string>();

          sample({
            source: idx,
            clock: attributeChanged,
            fn: (idx, attribute) => ({ attribute: attribute, idx: idx }),
            target: events.conditionAttributeChanged,
          });

          sample({
            source: idx,
            clock: operationChanged,
            fn: (idx, operation) => ({ operation: operation, idx: idx }),
            target: events.conditionOperationChanged,
          });

          sample({
            source: idx,
            clock: valueChanged,
            fn: (idx, value) => ({ value: value, idx: idx }),
            target: events.conditionValueChanged,
          });

          h("div", () => {
            spec({
              classList: ["grid", "gap-3", "place-items-stretch"],
              style: {
                gridTemplateColumns: "14fr 1fr 14fr 1fr",
              },
            });

            h("div", () => {
              spec({ classList: ["mb-6"] });

              Select({
                text: i18n("actions.form.conditions.fields.attribute.label"),
                value: remap(conditionField, "attribute"),
                options: $attributeList,
                optionSelected: attributeChanged,
                withBlank: createStore(false),
              });
            });

            h("div", () => {
              spec({ classList: ["mb-6"] });

              Select({
                text: i18n("actions.form.conditions.fields.operation.label"),
                value: remap(conditionField, "operation"),
                options: createStore<Condition["operation"][]>(["==", "!=", "<", "<=", ">", ">="]),
                optionSelected: operationChanged,
                withBlank: createStore(false),
              });
            });

            Input({
              type: "text",
              label: i18n("actions.form.conditions.fields.value.label"),
              required: true,
              value: remap(conditionField, "value"),
              inputChanged: valueChanged,
              errorText: actionForm.fields.conditions.$errorText,
              hint: i18n("actions.form.conditions.fields.value.hint"),
            });

            Button({
              event: trigger({ source: idx, target: events.removeCondition }),
              size: "plain",
              prevent: true,
              variant: "delete_icon",
              preIcon: TrashIcon,
            });
          });
        });
      });
    });

    h("div", () => {
      spec({ classList: ["pt-4"] });

      Button({
        text: i18n("buttons.create"),
        event: actionForm.submit,
        size: "base",
        prevent: true,
        variant: "default",
      });
    });

    ErrorHint($creationError, $creationError.map(Boolean));
  });
};
