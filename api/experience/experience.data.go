package main

import (
	"github.com/NDoolan360/go-htmx-site/web/templates"
)

var workExperiences = []Experience{
	{
		DateStart: "Jan 2024",
		DateEnd:   "Present",
		Location: Place{
			Name: "Kaluza",
			Logo: templates.KaluzaLogo(),
		},
		Positions: []Position{
			{Role: "Software Engineer", Active: true},
		},
		Link: "https://kaluza.com",
		Topics: []Topic{
			{Label: "Typescript", Link: "https://typescriptlang.org"},
			{Label: "Github Actions", Link: "https://github.com/features/actions"},
			{Label: "DataDog", Link: "https://datadoghq.com"},
			{Label: "Aiven", Link: "https://aiven.io"},
			{Label: "Kafka", Link: "https://kafka.apache.org"},
			{Label: "Argo CD", Link: "https://argoproj.github.io/cd"},
			{Label: "Kubernetes", Link: "https://kubernetes.io"},
			{Label: "Terraform", Link: "https://terraform.io"},
		},
	},
	{
		DateStart: "Jul 2021",
		DateEnd:   "Dec 2023",
		Location: Place{
			Name: "Gentrack",
			Logo: templates.GentrackLogo(),
		},
		Positions: []Position{
			{Role: "Intermediate Software Engineer", Active: false},
			{Role: "Junior Software Engineer", Active: false},
			{Role: "Graduate Software Engineer", Active: false},
		},
		Link: "https://gentrack.com",
		Topics: []Topic{
			{Label: "SQL"},
			{Label: "API Design"},
			{Label: "Unit Testing"},
			{Label: "Docker", Link: "https://docker.com"},
			{Label: "Jenkins", Link: "https://jenkins.io"},
		},
	},
	{
		DateStart: "Feb 2018",
		DateEnd:   "Jul 2021",
		Location: Place{
			Name: "Proquip Rental & Sales",
			Logo: templates.ProquipLogo(),
		},
		Positions: []Position{
			{Role: "IT Support Specialist", Active: false},
			{Role: "IT/Marketing Assistant", Active: false},
			{Role: "Administrative Assistant", Active: false},
		},
		Link: "https://pqrs.com.au",
		Topics: []Topic{
			{Label: "IT Support"},
			{Label: "Social Media Marketing"},
			{Label: "Adobe Suite", Link: "https://adobe.com/products/catalog.html"},
			{Label: "Wordpress", Link: "https://wordpress.com"},
			{Label: "Google Analytics", Link: "https://analytics.google.com/analytics"},
		},
	},
}

var educationExperiences = []Experience{
	{
		DateStart: "Feb 2018",
		DateEnd:   "Feb 2021",
		Location: Place{
			Name: "University of Melbourne",
			Logo: templates.MelbourneUniversityLogo(),
		},
		Positions: []Position{
			{Role: "Bachelor of Science: Computing and Software Systems", Active: true},
		},
		Link: "https://unimelb.edu.au",
		Topics: []Topic{
			{Label: "Course Overview", Link: "https://study.unimelb.edu.au/find/courses/major/computing-and-software-systems"},
		},
	},
}
