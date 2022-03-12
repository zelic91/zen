package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
)

const (
	GenDocker        = "GenDocker"
	GenGithubActions = "GenGithubActions"
	GenMakefile      = "GenMakefile"

	TemplateDocker        = "docker-go.tmpl"
	TemplateGithubActions = "github-action-go.tmpl"
	TemplateMakefile      = "make-go.tmpl"
)

var (
	//go:embed templates/*.tmpl
	rootFs embed.FS

	mapping map[string]string

	choices []string

	templates *template.Template
)

func init() {
	mapping = map[string]string{
		TemplateDocker:        "Dockerfile",
		TemplateGithubActions: "./.github/workflows/main.yaml",
		TemplateMakefile:      "Makefile",
	}

	choices = []string{
		GenDocker,
		GenGithubActions,
		GenMakefile,
	}

	var err error
	templates, err = template.ParseFS(rootFs, "templates/*.tmpl")
	if err != nil {
		fmt.Printf("Error parsing templates: %v\n", err)
	}
}

func main() {

	prompt := promptui.Select{
		Label: "Select generating option:",
		Items: choices,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		return
	}

	switch result {
	case GenDocker:
		ExeGenDocker()
	case GenMakefile:
		ExeGenMakefile()
	case GenGithubActions:
		ExeGenGithubActions()
	default:
		fmt.Println("Not implemented.")
	}
}

func ExeGenDocker() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	elements := strings.Split(path, "/")
	folder := elements[len(elements)-1]
	log.Println(folder)

	appVars := struct {
		Folder string
	}{
		Folder: folder,
	}

	var filePath *os.File
	out := mapping[TemplateDocker]

	if filePath, err = os.Create(out); err != nil {
		log.Fatal(err)
	}

	defer filePath.Close()

	if err = templates.ExecuteTemplate(filePath, TemplateDocker, appVars); err != nil {
		log.Fatal(err)
	}
}

func ExeGenGithubActions() {
	prompt := promptui.Prompt{
		Label: "Docker Repository",
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("error reading Docker repository %v", err)
		return
	}

	appVars := struct {
		DockerRepo string
	}{
		DockerRepo: result,
	}

	var filePath *os.File
	out := mapping[TemplateGithubActions]

	dirName := filepath.Dir(out)
	err = os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		fmt.Printf("error creating parent dirs %v", err)
		return
	}

	filePath, err = os.Create(out)
	if err != nil {
		fmt.Printf("error creating output file %v", err)
		return
	}

	defer filePath.Close()

	if err := templates.ExecuteTemplate(filePath, TemplateGithubActions, appVars); err != nil {
		log.Fatal(err)
	}
}

func ExeGenMakefile() {
	var filePath *os.File
	out := mapping[TemplateMakefile]

	filePath, err := os.Create(out)
	if err != nil {
		log.Fatal(err)
	}

	defer filePath.Close()

	if err := templates.ExecuteTemplate(filePath, TemplateMakefile, nil); err != nil {
		log.Fatal(err)
	}
}
