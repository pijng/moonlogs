module.exports = {
  root: true,
  env: {
    node: true,
  },
  parser: "@typescript-eslint/parser",
  // eslint-disable-next-line prettier/prettier
  extends: [
    "plugin:@typescript-eslint/recommended",
    "plugin:prettier/recommended",
  ],
  rules: {
    "@typescript-eslint/no-explicit-any": "off",
    "@typescript-eslint/no-non-null-assertion": "off",
    "no-debugger": process.env.NODE_ENV === "production" ? "error" : "warn",
    "prettier/prettier": [
      "error",
      {
        printWidth: 120,
        endOfLine: "auto",
      },
    ],
  },
  ignorePatterns: ["./build", "./dist"],
  globals: {},
  parserOptions: {
    ecmaVersion: "latest",
    sourceType: "module",
  },
};
