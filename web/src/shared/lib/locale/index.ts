export const getLocale = () => {
  if (navigator.languages && navigator.languages.length) {
    return navigator.languages[0];
  }
  return navigator.language;
};
