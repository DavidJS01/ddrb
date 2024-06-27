package skills

import (
	"fmt"
)

type Skill struct {
	Category string   `json:"category" yaml:"category"`
	Items    []string `json:"skills" yaml:"skills"`
}

func buildSkillsBlock(skill Skill) string {
	skillTemplate := "%s: & %s"
	categorySkills := ""

	for index, item := range skill.Items {
		// create a string for each item
		if index == len(skill.Items)-1 {
			categorySkills = categorySkills + item + `\\`
		} else {
			categorySkills = categorySkills + item + ", "
		}
	}

	return fmt.Sprintf(skillTemplate, skill.Category, categorySkills)
}

func BuildSkills(skills []Skill) string {
	template := `\header{Skills}
\begin{tabular}{ l l }
%s
\end{tabular}
\vspace{2mm}
`
	skillBlock := ``
	for _, skill := range skills {
		skillBlock = skillBlock + buildSkillsBlock(skill)
	}

	return fmt.Sprintf(template, skillBlock)

}
