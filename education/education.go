
package education

import (
	"fmt"
)

/*
 DatesAttended fields probably is a bad name.
 That field is a single string because I want to support use cases
 like: "2024" "Aug 2023 - Aug 2024" etc
*/
type Education struct {
	Name          string `json:"name" yaml:"name"`
	Location      string `json:"location" yaml:"location"`
	Degree        string `json:"degree" yaml:"degree"`
	DatesAttended string `json:"dates_attended" yaml:"dates_attended"`
}

func BuildEducation(institutions []Education) string {
	headerTemplate := `\header{Education}`
	educationTemplate := `\textbf{%s}\hfill %s\\
%s \hfill %s\\
\vspace{2mm}`

	output := ""
	for _, i := range institutions {
		education := fmt.Sprintf(educationTemplate, i.Name, i.Location, i.Degree, i.DatesAttended)
		output = fmt.Sprintf("%s\n%s", output, education)
	}

	
	return fmt.Sprintf("%s\n%s", headerTemplate, output)
}
