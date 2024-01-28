export const hashCode = (s: string): string => {
  let h: number = 0;
  for (let i = 0; i < s.length; i++) h = (Math.imul(31, h) + s.charCodeAt(i)) | 0;

  return h.toString();
};
