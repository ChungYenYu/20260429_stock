/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'apple-blue': '#0071e3',
        'apple-gray': '#f5f5f7',
        'near-black': '#1d1d1f',
      },
      fontFamily: {
        'sf-pro': ['system-ui', '-apple-system', 'sans-serif'],
      },
      borderRadius: {
        'apple-sm': '5px',
        'apple-md': '8px',
        'apple-lg': '12px',
        'apple-pill': '980px',
      },
      spacing: {
        'apple-base': '8px',
      }
    },
  },
  plugins: [],
}
