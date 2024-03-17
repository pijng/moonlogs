/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{html,ts}"],
  theme: {
    extend: {
      colors: {
        eigengrau: "#1b1b1f",
        "raisin-black": "#202127",
        "squid-ink": "#32363f",
        "shadow-gray": "#3c3f44",
        "slate-gray": "#65758529",
      },
      maxWidth: {
        5: "5rem",
      },
    },
  },
  darkMode: "class",
};
