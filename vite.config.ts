import autoprefixer from "autoprefixer";
import tailwindcss from "tailwindcss";
import { defineConfig } from "vite";
import { VitePWA } from "vite-plugin-pwa";
import sitemap from "vite-plugin-sitemap";

export default defineConfig({
	css: {
		postcss: {
			plugins: [autoprefixer(), tailwindcss()],
		},
	},
	plugins: [
		sitemap(),
		VitePWA({
			registerType: "prompt",
			includeAssets: ["/favicon.ico", "images/icons/apple-touch-icon.png"],
			manifest: {
				name: "Nathan Doolan",
				description:
					"A personal website showcasing Nathan Doolan's journey as a full-time software engineer in Melbourne. Explore his professional experience, projects, and interests in technology, board games, and 3D printing.",
				// biome-ignore lint/style/useNamingConvention: External API name
				short_name: "/ND",
				// biome-ignore lint/style/useNamingConvention: External API name
				theme_color: "#283c37",
				// biome-ignore lint/style/useNamingConvention: External API name
				background_color: "#192926",
				icons: [
					{
						src: "images/icons/pwa-192x192.png",
						sizes: "192x192",
						type: "image/png",
					},
					{
						src: "images/icons/pwa-512x512.png",
						sizes: "512x512",
						type: "image/png",
					},
					{
						src: "images/icons/pwa-512x512.png",
						sizes: "512x512",
						type: "image/png",
						purpose: "any maskable",
					},
				],
			},
		}),
	],
});
