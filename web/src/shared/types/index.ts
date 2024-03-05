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
    buttons: {
      create: string;
      settings: string;
      save: string;
      delete: string;
    };
    alerts: {
      delete: string;
    };
  };
};

export type TranslationPath = Join<PathsToStringProps<Translation>>;
