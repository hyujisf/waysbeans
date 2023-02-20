/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
    "node_modules/flowbite-react/**/*.{js,jsx,ts,tsx}"
  ],
  theme: {
    extend: {
      colors: {
        coffee: {
          100: '#F6E6DA',
          200: '#DBB699',
          300: '#974A4A',
          400: '#613D2B',
        },
      }
    },
  },
  plugins: [
    require('@tailwindcss/line-clamp'),
    require('flowbite/plugin')
  ],
}
