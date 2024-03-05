import { Translation } from "@/shared/types";

export const en: Translation = {
  profile: {
    label: "Profile",
    name: "Name",
    email: "Email",
    role: "Role",
    tags: "Tags",
    language: "Language",
    theme: "Theme",
  },
  log_groups: {
    label: "Log groups",
    general: "General",
    form: {
      actions: {
        create: "Create log group",
        edit: "Edit log group",
      },
      name: {
        label: "Name",
        hint: "Name - used as a textual identifier for the group. Must be specified in Latin, in lowercase, and with underscores as separators",
      },
      title: {
        label: "Title",
        hint: "Title - used for the human-readable name of the group in the web interface. Group search will also search for groups based on this characteristic",
      },
      description: {
        label: "Description",
        hint: "Description - used for the human-readable description of group details in the web interface. Group search will also search for groups based on this characteristic",
      },
      tag: {
        label: "Select a tag",
      },
      retention_days: {
        label: "Retention days",
        hint: "Retention days - the number of days during which logs will be available after their creation. After the specified number of days elapses, the logs will be deleted. To set an infinite lifespan, specify 0 or leave the field empty",
      },
      group_query_fields: {
        label: "Group query fields",
        hint: "Group query fields - a set of fields by which log grouping will occur",
        fields: {
          title: {
            label: "Title",
            hint: "Title - used for the human-readable name of the field in the web interface for log filtering",
          },
          name: {
            label: "Name",
            hint: "Name - used as a textual identifier for the group. Must be specified in Latin, in lowercase, and with underscores as separators",
          },
        },
      },
      kinds: {
        label: "Kinds",
        hint: "Kinds - a set of select options by which log grouping will occur",
        fields: {
          title: {
            label: "Title",
            hint: "Title - used for the human-readable name of the kind in the web interface for log filtering",
          },
          name: {
            label: "Name",
            hint: "Name - used as a textual identifier for the kind. Must be specified in Latin, in lowercase, and with underscores as separators",
          },
        },
      },
    },
    buttons: {
      create: "Create",
      settings: "Settings",
      save: "Save",
      delete: "Delete",
    },
    alerts: {
      delete: "Are you sure you want to delete this log group?",
    },
  },
};
