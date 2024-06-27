package main

import (
	"ddrb/file"
	"flag"
	"fmt"
	"os"
	"strings"
)

type supportedInputFormats string

const (
	jsonFormat supportedInputFormats = "json"
	yamlFormat supportedInputFormats = "yaml"
)

var convertToPDF = flag.Bool("pdf", false, "also output pdf using pdflatex")
var inputFilePath = flag.String("file_path", "input.yaml", "absolute input file path")

func exit(message string, exitCode int) {
	fmt.Printf(message)
	os.Exit(exitCode)
}

func parseFlags() {
	flag.Parse()
	fileEnding := strings.Split(*inputFilePath, ".")
	if fileEnding[len(fileEnding)-1] != string(jsonFormat) && fileEnding[len(fileEnding)-1] != string(yamlFormat) {
		exit(fmt.Sprintf("Invalid input file format %s, expected one of %s, %s\n", fileEnding, jsonFormat, yamlFormat), 1)
	}
}

func main() {
	parseFlags()
	inputFile, err := file.GetInput(*inputFilePath)
	if err != nil {
		exit(fmt.Sprintf("Error parsing resume input file: %v\n", err), 1)
	}
	latexOutput := file.BuildLatexOutput(inputFile)
	createdFile, err := file.WriteLatexToFile(latexOutput)
	if err != nil {
		exit(fmt.Sprintf("Error creating latex file: %v\n", err), 1)
	}
	if *convertToPDF {
		if err := file.LatexToPDF(createdFile); err != nil {
			exit(fmt.Sprintf("Error converting latex file %s to PDF: %v\n", createdFile, err), 1)
		}
	}
}
