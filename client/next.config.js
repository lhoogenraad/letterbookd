module.exports = {
	output: 'standalone', // Ensures the app is bundled properly for Netlify
	experimental: {
		appDir: true, // Use this if you're using Next.js 14+ and the new app directory
	},
	async rewrites() {
		return [
			{
				source: "/api/:path*",
				destination: "http://localhost:8080/api/:path*", // Matches backend
			},
		];
	},
};
