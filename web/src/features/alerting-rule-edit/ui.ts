import { Button, ErrorHint, Input, Multiselect, Select, Subheader } from "@/shared/ui";
import { h, spec } from "forest";
import { $creationError, events, ruleForm } from "./model";
import { i18n } from "@/shared/lib";
import { createStore } from "effector";
import { AlertingCondition, AlertingSeverity, Level } from "@/shared/api";
import { schemaModel } from "@/entities/schema";
import { AggregationType } from "@/shared/api/alerting-rule";

export const EditAlertingRuleForm = () => {
  h("form", () => {
    Input({
      type: "text",
      label: i18n("alerting_rules.form.name.label"),
      hint: i18n("alerting_rules.form.name.hint"),
      value: ruleForm.fields.name.$value,
      inputChanged: ruleForm.fields.name.changed,
      errorText: ruleForm.fields.name.$errorText,
    });

    Input({
      type: "text",
      label: i18n("alerting_rules.form.description.label"),
      hint: i18n("alerting_rules.form.name.hint"),
      value: ruleForm.fields.description.$value,
      inputChanged: ruleForm.fields.description.changed,
      errorText: ruleForm.fields.description.$errorText,
    });

    Input({
      type: "checkbox",
      label: i18n("alerting_rules.form.enabled.label"),
      value: ruleForm.fields.enabled.$value,
      inputChanged: ruleForm.fields.enabled.changed,
      errorText: ruleForm.fields.enabled.$errorText,
    });

    h("div", () => {
      spec({ classList: ["mb-6", "mt-10"] });
      Subheader(i18n("alerting_rules.form.alert"));
    });

    h("div", () => {
      spec({ classList: ["mb-6"] });

      Select({
        text: i18n("alerting_rules.form.severity.label"),
        hint: i18n("alerting_rules.form.severity.hint"),
        value: ruleForm.fields.severity.$value,
        options: createStore<AlertingSeverity[]>(["Error", "Fatal", "Warn", "Info"]),
        optionSelected: ruleForm.fields.severity.onChange,
        withBlank: createStore(false),
      });
    });

    Input({
      type: "text",
      label: i18n("alerting_rules.form.interval.label"),
      hint: i18n("alerting_rules.form.interval.hint"),
      value: ruleForm.fields.interval.$value,
      inputChanged: ruleForm.fields.interval.changed,
      errorText: ruleForm.fields.interval.$errorText,
    });

    h("div", () => {
      spec({ classList: ["mb-6"] });

      Select({
        text: i18n("alerting_rules.form.condition.label"),
        hint: i18n("alerting_rules.form.condition.hint"),
        value: ruleForm.fields.condition.$value,
        options: createStore<AlertingCondition[]>([">", "<", "=="]),
        optionSelected: ruleForm.fields.condition.onChange,
        withBlank: createStore(false),
      });
    });

    Input({
      type: "number",
      label: i18n("alerting_rules.form.threshold.label"),
      hint: i18n("alerting_rules.form.threshold.hint"),
      value: ruleForm.fields.threshold.$value,
      inputChanged: ruleForm.fields.threshold.changed,
      errorText: ruleForm.fields.threshold.$errorText,
    });

    h("div", () => {
      spec({ classList: ["mb-6", "mt-12"] });
      Subheader(i18n("alerting_rules.form.filters.label"));
    });

    h("div", () => {
      spec({ classList: ["mb-6"] });

      Select({
        text: i18n("alerting_rules.form.filters.level.label"),
        hint: i18n("alerting_rules.form.filters.level.hint"),
        value: ruleForm.fields.filter_level.$value,
        options: createStore<Level[]>(["Info", "Error", "Warn", "Debug", "Trace", "Fatal"]),
        optionSelected: ruleForm.fields.filter_level.onChange,
        withBlank: createStore(false),
      });
    });

    h("div", () => {
      spec({ classList: ["mb-6"] });

      Multiselect({
        text: i18n("alerting_rules.form.filters.schema_name.label"),
        hint: i18n("alerting_rules.form.filters.schema_name.hint"),
        options: schemaModel.$schemas.map((schema) => schema.map((s) => ({ name: s.title, id: s.id }))),
        selectedOptions: ruleForm.fields.filter_schema_ids.$value,
        optionChecked: events.schemaChecked,
        optionUnchecked: events.schemaUnchecked,
      });
    });

    h("div", () => {
      spec({ classList: ["mb-6"] });

      Multiselect({
        text: i18n("alerting_rules.form.filters.schema_fields.label"),
        hint: i18n("alerting_rules.form.filters.schema_fields.hint"),
        options: schemaModel.$schemas.map((schemas) => {
          const allFields = schemas.flatMap((schema) => {
            const fields = schema.fields || [];
            return fields.map((f) => ({ name: f.title, id: f.name }));
          });

          return Array.from(new Map(allFields.map((f) => [f.id, f])).values());
        }),
        selectedOptions: ruleForm.fields.filter_schema_fields.$value,
        optionChecked: events.schemaFieldChecked,
        optionUnchecked: events.schemaFieldUnchecked,
      });
    });

    h("div", () => {
      spec({ classList: ["mb-6"] });

      Multiselect({
        text: i18n("alerting_rules.form.filters.schema_kinds.label"),
        hint: i18n("alerting_rules.form.filters.schema_kinds.hint"),
        options: schemaModel.$schemas.map((schemas) => {
          const allKinds = schemas.flatMap((schema) => {
            const kinds = schema.kinds || [];
            return kinds.map((k) => ({ name: k.title, id: k.name }));
          });

          return Array.from(new Map(allKinds.map((f) => [f.id, f])).values());
        }),
        selectedOptions: ruleForm.fields.filter_schema_kinds.$value,
        optionChecked: events.schemaKindChecked,
        optionUnchecked: events.schemaKindUnchecked,
      });
    });

    h("div", () => {
      spec({ classList: ["mb-6", "mt-12"] });
      Subheader(i18n("alerting_rules.form.aggregation.label"));
    });

    h("div", () => {
      spec({ classList: ["mb-6"] });

      Select({
        text: i18n("alerting_rules.form.aggregation.type.label"),
        hint: i18n("alerting_rules.form.aggregation.type.hint"),
        value: ruleForm.fields.aggregation_type.$value,
        options: createStore<AggregationType[]>(["count"]),
        optionSelected: ruleForm.fields.aggregation_type.onChange,
        withBlank: createStore(false),
      });
    });

    h("div", () => {
      spec({ classList: ["mb-6"] });

      Multiselect({
        text: i18n("alerting_rules.form.aggregation.group_by.label"),
        hint: i18n("alerting_rules.form.aggregation.group_by.hint"),
        selectedOptions: ruleForm.fields.aggregation_group_by.$value,
        options: ruleForm.fields.filter_schema_fields.$value.map((fields) => {
          const defFields = [{ name: "schema_name", id: "schema_name" }];
          const schemasFields = fields.map((f) => ({ name: f, id: f }));
          return defFields.concat(schemasFields);
        }),
        optionChecked: events.aggregationGroupByChecked,
        optionUnchecked: events.aggregationGroupByUnchecked,
        errorText: ruleForm.fields.aggregation_group_by.$errorText,
      });
    });

    Input({
      type: "text",
      label: i18n("alerting_rules.form.aggregation.time_window.label"),
      hint: i18n("alerting_rules.form.aggregation.time_window.hint"),
      value: ruleForm.fields.aggregation_time_window.$value,
      inputChanged: ruleForm.fields.aggregation_time_window.onChange,
      errorText: ruleForm.fields.aggregation_time_window.$errorText,
    });

    h("div", () => {
      spec({ classList: ["flex", "justify-start", "space-x-2"] });

      Button({
        text: i18n("buttons.save"),
        event: ruleForm.submit,
        size: "base",
        prevent: true,
        variant: "default",
      });

      Button({
        text: i18n("buttons.delete"),
        event: events.deleteRuleClicked,
        size: "base",
        prevent: true,
        variant: "delete",
      });
    });

    ErrorHint($creationError, $creationError.map(Boolean));
  });
};
