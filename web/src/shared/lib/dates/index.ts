export const DATEFORMAT_OPTIONS: Intl.DateTimeFormatOptions = {
  hour: "numeric",
  minute: "numeric",
  second: "numeric",
  year: "numeric",
  month: "numeric",
  day: "numeric",
};

export const TIMEZONE = Intl.DateTimeFormat().resolvedOptions().timeZone;
