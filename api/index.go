package api

import (
	"html/template"
	"net/http"

	"github.com/NDoolan360/go-htmx-site/utils"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		utils.GetApiResource("template/index.gohtml"),
		utils.GetApiResource("template/head.gohtml"),
	))

	err := tmpl.Execute(w, utils.IndexTemplate{
		Title:       "Nathan Doolan",
		Description: "A personal website showcasing Nathan Doolan's journey as a full-time software engineer in Melbourne. Explore his professional experience, projects, and interests in technology, board games, and 3D printing.",
		InternalLinks: []utils.Link{
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
		ThemeSwitchSVG: utils.GetSVGLogo("theme_switch"),
		ExternalLinks: []utils.Link{
			{
				Label: "Github",
				URL:   "https://github.com/NDoolan360",
				Logo:  utils.GetSVGLogo("github"),
			},
			{
				Label: "LinkedIn",
				URL:   "https://www.linkedin.com/in/nathan-doolan-835a13171",
				Logo:  utils.GetSVGLogo("linkedin"),
			},
			{
				Label: "Discord",
				URL:   "https://discord.com/users/nothindoin",
				Logo:  utils.GetSVGLogo("discord"),
			},
			{
				Label: "Cults3D",
				URL:   "https://cults3d.com/en/users/ND360",
				Logo:  utils.GetSVGLogo("cults3d"),
			},
			{
				Label: "Boardgame Geek",
				URL:   "https://boardgamegeek.com/user/Nothin_Doin",
				Logo:  utils.GetSVGLogo("bgg"),
			},
		},
		Profile: utils.Profile{
			ImageAttr: []template.HTMLAttr{
				template.HTMLAttr(`title="Picture of me at the Strawberry Field in Liverpool"`),
				template.HTMLAttr(`srcset="/images/profile-192.webp, /images/profile-512.webp 2.66x, /images/profile-792.webp 4.125x"`),
			},
			Paragraphs: []string{
				"Hey there, I'm Nathan Doolan, a full-time software engineer based in Melbourne, Australia. I spend my days coding and my spare time collecting board games, enjoying nature, and experimenting with my 3D printer.",
				"With a background in computer science, I'm always eager to learn new things, tackle new challenges, and find satisfaction in the simple joy of crafting an elegant solution to a problem.",
			},
		},
		Experiences: []utils.Experience{
			{
				Date: utils.Date{Start: "Jan 2024", End: "Present"},
				Workplace: utils.Link{
					Label: "Kaluza",
					URL:   "https://kaluza.com/",
					Logo:  utils.GetSVGLogo("kaluza"),
				},
				Positions: []utils.Position{{Title: "Software Engineer", Current: true}},
				Topics:    []string{"Typescript", "Git", "Github Actions"},
			},
			{
				Date: utils.Date{Start: "Jul 2021", End: "Dec 2023"},
				Workplace: utils.Link{
					Label: "Gentrack",
					URL:   "https://gentrack.com/",
					Logo:  utils.GetSVGLogo("gentrack"),
				},
				Positions: []utils.Position{
					{Title: "Intermediate Software Engineer"},
					{Title: "Junior Software Engineer"},
					{Title: "Graduate Software Engineer"},
				},
				Topics: []string{"Git", "SQL", "Github Actions", "Docker", "Jenkins", "API Design", "Unit Testing"},
			},
			{
				Date: utils.Date{Start: "Feb 2018", End: "Jul 2021"},
				Workplace: utils.Link{
					Label: "Proquip Rental & Sales",
					URL:   "https://pqrs.com.au/",
					Logo:  utils.GetSVGLogo("proquip"),
				},
				Positions: []utils.Position{
					{Title: "IT Support Specialist"},
					{Title: "IT/Marketing Assistant"},
					{Title: "Administrative Assistant"},
				},
				Topics: []string{"IT Support", "Adobe Suite", "Social Media Marketing", "Wordpress", "Google Analytics"},
			},
			{
				Date: utils.Date{Start: "Feb 2018", End: "Feb 2021"},
				Workplace: utils.Link{
					Label: "University of Melbourne",
					URL:   "https://www.unimelb.edu.au/",
					Logo:  utils.GetSVGLogo("melbourneuniversity"),
				},
				Positions: []utils.Position{{Title: "Bachelor of Science: Computing and Software Systems"}},
				Education: true,
			},
		},
		ToolSections: []utils.ToolSection{
			{
				Title: "Built with:",
				Links: []utils.Link{
					{
						Label: "htmx",
						URL:   "https://htmx.org",
						Logo:  utils.GetSVGLogo("htmx"),
					},
					{
						Label: "hyperscript",
						URL:   "https://hyperscript.org",
						Logo:  utils.GetSVGLogo("hyperscript"),
					},
					{
						Label: "Tailwind CSS",
						URL:   "https://tailwindcss.com",
						Logo:  utils.GetSVGLogo("tailwind"),
					},
				},
			},
			{
				Title: "Deployed with:",
				Links: []utils.Link{
					{
						Label: "Vercel",
						URL:   "https://vercel.com",
						Logo:  utils.GetSVGLogo("vercel"),
					},
				},
			},
		},
		Copyright: utils.Copyright("Nathan Doolan"),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
