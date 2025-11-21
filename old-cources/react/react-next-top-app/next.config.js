/** @type {import('next').NextConfig} */
const nextConfig = {
	images: {
		domains: ['courses-top.ru', 'cdn-bucket.hb.bizmrg.com']
	},
	webpack(config, options) {
		config.module.rules.push({
			loader: '@svgr/webpack',
			options: {
				prettier: false,
				issuer: /\.[jt]sx?$/,
				svgo: true,
				svgoConfig: {
					plugins: [
						{
							name: 'preset-default',
							params: {
								overrides: {
									// disable plugins
									removeViewBox: false,
								},
							},
						}
					],
				},
				titleProp: true,
			},
			test: /\.svg$/,
		});

		return config;
	},
	reactStrictMode: true,
};

module.exports = nextConfig;
