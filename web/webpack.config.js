const path = require('path');

module.exports = {
	entry: './assets/js/main.ts',
	output: {
		filename: 'bundle.js',
		path: path.resolve(__dirname, 'static/js'),
	},
	mode: 'development',
	devServer: {
		static: path.resolve(__dirname, 'static'),
		port: 3000,
		hot: true,
	},
	resolve: {
		extensions: ['.ts', '.js'],
	},
	module: {
		rules: [
			{
				test: /\.ts$/,
				use: 'ts-loader',
				exclude: /node_modules/,
			},
		],
	},
};
