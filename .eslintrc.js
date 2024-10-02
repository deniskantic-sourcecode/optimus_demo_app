/**
 * .eslint.js
 *
 * ESLint configuration file.
 */

module.exports = {
  root: true,
  env: {
    node: true,
    browser: true,
    es2021: true,
  },
  extends: [
    'vuetify',
    '@vue/eslint-config-typescript',
    './.eslintrc-auto-import.json',
    'eslint:recommended',
    'plugin:vue/recommended', // Include Vue.js rules
  ],
  rules: {
    'vue/multi-word-component-names': 'off',
    'no-trailing-spaces': 'error', // Enable the trailing spaces rule

  },
}
