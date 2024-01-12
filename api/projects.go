package api

import (
	"fmt"
	"html/template"
	"net/http"

	githublangsgo "github.com/NDoolan360/github-langs-go"
	"github.com/NDoolan360/go-htmx-site/utils"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {
	projects, errs := FetchProjects(r.URL.Query()["host"])
	if len(errs) > 0 {
		var errorMessages string
		for _, err := range errs {
			errorMessages += err.Error() + "\n"
		}
		http.Error(w, errorMessages, http.StatusInternalServerError)
	} else {
		tmpl := template.Must(template.ParseFiles(
			utils.GetApiResource("template/projects.gohtml"),
		))
		err := tmpl.Execute(w, utils.ProjectsTemplate{Projects: projects})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func FetchProjects(hosts []string) ([]utils.Project, []error) {
	projects := []utils.Project{}
	errs := []error{}
	for _, host := range hosts {
		if site, ok := utils.HostMap[host]; !ok {
			errs = append(errs, fmt.Errorf("URL not found for host: %s", host))

		} else if content, fetchErr := utils.Fetch(site.Path); fetchErr != nil {
			errs = append(errs, fmt.Errorf("error fetching content from host %s: %s", host, fetchErr.Error()))

		} else if hostProjects, ParseErr := utils.Parse(content, host, site.Type); ParseErr != nil {
			errs = append(errs, fmt.Errorf("error parsing content from host %s: %s", host, ParseErr.Error()))

		} else {
			for _, project := range hostProjects {
				// Skip unimportant Github Repos
				if host == "github" && (project.Fork || len(project.Topics) == 0) {
					continue
				}
				project.Host = site.Name
				project.LogoSVG = utils.GetSVGLogo(host)
				if project.Language != "" {
					if lang, err := githublangsgo.GetLanguage(project.Language); err == nil {
						project.LanguageColour = template.CSS(lang.Color)
					}
				}
				projects = append(projects, *project)
			}
		}
	}

	return projects, errs
}
