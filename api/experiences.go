package api

import (
	"html/template"
	"net/http"

	utils "github.com/NDoolan360/go-htmx-site/src"
)

type Experiences struct {
	Experiences []Experience
}

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
	Name string
	Logo template.HTML
	Link string
}

type Position struct {
	Title   string
	Current bool
}

func GetExperiences(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles(
		utils.GetTemplatePath("experiences.gohtml"),
	)).Execute(w, Experiences{
		[]Experience{
			{
				Date{"Jan 2024", "Present"},
				Workplace{
					Name: "Kaluza",
					Logo: utils.IgnoreErr(utils.GetSVGLogo("kaluza")),
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
					Logo: utils.IgnoreErr(utils.GetSVGLogo("gentrack")),
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
					Logo: utils.IgnoreErr(utils.GetSVGLogo("proquip")),
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
					Logo: utils.IgnoreErr(utils.GetSVGLogo("melbourneuniversity")),
					Link: "https://www.unimelb.edu.au/",
				},
				[]Position{{"Bachelor of Science: Computing and Software Systems", false}},
				[]string{},
				true,
			},
		},
	})
}
