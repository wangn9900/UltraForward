/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: {
          DEFAULT: '#00f2ff',
          300: '#66f7ff',
        },
        accent: '#7000ff',
      }
    },
  },
  plugins: [],
}
