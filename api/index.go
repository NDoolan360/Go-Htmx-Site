package api

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"testing"
	"time"
)

// GetIndex handles the request for rendering the index page.
func GetIndex(w http.ResponseWriter, r *http.Request) {
	indexTemplate := template.Must(template.ParseFiles(
		GetApiAsset("template/index.gohtml"),
		GetApiAsset("template/head.gohtml"),
		GetApiAsset("template/theme-switch.gohtml"),
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
				Logo:  GetSVGLogo("github"),
			},
			{
				Label: "LinkedIn",
				URL:   "https://www.linkedin.com/in/nathan-doolan-835a13171",
				Logo:  GetSVGLogo("linkedin"),
			},
			{
				Label: "Discord",
				URL:   "https://discord.com/users/nothindoin",
				Logo:  GetSVGLogo("discord"),
			},
			{
				Label: "Cults3D",
				URL:   "https://cults3d.com/en/users/ND360",
				Logo:  GetSVGLogo("cults3d"),
			},
			{
				Label: "Boardgame Geek",
				URL:   "https://boardgamegeek.com/user/Nothin_Doin",
				Logo:  GetSVGLogo("bgg"),
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
					Logo:  GetSVGLogo("kaluza"),
				},
				Positions: []Position{{Title: "Software Engineer", Current: true}},
				Knowledge: []Link{{Label: "Git"}, {Label: "Typescript", URL: "https://www.typescriptlang.org"}, {Label: "Docker", URL: "https://www.docker.com"}, {Label: "Github Actions"}, {Label: "CircleCI", URL: "https://circleci.com"}, {Label: "NestJS", URL: "https://nestjs.com"}, {Label: "GraphQL", URL: "https://graphql.org"}, {Label: "Terraform", URL: "https://www.terraform.io"}, {Label: "Kafka", URL: "https://kafka.apache.org"}, {Label: "Kubernetes", URL: "https://kubernetes.io"}},
			},
			{
				Date: Date{Start: "Jul 2021", End: "Dec 2023"},
				Link: Link{
					Label: "Gentrack",
					URL:   "https://gentrack.com/",
					Logo:  GetSVGLogo("gentrack"),
				},
				Positions: []Position{
					{Title: "Intermediate Software Engineer"},
					{Title: "Junior Software Engineer"},
					{Title: "Graduate Software Engineer"},
				},
				Knowledge: []Link{{Label: "Git"}, {Label: "SQL"}, {Label: "Docker", URL: "https://www.docker.com"}, {Label: "Jenkins", URL: "https://www.jenkins.io"}, {Label: "API Design"}, {Label: "Unit Testing"}},
			},
			{
				Date: Date{Start: "Feb 2018", End: "Jul 2021"},
				Link: Link{
					Label: "Proquip Rental & Sales",
					URL:   "https://pqrs.com.au/",
					Logo:  GetSVGLogo("proquip"),
				},
				Positions: []Position{
					{Title: "IT Support Specialist"},
					{Title: "IT/Marketing Assistant"},
					{Title: "Administrative Assistant"},
				},
				Knowledge: []Link{{Label: "IT Support"}, {Label: "Adobe Suite"}, {Label: "Social Media Marketing"}, {Label: "Wordpress"}, {Label: "Google Analytics"}},
			},
			{
				Date: Date{Start: "Feb 2018", End: "Feb 2021"},
				Link: Link{
					Label: "University of Melbourne",
					URL:   "https://www.unimelb.edu.au/",
					Logo:  GetSVGLogo("melbourneuniversity"),
				},
				Positions: []Position{{Title: "Bachelor of Science: Computing and Software Systems"}},
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
						Logo:  GetSVGLogo("go"),
					},
					{
						Label: "htmx",
						URL:   "https://htmx.org",
						Logo:  GetSVGLogo("htmx"),
					},
					{
						Label: "Tailwind CSS",
						URL:   "https://tailwindcss.com",
						Logo:  GetSVGLogo("tailwind"),
					},
				},
			},
			{
				Title: "Deployed with:",
				Links: []Link{
					{
						Label: "Vercel",
						URL:   "https://vercel.com",
						Logo:  GetSVGLogo("vercel"),
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

// GetApiAsset returns the correct resource path for resources in the api directory,
// based on the current environment.
func GetApiAsset(path string) string {
	if _, inVercel := os.LookupEnv("VERCEL"); inVercel {
		// When running as serverless in Vercel environment, /api is the root.
		// Note: The root directory is not explicitly named "api".
		return "assets/" + path
	} else if testing.Testing() {
		// During testing, api resources are expected to be in the ../api/ directory.
		return "../api/assets/" + path
	} else {
		// Assume it is being run from the root, and api resources are in the api/ directory.
		return "api/assets/" + path
	}
}

// GetSVGLogo reads an SVG logo file and returns it as an HTML template.
func GetSVGLogo(logo string) template.HTML {
	svg, err := os.ReadFile(GetApiAsset("logos/" + logo + ".svg"))
	if err != nil {
		panic(err)
	}
	return template.HTML(svg)
}
