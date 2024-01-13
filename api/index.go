package api

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"testing"
	"time"
)

// IndexTemplate represents the data structure for the index.gohtml template.
type IndexTemplate struct {
	// Basic page information for <hrad>
	Title       string
	Description string

	// Navigation links
	InternalLinks  []Link
	ExternalLinks  []Link
	ThemeSwitchSVG template.HTML

	// Profile information
	Profile Profile

	// Professional and educational experiences
	Experiences []Experience

	// Tools and technologies used section
	ToolSections []ToolSection

	// Copyright string
	Copyright string
}

// Link represents a hyperlink with label, URL, and optional logo.
type Link struct {
	Label string
	URL   template.URL
	Logo  template.HTML
}

// Profile represents user profile information with image and paragraphs.
type Profile struct {
	ImageAttr  []template.HTMLAttr
	Paragraphs []string
}

// Experience represents a work experience or education entry.
type Experience struct {
	Date
	Workplace Link
	Positions []Position
	Topics    []string
	Education bool
}

// ToolSection represents a section with a title and links to tools used.
type ToolSection struct {
	Title string
	Links []Link
}

// GetIndex handles the request for rendering the index page.
func GetIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		GetApiResource("template/index.gohtml"),
		GetApiResource("template/head.gohtml"),
	))

	err := tmpl.Execute(w, IndexTemplate{
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
		ThemeSwitchSVG: GetSVGLogo("theme_switch"),
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
				Workplace: Link{
					Label: "Kaluza",
					URL:   "https://kaluza.com/",
					Logo:  GetSVGLogo("kaluza"),
				},
				Positions: []Position{{Title: "Software Engineer", Current: true}},
				Topics:    []string{"Typescript", "Git", "Github Actions"},
			},
			{
				Date: Date{Start: "Jul 2021", End: "Dec 2023"},
				Workplace: Link{
					Label: "Gentrack",
					URL:   "https://gentrack.com/",
					Logo:  GetSVGLogo("gentrack"),
				},
				Positions: []Position{
					{Title: "Intermediate Software Engineer"},
					{Title: "Junior Software Engineer"},
					{Title: "Graduate Software Engineer"},
				},
				Topics: []string{"Git", "SQL", "Github Actions", "Docker", "Jenkins", "API Design", "Unit Testing"},
			},
			{
				Date: Date{Start: "Feb 2018", End: "Jul 2021"},
				Workplace: Link{
					Label: "Proquip Rental & Sales",
					URL:   "https://pqrs.com.au/",
					Logo:  GetSVGLogo("proquip"),
				},
				Positions: []Position{
					{Title: "IT Support Specialist"},
					{Title: "IT/Marketing Assistant"},
					{Title: "Administrative Assistant"},
				},
				Topics: []string{"IT Support", "Adobe Suite", "Social Media Marketing", "Wordpress", "Google Analytics"},
			},
			{
				Date: Date{Start: "Feb 2018", End: "Feb 2021"},
				Workplace: Link{
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
						Label: "htmx",
						URL:   "https://htmx.org",
						Logo:  GetSVGLogo("htmx"),
					},
					{
						Label: "hyperscript",
						URL:   "https://hyperscript.org",
						Logo:  GetSVGLogo("hyperscript"),
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

// GetApiResource returns the correct resource path for resources in the api dirextory,
// based on the current environment.
func GetApiResource(path string) string {
	if _, inVercel := os.LookupEnv("VERCEL"); inVercel {
		// When running as serverless in Vercel environment, /api is the root.
		// Note: The root directory is not explicitly named "api".
		return path
	} else if testing.Testing() {
		// During testing, api resources are expected to be in the ../api/ directory.
		return "../api/" + path
	} else {
		// Assume it is being run from the root, and api resources are in the api/ directory.
		return "api/" + path
	}
}

// GetSVGLogo reads an SVG logo file and returns it as an HTML template.
func GetSVGLogo(logo string) template.HTML {
	svg, err := os.ReadFile(GetApiResource("logos/" + logo + ".svg"))
	if err != nil {
		panic(err)
	}
	return template.HTML(svg)
}
