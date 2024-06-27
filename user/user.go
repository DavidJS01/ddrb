package user

import (
	"fmt"
)

type User struct {
	Name        string `json:"name" yaml:"name"`
	Email       string `json:"email" yaml:"email"`
	PhoneNumber string `json:"phone_number" yaml:"phone_number"`
	GithubURL   string `json:"github_url" yaml:"github_url"`
}

func BuildProfile(user User) string {
	return fmt.Sprintf(`
\vspace*{-10pt}
\begin{center}
	{\Huge \scshape {%s}}\\
	%s $\cdot$ %s $\cdot$ %s
\end{center}`, user.Name, user.Email, user.PhoneNumber, user.GithubURL)
}
