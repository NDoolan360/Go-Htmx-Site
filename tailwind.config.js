/** @type {import('tailwindcss').Config} */
export default ({
	content: ["./public/index.html", "./api/assets/template/*.html.tmpl"],
	darkMode: "class",
	theme: {
		extend: {
			colors: {
				text: "var(--text)",
				background: "var(--background)",
				primary: "var(--primary)",
				secondary: "var(--secondary)",
				accent: "var(--accent)",
			},
			fontSize: {
				sm: "0.750rem",
				base: "1rem",
				xl: "1.333rem",
				"2xl": "1.777rem",
				"3xl": "2.369rem",
				"4xl": "3.158rem",
				"5xl": "4.210rem",
			},
			fontFamily: {
				heading: "Fragment Mono, mono",
				body: "Josefin Sans, sans-serif"
			},
			fontWeight: {
				normal: "400",
				bold: "700",
			},
		},
	},
	future: {
		hoverOnlyWhenSupported: true,
	},
});
