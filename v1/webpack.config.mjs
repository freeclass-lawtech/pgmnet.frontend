import { dirname, resolve } from 'path';
import { fileURLToPath } from 'url';
import TerserPlugin from 'terser-webpack-plugin';

// Get the current directory equivalent to __dirname
const __dirname = dirname(fileURLToPath(import.meta.url));

export default {
    entry: './source/index.js',
    output: {
        filename: 'main.js',
        path: resolve(__dirname, 'dist'),
    },
    module: {
        rules: [
            {
                test: /\.css$/i,
                use: ['style-loader', 'css-loader'],
            },
        ],
    },
    optimization: {
        minimize: true, // Enable minimization
        minimizer: [
            new TerserPlugin({
                terserOptions: {
                    format: {
                        comments: false, // Remove comments
                    },
                    compress: {
                        drop_console: true, // Remove console.log statements
                    },
                },
                extractComments: false, // Prevent comments from being extracted into a separate file
            }),
        ],
    },
};
