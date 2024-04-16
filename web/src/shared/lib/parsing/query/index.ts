import { isObjectPresent } from "../../presence";

export const queryStringToObject = (queryString: string | undefined): Record<string, string> => {
  queryString = queryString ?? "";

  const jsonObject: Record<string, string> = {};
  const parts = queryString.split("&");

  if (parts.length === 0) {
    return {};
  }

  for (const part of parts) {
    const [k, v] = part.split("=");
    if (!Boolean(k) || !Boolean(v)) {
      continue;
    }

    jsonObject[k] = v;
  }

  return jsonObject;
};

export const objectToQueryString = (object: Record<string, string>): string => {
  if (!isObjectPresent(object)) {
    return "";
  }

  return Object.keys(object)
    .map((key) => encodeURIComponent(key) + "=" + encodeURIComponent(object[key]))
    .join("&");
};
