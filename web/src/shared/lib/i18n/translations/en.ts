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
    api_tokens: {
      name: "Name",
      token: "Token",
      revoked: "Revoked",
      actions: "Actions ",
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
    not_found: "The requested resource could not be found",
    forbidden: "You do not have permission to access this resource",
    to_home: "Go to Home page",
  },
};
