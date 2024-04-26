package api

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/NDoolan360/go-htmx-site/logos"
)

// GetIndex handles the request for rendering the index page.
func GetIndex(w http.ResponseWriter, r *http.Request) {
	indexTemplate := template.Must(template.ParseFiles(
		"templates/index.html.tmpl",
		"templates/head.html.tmpl",
		"templates/theme-switch.html.tmpl",
	))

	err := indexTemplate.Execute(w, IndexTemplate{
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
		ExternalLinks: []Link{
			{
				Label: "Github",
				URL:   "https://github.com/NDoolan360",
				Logo:  logos.GetSVGLogo("github"),
			},
			{
				Label: "LinkedIn",
				URL:   "https://www.linkedin.com/in/nathan-doolan-835a13171",
				Logo:  logos.GetSVGLogo("linkedin"),
			},
			{
				Label: "Discord",
				URL:   "https://discord.com/users/nothindoin",
				Logo:  logos.GetSVGLogo("discord"),
			},
			{
				Label: "Cults3D",
				URL:   "https://cults3d.com/en/users/ND360",
				Logo:  logos.GetSVGLogo("cults3d"),
			},
			{
				Label: "Boardgame Geek",
				URL:   "https://boardgamegeek.com/user/Nothin_Doin",
				Logo:  logos.GetSVGLogo("bgg"),
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
				Date: Date{Start: "Jan 2024", End: "Present"},
				Link: Link{
					Label: "Kaluza",
					URL:   "https://kaluza.com/",
					Logo:  logos.GetSVGLogo("kaluza"),
				},
				Positions: []Position{{Title: "Software Engineer", Current: true}},
				Knowledge: []Link{{Label: "Typescript", URL: "https://www.typescriptlang.org"}, {Label: "Github Actions", URL: "https://github.com/features/actions"}, {Label: "CircleCI", URL: "https://circleci.com"}, {Label: "NestJS", URL: "https://nestjs.com"}, {Label: "Terraform", URL: "https://www.terraform.io"}, {Label: "Kafka", URL: "https://kafka.apache.org"}, {Label: "DataDog", URL: "https://www.datadoghq.com"}},
			},
			{
				Date: Date{Start: "Jul 2021", End: "Dec 2023"},
				Link: Link{
					Label: "Gentrack",
					URL:   "https://gentrack.com/",
					Logo:  logos.GetSVGLogo("gentrack"),
				},
				Positions: []Position{
					{Title: "Intermediate Software Engineer"},
					{Title: "Junior Software Engineer"},
					{Title: "Graduate Software Engineer"},
				},
				Knowledge: []Link{{Label: "SQL"}, {Label: "API Design"}, {Label: "Unit Testing"}, {Label: "Docker", URL: "https://www.docker.com"}, {Label: "Jenkins", URL: "https://www.jenkins.io"}},
			},
			{
				Date: Date{Start: "Feb 2018", End: "Jul 2021"},
				Link: Link{
					Label: "Proquip Rental & Sales",
					URL:   "https://pqrs.com.au/",
					Logo:  logos.GetSVGLogo("proquip"),
				},
				Positions: []Position{
					{Title: "IT Support Specialist"},
					{Title: "IT/Marketing Assistant"},
					{Title: "Administrative Assistant"},
				},
				Knowledge: []Link{{Label: "IT Support"}, {Label: "Social Media Marketing"}, {Label: "Adobe Suite", URL: "https://www.adobe.com/products/catalog.html"}, {Label: "Wordpress", URL: "https://wordpress.com"}, {Label: "Google Analytics", URL: "https://analytics.google.com/analytics"}},
			},
			{
				Date: Date{Start: "Feb 2018", End: "Feb 2021"},
				Link: Link{
					Label: "University of Melbourne",
					URL:   "https://www.unimelb.edu.au/",
					Logo:  logos.GetSVGLogo("melbourneuniversity"),
				},
				Positions: []Position{{Title: "Bachelor of Science: Computing and Software Systems"}},
				Knowledge: []Link{{Label: "Course Overview", URL: "https://study.unimelb.edu.au/find/courses/major/computing-and-software-systems"}},
				Education: true,
			},
		},
		ToolSections: []ToolSection{
			{
				Title: "Built with:",
				Links: []Link{
					{
						Label: "Go",
						URL:   "https://go.dev/",
						Logo:  logos.GetSVGLogo("go"),
					},
					{
						Label: "htmx",
						URL:   "https://htmx.org",
						Logo:  logos.GetSVGLogo("htmx"),
					},
					{
						Label: "Tailwind CSS",
						URL:   "https://tailwindcss.com",
						Logo:  logos.GetSVGLogo("tailwind"),
					},
				},
			},
			{
				Title: "Deployed with:",
				Links: []Link{
					{
						Label: "Netlify",
						URL:   "https://netlify.com",
						Logo:  logos.GetSVGLogo("netlify"),
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

// Copyright generates a copyright string with the current year with a given name.
func Copyright(name string) string {
	year := Now().Year()
	return fmt.Sprintf("© %s %d", name, year)
}
