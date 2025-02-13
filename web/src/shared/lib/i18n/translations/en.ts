import { Translation } from "./types";

export const en: Translation = {
  profile: {
    label: "Profile",
    name: "Name",
    email: "Email",
    role: "Role",
    tags: "Tags",
    language: "Language",
    theme: "Theme",
    clipboard_mode: "Clipboard mode",
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
    alerts: {
      delete: "Are you sure you want to delete this log group?",
    },
    filters: {
      query: {
        label: "Filter",
        kind: "Kind",
      },
      time: {
        label: "Time",
        from: "From",
        to: "To",
      },
      level: {
        label: "Level",
      },
    },
  },
  buttons: {
    create: "Create",
    settings: "Settings",
    save: "Save",
    delete: "Delete",
    edit: "Edit",
    open: "Open",
    log_in: "Log in",
    register: "Register",
    previous: "Previous",
    next: "Next",
    clear: "Clear",
  },
  auth: {
    email: "Email",
    password: "Password",
  },
  validations: {
    required: "Required field",
    invalid_email: "Invalid email",
    passwords_dont_match: "Passwords must match",
  },
  members: {
    label: "Members",
    form: {
      actions: {
        create: "Create member",
        edit: "Edit member",
      },
      name: "Name",
      email: "Email",
      role: "Select a role",
      tag: {
        label: "Select tags",
        hint: "Click to select tags",
      },
      password: "Password",
      confirm_password: "Confirm password",
      revoked: "Revoked",
    },
    alerts: {
      delete: "Are you sure you want to delete this user?",
    },
  },
  tags: {
    label: "Tags",
    form: {
      actions: {
        create: "Create tag",
        edit: "Edit tag",
      },
      name: {
        label: "Name",
        hint: "Name - used for the human-readable name of the tag in the web interface",
      },
      view_order: {
        label: "View order",
        hint: "View order - determines the priority order of the tags. Tags with lower values will be displayed higher in the list of log groups",
      },
    },
    alerts: {
      delete: "Are you sure you want to delete this tag?",
    },
  },
  actions: {
    label: "Custom actions",
    form: {
      actions: {
        create: "Create custom action",
        edit: "Edit custom action",
      },
      name: {
        label: "Name",
        hint: "Name - used for the human-readable name of custom action in the web interface",
      },
      pattern: {
        label: "Pattern",
        hint: "Pattern - a template of a link that will be opened by the user when clicked on an action button. Template variables can be used inside the pattern",
        variables: "You can use the following variables for pattern",
      },
      method: {
        label: "Method",
        hint: "Method - the request method when clicking on an action button",
      },
      conditions: {
        label: "Conditions",
        hint: "Conditions - a list of conditions by which it is determined whether to display the action button in a specific log group",
        fields: {
          attribute: {
            label: "Attribute",
            hint: "Attribute - the attribute of the log group on which the condition is based",
          },
          operation: {
            label: "Operation",
            hint: "Operation - the operation that needs to be performed on the attribute",
          },
          value: {
            label: "Value",
            hint: "Value - the value that is compared to the attribute based on the specified operation",
          },
        },
      },
      schema_name: {
        label: "Schema",
        hint: "Schema - the schema for which the action button needs to be displayed",
      },
      disabled: {
        label: "Disabled",
        hint: "Disabled - permanently disables the display of the action button regardless of conditions",
      },
    },
    alerts: {
      delete: "Are you sure you want to delete this custom action?",
    },
  },
  api_tokens: {
    label: "API tokens",
    form: {
      actions: {
        create: "Create API token",
        edit: "Edit API token",
      },
      name: {
        label: "Name",
        hint: "Name - is used to indicate which service will use this API token. It does not affect the token functionally",
      },
      revoked: "Revoked",
      creation_hint:
        "Your integration API token has been successfully created. Make sure to save it securely now, as it won't be displayed again for security reasons",
    },
    alerts: {
      delete: "Are you sure you want to delete this API token?",
    },
  },
  tables: {
    log_groups: {
      time: "Time",
      level: "Level",
      text: "Text",
      request: "Request",
      response: "Response",
    },
    members: {
      email: "Email",
      name: "Name",
      role: "Role",
      revoked: "Revoked",
      actions: "Actions",
    },
    tags: {
      name: "Name",
      actions: "Actions",
      view_order: "View order",
    },
    alerting_rules: {
      name: "Name",
      enabled: "Enabled",
      severity: "Severity",
      interval: "Interval",
      threshold: "Threshold",
      actions: "Actions",
    },
    api_tokens: {
      name: "Name",
      token: "Token",
      revoked: "Revoked",
      actions: "Actions",
    },
    actions: {
      name: "Name",
      pattern: "Pattern",
      method: "Method",
      conditions: "Conditions",
      schema_name: "Schema name",
      disabled: "Disabled",
      actions: "Actions",
    },
    notification_profiles: {
      name: "Name",
      description: "Description",
      enabled: "Enabled",
      actions: "Actions",
    },
  },
  pagination: {
    first_page: "1",
  },
  components: {
    search: {
      text: "Search",
    },
    sidebar: {
      open: "Open sidebar",
    },
  },
  miscellaneous: {
    loading: "Loading...",
    brand: "Moonlogs",
    blank_option: "—",
    copied_to_clipboard: "Copied!",
    not_found: "The requested resource could not be found",
    forbidden: "You do not have permission to access this resource",
    to_home: "Go to Home page",
    empty_search_result: "No logs were found matching your search. Please adjust the filters or try again later",
  },
  alerting_rules: {
    label: "Alerting Rules",
    form: {
      actions: {
        create: "Create Alerting Rule",
        edit: "Edit Alerting Rule",
      },
      alert: "Alert Conditions",
      name: {
        label: "Name",
        hint: "A unique name for the alert",
      },
      description: {
        label: "Description",
        hint: "A short description of the purpose of this alert",
      },
      enabled: {
        label: "Enabled",
        hint: "Enable or disable the alert",
      },
      severity: {
        label: "Severity",
        hint: "Defines the priority level of the alert",
      },
      interval: {
        label: "Interval",
        hint: "The frequency at which alert conditions are checked. Specified in duration format, e.g., '30s' (30 seconds), '5m' (5 minutes), '1h' (1 hour)",
      },
      condition: {
        label: "Condition",
        hint: "A logical expression that determines when the alert should trigger",
      },
      threshold: {
        label: "Threshold",
        hint: "A value against which actual alert metrics are compared",
      },
      filters: {
        label: "Filters",
        hint: "Additional parameters to refine alert conditions",
        level: {
          label: "Level",
          hint: "Filter by log level",
        },
        schema_name: {
          label: "Schemas",
          hint: "Filter by log schemas. Leave empty to analyze all schemas",
        },
        schema_fields: {
          label: "Schema Fields",
          hint: "Filter by schema fields. Checks only for the presence of a specific field, regardless of its value",
        },
        schema_kinds: {
          label: "Schema Types",
          hint: "Filter by schema types. Checks only for the presence of specific types in schemas. Leave empty to analyze all log schema types",
        },
      },
      aggregation: {
        label: "Aggregation",
        hint: "Settings for grouping and counting events",
        type: {
          label: "Type",
          hint: "Type of data aggregation",
        },
        group_by: {
          label: "Group By",
          hint: "Specify the fields by which events will be grouped",
        },
        time_window: {
          label: "Time Window",
          hint: "Specify the time period for data aggregation. Specified in duration format, e.g., '30s' (30 seconds), '5m' (5 minutes), '1h' (1 hour)",
        },
      },
    },
    alerts: {
      delete: "Are you sure you want to delete this alerting rule?",
    },
  },
  notification_profiles: {
    label: "Alert Manager",
    form: {
      actions: {
        create: "Create Notification Profile",
        edit: "Edit Notification Profile",
      },
      name: {
        label: "Name",
        hint: "A unique name for the notification profile",
      },
      description: {
        label: "Description",
        hint: "A short description to clarify the purpose of this profile",
      },
      rule_name: {
        label: "Alerting Rules",
        hint: "Select the alerting rules associated with this profile",
      },
      enabled: {
        label: "Enabled",
        hint: "Enable or disable the notification profile",
      },
      silence_for: {
        label: "Silence For",
        hint: "The period during which notifications from this profile will not be sent after it has already triggered (e.g., '5m', '1h')",
      },
      url: {
        label: "Route",
        hint: "The URL where notifications will be sent",
      },
      method: {
        label: "Method",
        hint: "HTTP request method",
      },
      headers: {
        label: "Headers",
        hint: "Additional headers that will be included in the HTTP request",
        fields: {
          key: {
            label: "Key",
            hint: "Header name (e.g., 'Authorization')",
          },
          value: {
            label: "Value",
            hint: "Header value (e.g., 'Bearer token')",
          },
        },
      },
      payload: {
        label: "Request Body",
        hint: "The format of the request body that will be sent in the notification",
        variables: "You can use the following variables in the request body",
      },
    },
    alerts: {
      delete: "Are you sure you want to delete this notification profile?",
    },
  },
};
