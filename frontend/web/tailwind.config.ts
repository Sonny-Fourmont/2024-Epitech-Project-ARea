module.exports = {
  content: [
    './src/**/*.{js,ts,jsx,tsx}', // Include all JS, TS, JSX, and TSX files in the src directory
    './public/index.html', // Include the index.html file in the public directory
  ],
  theme: {
    extend: {
      colors: {
        placeholder: '#A6A39F',
        buttonColor: '#3b82f6',
        buttonHoverColor: '#2563eb',
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