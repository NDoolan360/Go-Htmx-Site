package api

import (
	"fmt"
	"html/template"
	"net/http"

	utils "github.com/NDoolan360/go-htmx-site/src"
)

type Index struct {
	Title          string
	Description    string
	InternalLinks  []Link
	ExternalLinks  []Link
	ThemeSwitchSVG template.HTML
	Profile
	ToolSections  []ToolSection
	CopyrightYear string
}

type Link struct {
	Label string
	URL   template.URL
	Logo  template.HTML
}

type Profile struct {
	ImageAttr  []template.HTMLAttr
	Paragraphs []string
}

type ToolSection struct {
	Title string
	Links []Link
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles(
		utils.GetTemplatePath("index.gohtml"),
		utils.GetTemplatePath("head.gohtml"),
	)).Execute(w, Index{
		Title:       "Nathan Doolan",
		Description: "A personal website showcasing Nathan Doolan's journey as a full-time software engineer in Melbourne. Explore his professional experience, projects, and interests in technology, board games, and 3D printing.",
		InternalLinks: []Link{
			{
				Label: "About",
				URL:   template.URL("#about"),
			},
			{
				Label: "Experience",
				URL:   template.URL("#experience"),
			},
			{
				Label: "Projects",
				URL:   template.URL("#projects"),
			},
		},
		ThemeSwitchSVG: utils.IgnoreErr(utils.GetSVGLogo("theme_switch")),
		ExternalLinks: []Link{
			{
				Label: "Github",
				URL:   template.URL("https://github.com/NDoolan360"),
				Logo:  utils.IgnoreErr(utils.GetSVGLogo("github")),
			},
			{
				Label: "LinkedIn",
				URL:   template.URL("https://www.linkedin.com/in/nathan-doolan-835a13171"),
				Logo:  utils.IgnoreErr(utils.GetSVGLogo("linkedin")),
			},
			{
				Label: "Discord",
				URL:   template.URL("https://discord.com/users/nothindoin"),
				Logo:  utils.IgnoreErr(utils.GetSVGLogo("discord")),
			},
			{
				Label: "Cults3D",
				URL:   template.URL("https://cults3d.com/en/users/ND360"),
				Logo:  utils.IgnoreErr(utils.GetSVGLogo("cults3d")),
			},
			{
				Label: "Boardgame Geek",
				URL:   template.URL("https://boardgamegeek.com/user/Nothin_Doin"),
				Logo:  utils.IgnoreErr(utils.GetSVGLogo("bgg")),
			},
		},
		Profile: Profile{
			ImageAttr: []template.HTMLAttr{
				template.HTMLAttr(`title="Picture of me at the Strawberry Field in Liverpool"`),
				template.HTMLAttr(`srcset="/images/profile-192.webp, /images/profile-512.webp 2.66x, /images/profile-792.webp 4.125x"`),
			},
			Paragraphs: []string{
				"Hey there, I'm Nathan Doolan, a full-time software engineer based in Melbourne, Australia. I spend my days coding and my spare time collecting board games, enjoying nature, and experimenting with my 3D printer.",
				"With a background in computer science, I'm always eager to learn new things, tackle new challenges, and find satisfaction in the simple joy of crafting an elegant solution to a problem.",
			},
		},
		ToolSections: []ToolSection{
			{
				Title: "Built with:",
				Links: []Link{
					{
						Label: "htmx",
						URL:   template.URL("https://htmx.org"),
						Logo:  utils.IgnoreErr(utils.GetSVGLogo("htmx")),
					},
					{
						Label: "hyperscript",
						URL:   template.URL("https://hyperscript.org"),
						Logo:  utils.IgnoreErr(utils.GetSVGLogo("hyperscript")),
					},
					{
						Label: "Tailwind CSS",
						URL:   template.URL("https://tailwindcss.com"),
						Logo:  utils.IgnoreErr(utils.GetSVGLogo("tailwind")),
					},
				},
			},
			{
				Title: "Deployed with:",
				Links: []Link{
					{
						Label: "Vercel",
						URL:   template.URL("https://vercel.com"),
						Logo:  utils.IgnoreErr(utils.GetSVGLogo("vercel")),
					},
				},
			},
		},
		CopyrightYear: fmt.Sprint(Now().Year()),
	})
}
