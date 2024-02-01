/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{html,ts}"],
  theme: {
    extend: {
      maxWidth: {
        5: "5rem",
      },
    },
  },
  darkMode: "class",
};
