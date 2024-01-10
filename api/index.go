package api

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	utils "github.com/NDoolan360/go-htmx-site/src"
)

var Now = func() time.Time {
	return time.Now()
}

type Index struct {
	Title          string
	Description    string
	InternalLinks  []Link
	ExternalLinks  []Link
	ThemeSwitchSVG template.HTML
	Profile
	Experiences  []Experience
	ToolSections []ToolSection
	Copyright    string
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

type Experience struct {
	Date
	Workplace Link
	Positions []Position
	Topics    []string
	Education bool
}

type Date struct {
	Start string
	End   string
}

type Position struct {
	Title   string
	Current bool
}

type ToolSection struct {
	Title string
	Links []Link
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles(
		utils.GetTemplatePath("index.gohtml"),
		utils.GetTemplatePath("head.gohtml"),
	))

	err := tmp.Execute(w, Index{
		Title:       "Nathan Doolan",
		Description: "A personal website showcasing Nathan Doolan's journey as a full-time software engineer in Melbourne. Explore his professional experience, projects, and interests in technology, board games, and 3D printing.",
		InternalLinks: []Link{
			{
				Label: "About",
				URL:   "#about",
			},
			{
				Label: "Experience",
				URL:   "#experience",
			},
			{
				Label: "Projects",
				URL:   "#projects",
			},
		},
		ThemeSwitchSVG: utils.IgnoreErr(utils.GetSVGLogo("theme_switch")),
		ExternalLinks: []Link{
			{
				"Github",
				"https://github.com/NDoolan360",
				utils.IgnoreErr(utils.GetSVGLogo("github")),
			},
			{
				"LinkedIn",
				"https://www.linkedin.com/in/nathan-doolan-835a13171",
				utils.IgnoreErr(utils.GetSVGLogo("linkedin")),
			},
			{
				"Discord",
				"https://discord.com/users/nothindoin",
				utils.IgnoreErr(utils.GetSVGLogo("discord")),
			},
			{
				"Cults3D",
				"https://cults3d.com/en/users/ND360",
				utils.IgnoreErr(utils.GetSVGLogo("cults3d")),
			},
			{
				"Boardgame Geek",
				"https://boardgamegeek.com/user/Nothin_Doin",
				utils.IgnoreErr(utils.GetSVGLogo("bgg")),
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
		Experiences: []Experience{
			{
				Date: Date{"Jan 2024", "Present"},
				Workplace: Link{
					"Kaluza",
					"https://kaluza.com/",
					utils.IgnoreErr(utils.GetSVGLogo("kaluza")),
				},
				Positions: []Position{{"Software Engineer", true}},
				Topics:    []string{"Typescript", "Git", "Github Actions"},
			},
			{
				Date: Date{"Jul 2021", "Dec 2023"},
				Workplace: Link{
					"Gentrack",
					"https://gentrack.com/",
					utils.IgnoreErr(utils.GetSVGLogo("gentrack")),
				},
				Positions: []Position{
					{"Intermediate Software Engineer", false},
					{"Junior Software Engineer", false},
					{"Graduate Software Engineer", false},
				},
				Topics: []string{"Git", "SQL", "Github Actions", "Docker", "Jenkins", "API Design", "Unit Testing"},
			},
			{
				Date: Date{"Feb 2018", "Jul 2021"},
				Workplace: Link{
					"Proquip Rental & Sales",
					"https://pqrs.com.au/",
					utils.IgnoreErr(utils.GetSVGLogo("proquip")),
				},
				Positions: []Position{
					{"IT Support Specialist", false},
					{"IT/Marketing Assistant", false},
					{"Administrative Assistant", false},
				},
				Topics: []string{"IT Support", "Adobe Suite", "Social Media Marketing", "Wordpress", "Google Analytics"},
			},
			{
				Date: Date{"Feb 2018", "Feb 2021"},
				Workplace: Link{
					"University of Melbourne",
					"https://www.unimelb.edu.au/",
					utils.IgnoreErr(utils.GetSVGLogo("melbourneuniversity")),
				},
				Positions: []Position{{"Bachelor of Science: Computing and Software Systems", false}},
				Education: true,
			},
		},
		ToolSections: []ToolSection{
			{
				"Built with:",
				[]Link{
					{
						"htmx",
						"https://htmx.org",
						utils.IgnoreErr(utils.GetSVGLogo("htmx")),
					},
					{
						"hyperscript",
						"https://hyperscript.org",
						utils.IgnoreErr(utils.GetSVGLogo("hyperscript")),
					},
					{
						"Tailwind CSS",
						"https://tailwindcss.com",
						utils.IgnoreErr(utils.GetSVGLogo("tailwind")),
					},
				},
			},
			{
				"Deployed with:",
				[]Link{
					{
						"Vercel",
						"https://vercel.com",
						utils.IgnoreErr(utils.GetSVGLogo("vercel")),
					},
				},
			},
		},
		Copyright: Copyright("Nathan Doolan"),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Copyright(name string) string {
	year := Now().Year()
	return fmt.Sprintf("© %s %d", name, year)
}
