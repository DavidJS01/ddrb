package file

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"ddrb/education"
	"ddrb/experience"
	"ddrb/frontmatter"
	"ddrb/projects"
	"ddrb/skills"
	"ddrb/user"

	"gopkg.in/yaml.v3"
)

type InputFile struct {
	User       user.User               `json:"user" yaml:"user"`
	Education  []education.Education   `json:"education" yaml:"education"`
	Skills     []skills.Skill          `json:"skills" yaml:"skills"`
	Experience []experience.Experience `json:"experience" yaml:"experience"`
	Projects   []projects.Project      `json:"projects" yaml:"projects"`
}

func GetInput(inputFilePath string) (InputFile, error) {
	file, err := os.Open(fmt.Sprintf(inputFilePath))
	if err != nil {
		return InputFile{}, fmt.Errorf("getInput: error opening file %w", err)

	}
	fmt.Printf("attempting to parse file %s\n", file.Name())
	filePathEnding := strings.Split(inputFilePath, ".")
	input, err := parseFile(file, filePathEnding[len(filePathEnding)-1])
	if err != nil {
		return input, err
	}
	return input, nil
}

func parseFile(file *os.File, fileFormat string) (InputFile, error) {
	var input InputFile

	switch fileFormat {
	case "json":
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&input); err != nil {
			return InputFile{}, fmt.Errorf("parseFile: error decoding file: %w", err)
		}
	case "yaml":
		decoder := yaml.NewDecoder(file)
		if err := decoder.Decode(&input); err != nil {
			return InputFile{}, fmt.Errorf("parseFile: error decoding file: %w", err)
		}
	default:
		return InputFile{}, fmt.Errorf("parseFile: unexpected file type %s found, expected one of `json` `yaml`", fileFormat)
	}

	return input, nil
}

func BuildLatexOutput(inp InputFile) string {
	frontmatter := frontmatter.GetFrontmatter()
	profile := user.BuildProfile(inp.User)
	education := education.BuildEducation(inp.Education)
	skills := skills.BuildSkills(inp.Skills)
	experience := experience.BuildExperience(inp.Experience)
	projects := projects.BuildProjects(inp.Projects)

	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n\\end{document}", frontmatter, profile, education, skills, experience, projects)

}

func WriteLatexToFile(latexOutput string) (string, error) {
	now := time.Now()

	formattedDate := now.Format("2006_01_02")
	resumeFileName := fmt.Sprintf("resume_%s.tex", formattedDate)

	file, err := os.Create(resumeFileName)
	if err != nil {
		return "", fmt.Errorf("writeLatexToFile: error creating file:", err)
	}
	defer file.Close()

	_, err = file.WriteString(latexOutput)
	if err != nil {
		fmt.Errorf("Error writing to file:", err)
		return "", fmt.Errorf("writeLatexToFile: error writing to file:", err)
	}

	fmt.Printf("\nWrote latex output to %s successfully", resumeFileName)
	return resumeFileName, nil

}

func LatexToPDF(fileName string) error {
	cmd := exec.Command("pdflatex", fileName)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
