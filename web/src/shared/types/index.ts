export type Keys<T extends Record<string, any>> = Extract<keyof T, string>;

export type PathsToStringProps<T> = T extends string
  ? []
  : {
      [K in Extract<keyof T, string>]: [K, ...PathsToStringProps<T[K]>];
    }[Extract<keyof T, string>];

export type Join<T extends string[]> = T extends []
  ? never
  : T extends [infer F]
  ? F
  : T extends [infer F, ...infer R]
  ? F extends string
    ? `${F}.${Join<Extract<R, string[]>>}`
    : never
  : string;

export type Translation = {
  profile: {
    label: string;
    name: string;
    email: string;
    role: string;
    tags: string;
    language: string;
    theme: string;
  };
  log_groups: {
    label: string;
    general: string;
    form: {
      actions: {
        create: string;
        edit: string;
      };
      name: {
        label: string;
        hint: string;
      };
      title: {
        label: string;
        hint: string;
      };
      description: {
        label: string;
        hint: string;
      };
      tag: {
        label: string;
      };
      retention_days: {
        label: string;
        hint: string;
      };
      group_query_fields: {
        label: string;
        hint: string;
        fields: {
          title: {
            label: string;
            hint: string;
          };
          name: {
            label: string;
            hint: string;
          };
        };
      };
      kinds: {
        label: string;
        hint: string;
        fields: {
          title: {
            label: string;
            hint: string;
          };
          name: {
            label: string;
            hint: string;
          };
        };
      };
    };
    alerts: {
      delete: string;
    };
    filters: {
      query: {
        label: string;
        kind: string;
      };
      time: {
        label: string;
        from: string;
        to: string;
      };
      level: {
        label: string;
      };
    };
  };
  members: {
    label: string;
    form: {
      actions: {
        create: string;
        edit: string;
      };
      name: string;
      email: string;
      role: string;
      tag: {
        label: string;
        hint: string;
      };
      revoked: string;
      password: string;
      confirm_password: string;
    };
    alerts: {
      delete: string;
    };
  };
  tags: {
    label: string;
    form: {
      actions: {
        create: string;
        edit: string;
      };
      name: {
        label: string;
        hint: string;
      };
      view_order: {
        label: string;
        hint: string;
      };
    };
    alerts: {
      delete: string;
    };
  };
  api_tokens: {
    label: string;
    form: {
      actions: {
        create: string;
        edit: string;
      };
      name: {
        label: string;
        hint: string;
      };
      revoked: string;
      creation_hint: string;
    };
    alerts: {
      delete: string;
    };
  };
  buttons: {
    create: string;
    settings: string;
    save: string;
    delete: string;
    edit: string;
    open: string;
    log_in: string;
    register: string;
    previous: string;
    next: string;
    clear: string;
  };
  tables: {
    log_groups: {
      time: string;
      level: string;
      text: string;
      request: string;
      response: string;
    };
    members: {
      email: string;
      name: string;
      role: string;
      revoked: string;
      actions: string;
    };
    tags: {
      name: string;
      view_order: string;
      actions: string;
    };
    api_tokens: {
      name: string;
      token: string;
      revoked: string;
      actions: string;
    };
  };
  auth: {
    email: string;
    password: string;
  };
  validations: {
    required: string;
    invalid_email: string;
    passwords_dont_match: string;
  };
  pagination: {
    first_page: string;
  };
  components: {
    search: {
      text: string;
    };
    sidebar: {
      open: string;
    };
  };
  miscellaneous: {
    loading: string;
    brand: "Moonlogs";
    blank_option: "—";
    not_found: string;
    forbidden: string;
    to_home: string;
  };
};

export type TranslationPath = Join<PathsToStringProps<Translation>>;
