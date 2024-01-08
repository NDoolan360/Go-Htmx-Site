package api

import (
	"html/template"
	"net/http"

	utils "github.com/NDoolan360/go-htmx-site/src"
)

type Experience struct {
	Date
	Workplace
	Positions []Position
	Topics    []string
	Education bool
}

type Date struct {
	Start string
	End   string
}

type Workplace struct {
	Name    string
	Logo    string
	LogoSVG template.HTML
	Link    string
}

type Position struct {
	Title   string
	Current bool
}

var experiences = []Experience{
	{
		Date{"Jan 2024", "Present"},
		Workplace{
			Name: "Kaluza",
			Logo: "kaluza.svg",
			Link: "https://kaluza.com/",
		},
		[]Position{{"Software Engineer", true}},
		[]string{"Typescript", "Git", "Github Actions"},
		false,
	},
	{
		Date{"Jul 2021", "Dec 2023"},
		Workplace{
			Name: "Gentrack",
			Logo: "gentrack.svg",
			Link: "https://gentrack.com/",
		},
		[]Position{
			{"Intermediate Software Engineer", false},
			{"Junior Software Engineer", false},
			{"Graduate Software Engineer", false},
		},
		[]string{"Git", "SQL", "Github Actions", "Docker", "Jenkins", "API Design", "Unit Testing"},
		false,
	},
	{
		Date{"Feb 2018", "Jul 2021"},
		Workplace{
			Name: "Proquip Rental & Sales",
			Logo: "proquip.svg",
			Link: "https://pqrs.com.au/",
		},
		[]Position{
			{"IT Support Specialist", false},
			{"IT/Marketing Assistant", false},
			{"Administrative Assistant", false},
		},
		[]string{"IT Support", "Adobe Suite", "Social Media Marketing", "Wordpress", "Google Analytics"},
		false,
	},
	{
		Date{"Feb 2018", "Feb 2021"},
		Workplace{
			Name: "University of Melbourne",
			Logo: "melbourneuniversity.svg",
			Link: "https://www.unimelb.edu.au/",
		},
		[]Position{{"Bachelor of Science: Computing and Software Systems", false}},
		[]string{},
		true,
	},
}

func GetExperience(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(utils.GetTemplate("experience.gohtml")))
	for _, experience := range experiences {
		if svg, err := utils.GetSVGLogo(experience.Workplace.Logo); err == nil {
			experience.Workplace.LogoSVG = svg
		} else {
			experience.Workplace.Logo = ""
		}
		tmpl.Execute(w, experience)
	}
}
