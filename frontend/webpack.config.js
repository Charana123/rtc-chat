const path = require('path');

module.exports = {
    // Entry
    entry: './src/index.jsx',
    mode: 'development',

    // Output
    output: {
        path: path.resolve(__dirname, 'public/dist'),
        filename: 'bundle.js'
    },

    // Loaders
    module: {
        rules : [
        // JavaScript/JSX Files
        {
            test: /\.jsx$/,
            exclude: /node_modules/,
            use: ['babel-loader'],
        },
        // // CSS Files
        // {
        //     test: /\.css$/,
        //     use: ['style-loader', 'css-loader'],
        // }
        ]
    },

    // Plugins
    plugins: [],
};