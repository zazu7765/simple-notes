/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: { backgroundColor: ['active'],},
  },
  plugins: [require('flowbite/plugin')],
  darkMode: 'class',
}
