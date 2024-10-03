module.exports = {
  content: [
    './src/**/*.{js,ts,jsx,tsx}', // Include all JS, TS, JSX, and TSX files in the src directory
    './public/index.html', // Include the index.html file in the public directory
  ],
  theme: {
    extend: {
      colors: {
        placeholder: '#A6A39F',
      },
    },
  },
  variants: {
    extend: {
      placeholderColor: ['focus'],
    },
  },
  plugins: [],
};