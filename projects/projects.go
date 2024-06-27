package projects

import (
	"fmt"
)

type Project struct {
	Title        string   `json:"title" yaml:"title"`
	Technologies []string `json:"technologies" yaml:"technologies"`
	URL          string   `json:"url" yaml:"url"`
	Description  string   `json:"description" yaml:"description"`
}

func buildProjectsBlock(project Project) string {
	// order: title, technologies, url, description
	projectTemplate := `{\textbf{%s}} {\sl %s} \hfill %s\\
%s\\\vspace{1mm}`
	technologies := ""
	for i, technology := range project.Technologies {
		if i == len(project.Technologies)-1 {
			technologies = technologies + technology
		} else {
			technologies = technologies + technology + ", "
		}
	}
	return fmt.Sprintf(projectTemplate, project.Title, technologies, project.URL, project.Description)
}

func BuildProjects(projects []Project) string {
	projectBlockTemplate := `\header{Projects}
\vspace{1mm}
%s
\vspace*{2mm}`

	blockContents := ""
	for _, project := range projects {
		blockContents = blockContents + buildProjectsBlock(project)
	}

	return fmt.Sprintf(projectBlockTemplate, blockContents)
}
