package api

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

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
		FromPWD("api/template/index.gohtml"),
		FromPWD("api/template/head.gohtml"),
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
		ThemeSwitchSVG: IgnoreErr(GetSVGLogo("theme_switch")),
		ExternalLinks: []Link{
			{
				"Github",
				"https://github.com/NDoolan360",
				IgnoreErr(GetSVGLogo("github")),
			},
			{
				"LinkedIn",
				"https://www.linkedin.com/in/nathan-doolan-835a13171",
				IgnoreErr(GetSVGLogo("linkedin")),
			},
			{
				"Discord",
				"https://discord.com/users/nothindoin",
				IgnoreErr(GetSVGLogo("discord")),
			},
			{
				"Cults3D",
				"https://cults3d.com/en/users/ND360",
				IgnoreErr(GetSVGLogo("cults3d")),
			},
			{
				"Boardgame Geek",
				"https://boardgamegeek.com/user/Nothin_Doin",
				IgnoreErr(GetSVGLogo("bgg")),
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
					IgnoreErr(GetSVGLogo("kaluza")),
				},
				Positions: []Position{{"Software Engineer", true}},
				Topics:    []string{"Typescript", "Git", "Github Actions"},
			},
			{
				Date: Date{"Jul 2021", "Dec 2023"},
				Workplace: Link{
					"Gentrack",
					"https://gentrack.com/",
					IgnoreErr(GetSVGLogo("gentrack")),
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
					IgnoreErr(GetSVGLogo("proquip")),
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
					IgnoreErr(GetSVGLogo("melbourneuniversity")),
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
						IgnoreErr(GetSVGLogo("htmx")),
					},
					{
						"hyperscript",
						"https://hyperscript.org",
						IgnoreErr(GetSVGLogo("hyperscript")),
					},
					{
						"Tailwind CSS",
						"https://tailwindcss.com",
						IgnoreErr(GetSVGLogo("tailwind")),
					},
				},
			},
			{
				"Deployed with:",
				[]Link{
					{
						"Vercel",
						"https://vercel.com",
						IgnoreErr(GetSVGLogo("vercel")),
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

var Now = func() time.Time {
	return time.Now()
}

func Copyright(name string) string {
	year := Now().Year()
	return fmt.Sprintf("© %s %d", name, year)
}
