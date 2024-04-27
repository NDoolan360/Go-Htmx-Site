package handlers

import (
	"html/template"
	"net/http"

	"github.com/NDoolan360/go-htmx-site/internal/components"
	"github.com/NDoolan360/go-htmx-site/internal/layouts"
)

// Index handles the request for rendering the index page.
func Index(w http.ResponseWriter, r *http.Request) {
	layouts.Index.Execute(w, layouts.IndexTemplate{
		Title:       "Nathan Doolan",
		Description: "A personal website showcasing Nathan Doolan's journey as a full-time software engineer in Melbourne. Explore his professional experience, projects, and interests in technology, board games, and 3D printing.",
		InternalLinks: []components.Link{
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
		ExternalLinks: []components.Link{
			{
				Label: "Github",
				URL:   "https://github.com/NDoolan360",
				Logo:  components.GetSVGLogo("github"),
			},
			{
				Label: "LinkedIn",
				URL:   "https://www.linkedin.com/in/nathan-doolan-835a13171",
				Logo:  components.GetSVGLogo("linkedin"),
			},
			{
				Label: "Discord",
				URL:   "https://discord.com/users/nothindoin",
				Logo:  components.GetSVGLogo("discord"),
			},
			{
				Label: "Cults3D",
				URL:   "https://cults3d.com/en/users/ND360",
				Logo:  components.GetSVGLogo("cults3d"),
			},
			{
				Label: "Boardgame Geek",
				URL:   "https://boardgamegeek.com/user/Nothin_Doin",
				Logo:  components.GetSVGLogo("bgg"),
			},
		},
		Profile: components.Profile{
			ImageAttr: []template.HTMLAttr{
				template.HTMLAttr(`title="Picture of me at the Strawberry Field in Liverpool"`),
				template.HTMLAttr(`srcset="/images/profile-192.webp, /images/profile-512.webp 2.66x, /images/profile-792.webp 4.125x"`),
			},
			Paragraphs: []string{
				"Hey there, I'm Nathan Doolan, a full-time software engineer based in Melbourne, Australia. I spend my days coding and my spare time collecting board games, enjoying nature, and experimenting with my 3D printer.",
				"With a background in computer science, I'm always eager to learn new things, tackle new challenges, and find satisfaction in the simple joy of crafting an elegant solution to a problem.",
			},
		},
		Experiences: []components.Experience{
			{
				Date: components.Date{Start: "Jan 2024", End: "Present"},
				Link: components.Link{
					Label: "Kaluza",
					URL:   "https://kaluza.com/",
					Logo:  components.GetSVGLogo("kaluza"),
				},
				Positions: []components.Position{{Title: "Software Engineer", Current: true}},
				Knowledge: []components.Link{
					{Label: "Typescript", URL: "https://www.typescriptlang.org"},
					{Label: "Github Actions", URL: "https://github.com/features/actions"},
					{Label: "CircleCI", URL: "https://circleci.com"},
					{Label: "NestJS", URL: "https://nestjs.com"},
					{Label: "Terraform", URL: "https://www.terraform.io"},
					{Label: "Kafka", URL: "https://kafka.apache.org"},
					{Label: "DataDog", URL: "https://www.datadoghq.com"},
				},
			},
			{
				Date: components.Date{Start: "Jul 2021", End: "Dec 2023"},
				Link: components.Link{
					Label: "Gentrack",
					URL:   "https://gentrack.com/",
					Logo:  components.GetSVGLogo("gentrack"),
				},
				Positions: []components.Position{
					{Title: "Intermediate Software Engineer"},
					{Title: "Junior Software Engineer"},
					{Title: "Graduate Software Engineer"},
				},
				Knowledge: []components.Link{
					{Label: "SQL"},
					{Label: "API Design"},
					{Label: "Unit Testing"},
					{Label: "Docker", URL: "https://www.docker.com"},
					{Label: "Jenkins", URL: "https://www.jenkins.io"},
				},
			},
			{
				Date: components.Date{Start: "Feb 2018", End: "Jul 2021"},
				Link: components.Link{
					Label: "Proquip Rental & Sales",
					URL:   "https://pqrs.com.au/",
					Logo:  components.GetSVGLogo("proquip"),
				},
				Positions: []components.Position{
					{Title: "IT Support Specialist"},
					{Title: "IT/Marketing Assistant"},
					{Title: "Administrative Assistant"},
				},
				Knowledge: []components.Link{
					{Label: "IT Support"},
					{Label: "Social Media Marketing"},
					{Label: "Adobe Suite", URL: "https://www.adobe.com/products/catalog.html"},
					{Label: "Wordpress", URL: "https://wordpress.com"},
					{Label: "Google Analytics", URL: "https://analytics.google.com/analytics"},
				},
			},
			{
				Date: components.Date{Start: "Feb 2018", End: "Feb 2021"},
				Link: components.Link{
					Label: "University of Melbourne",
					URL:   "https://www.unimelb.edu.au/",
					Logo:  components.GetSVGLogo("melbourneuniversity"),
				},
				Positions: []components.Position{
					{Title: "Bachelor of Science: Computing and Software Systems"},
				},
				Knowledge: []components.Link{
					{Label: "Course Overview", URL: "https://study.unimelb.edu.au/find/courses/major/computing-and-software-systems"},
				},
				Education: true,
			},
		},
		ToolSections: []components.ToolSection{
			{
				Title: "Built with:",
				Links: []components.Link{
					{
						Label: "Go",
						URL:   "https://go.dev/",
						Logo:  components.GetSVGLogo("go"),
					},
					{
						Label: "htmx",
						URL:   "https://htmx.org",
						Logo:  components.GetSVGLogo("htmx"),
					},
					{
						Label: "Tailwind CSS",
						URL:   "https://tailwindcss.com",
						Logo:  components.GetSVGLogo("tailwind"),
					},
				},
			},
			{
				Title: "Deployed with:",
				Links: []components.Link{
					{
						Label: "Netlify",
						URL:   "https://netlify.com",
						Logo:  components.GetSVGLogo("netlify"),
					},
				},
			},
		},
		Copyright: components.Copyright("Nathan Doolan"),
	})
}
