const path = require('path');

module.exports = {
  entry: './web/index.js',
  output: {
    clean: true,
    filename: 'js-bin/app.js',
    path: path.resolve(__dirname, 'out'),
  },
  module: {
    rules: [
      {
        test: /\.(s?css)$/,
        use: [{
          loader: 'style-loader'
        }, {
          loader: 'css-loader'
        }]
      },
    ],
  },
};
