package experience

import (
	"fmt"
)

type Experience struct {
	OrganizationName string   `json:"organization_name" yaml:"organization_name"`
	PositionTitle    string   `json:"position_title" yaml:"position_title"`
	Location         string   `json:"location" yaml:"location"`
	WorkedDates      string   `json:"worked_dates" yaml:"worked_dates"`
	BulletPoints     []string `json:"bullet_points" yaml:"bullet_points"`
}

func buildBulletPoints(job Experience) string {
	builtString := ``
	for _, bulletPoint := range job.BulletPoints {
		builtString = builtString + fmt.Sprintf(`\item %s`, bulletPoint)
	}
	return builtString
}

func buildExperienceBlock(job Experience) string {
	experienceItemTemplate := `\textbf{%s} \hfill %s \\
%s \hfill %s\\
\vspace{-3mm}
`
	experienceBulletPointsTemplate := `
\begin{itemize} \itemsep 1pt
    %s
\end{itemize}`

	bulletPoints := buildBulletPoints(job)

	return fmt.Sprintf(experienceItemTemplate, job.OrganizationName, job.Location, job.PositionTitle, job.WorkedDates) + fmt.Sprintf(experienceBulletPointsTemplate, bulletPoints)
}

func BuildExperience(jobs []Experience) string {
	headerTemplate := `\header{Experience}
\vspace{1mm}
%s
`
	jobExperience := ""
	for _, job := range jobs {
		jobExperience = jobExperience + buildExperienceBlock(job)
	}

	return fmt.Sprintf(headerTemplate, jobExperience)
}
