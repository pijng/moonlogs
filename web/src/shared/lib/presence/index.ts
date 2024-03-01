export const isObjectPresent = (object: Record<any, any>): boolean => {
  return Boolean(object) && Object.keys(object).length > 0;
};
