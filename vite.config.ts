import autoprefixer from "autoprefixer";
import tailwindcss from "tailwindcss";
import { defineConfig } from "vite";

export default defineConfig({
	css: {
		postcss: {
			plugins: [autoprefixer(), tailwindcss()],
		},
	},
	server: {
		proxy: {
			"/proxy/api/github": {
				target: "https://api.github.com/users/NDoolan360/repos",
				changeOrigin: true,
				rewrite: (path) => path.replace(/^\/proxy\/api\/github/, ""),
			},
			"/proxy/cults3d": {
				target: "https://cults3d.com/en/users/ND360/3d-models",
				changeOrigin: true,
				rewrite: (path) => path.replace(/^\/proxy\/cults3d/, ""),
			},
			"/proxy/boardgamegeek": {
				target:
					"https://boardgamegeek.com/geeksearch.php?action=search&advsearch=1&objecttype=boardgame&include%5Bdesignerid%5D=133893",
				changeOrigin: true,
				rewrite: (path) => path.replace(/^\/proxy\/boardgamegeek/, ""),
			},
			"/proxy/xmlapi/boardgamegeek/": {
				target: "https://api.geekdo.com/xmlapi/boardgame/",
				changeOrigin: true,
				rewrite: (path) => path.replace(/^\/proxy\/xmlapi\/boardgamegeek\//, ""),
			},
		},
	},
});
