/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'bg-dark': '#141414',
        'panel-dark': '#1e1f23',
        'sub-dark':'#2A2A36',
        'alt-dark': '#333a4b',
        'border-dark':'#404040',
        'text-gray':'#87888b',
        'text-gray-1':'#9ea2ab',
        'btn-primary':'#6577f4',
        'btn-primary-hover':'#5666D1',
      }
    },
  },
  plugins: [],
}

