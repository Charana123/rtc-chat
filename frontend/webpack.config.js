const path = require('path');
const MiniCssExtractPlugin = require("mini-css-extract-plugin");

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
            use: ['babel-loader']
        },
        // CSS Files
        {
            test: /\.css$/,
            use: [
                MiniCssExtractPlugin.loader, // create seperate .css file
                // 'style-loader', // embed into .js as base64 encoded string
                'css-loader'
            ]
        },
        { 
            test: /\.woff(2)?(\?v=[0-9]\.[0-9]\.[0-9])?$/, 
            loader: "url-loader?limit=10000&mimetype=application/font-woff" 
        },
        { 
            test: /\.(ttf|eot|svg)(\?v=[0-9]\.[0-9]\.[0-9])?$/, 
            loader: "file-loader" 
        },
        ]
    },

    // Plugins
    plugins: [
        new MiniCssExtractPlugin({
            filename: "style.css"
        })
    ]
};