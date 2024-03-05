import { Store } from "effector";

const VARIABLE_PATTERN = /\{([^}]+)\}/;
const REPLACE_PATTERN = /\{(.+?)\}/g;

export const invokeTemplate = (template: Store<string>, vars: Record<string, any> | undefined): Store<string> => {
  const $parts = template.map((t) => t.split("\n"));

  const $invokedParts = $parts.map((parts) => {
    return parts.map((part) => {
      const matches = part.match(VARIABLE_PATTERN);
      const variablePlaceholder = matches && matches[1];

      if (variablePlaceholder && !vars?.[variablePlaceholder]) {
        return null;
      }

      return part.replace(REPLACE_PATTERN, (_, name) => {
        return vars?.[name];
      });
    });
  });

  return $invokedParts.map((parts) => parts.filter((p) => p !== null).join("\n"));
};
