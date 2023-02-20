module.exports = {
  plugins: {
    cssnano: {
      preset: ["advanced", {
        discardComments: {
          removeAll: true,
        },
      }]
    },
    'postcss-import': {},
    tailwindcss: {},
    autoprefixer: {},
  },
}
