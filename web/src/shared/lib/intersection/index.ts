export const intersection = (arrays: string[][]): string[] => {
  if (arrays.length === 0) return [];

  return arrays.reduce((prevArray, currentArray) => {
    return prevArray.filter((value) => currentArray.includes(value));
  });
};
