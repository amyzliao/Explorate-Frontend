/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: 'jit',
  content: ["./src/**/*.{html,js,vue}",
            "./public/**/*.html",
            "./node_modules/flowbite/**/*.js"],
  theme: {
    extend: {
      colors: {
        primary: '#ef5350',
        primary_hover: '#bc413e',
        secondary_yellow: '#e8b974',
        secondary_blue: '#37bce5',
        secondary_gray: '#b8aaa0'
      },
    },
    fontFamily: {
      sans: ['Optima', 'sans-serif'],
      serif: ['Merriweather', 'serif'],
    },
  },
  plugins: [
    require('flowbite/plugin')
  ],
}

