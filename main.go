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
	GenDocker        = "🐳 Generate Dockerfile for Go"
	GenDockerSwagger = "🐳 Generate Dockerfile for Go with Swagger"
	GenK8sDeployment = "☸️ Generate k8s deployment"
	GenK8sService    = "☸️ Generate k8s service"
	GenGithubActions = "Generate Github Actions Config for Go"
	GenMakefile      = "Generate Makefile for Go"

	TemplateDocker        = "docker-go.tmpl"
	TemplateDockerSwagger = "docker-go-swagger.tmpl"
	TemplateK8sDeployment = "k8s-deployment.tmpl"
	TemplateK8sService    = "k8s-service.tmpl"
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
		TemplateDockerSwagger: "Dockerfile",
		TemplateK8sDeployment: "deployment.yaml",
		TemplateK8sService:    "service.yaml",
		TemplateGithubActions: "./.github/workflows/main.yaml",
		TemplateMakefile:      "Makefile",
	}

	choices = []string{
		GenDocker,
		GenDockerSwagger,
		GenK8sDeployment,
		GenK8sService,
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
		Label: "Select generating option",
		Items: choices,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		return
	}

	switch result {
	case GenDocker:
		ExecGenDocker()
	case GenDockerSwagger:
		ExecGenDockerSwagger()
	case GenK8sDeployment:
		ExecGenK8sDeployment()
	case GenK8sService:
		ExecGenK8sService()
	case GenMakefile:
		ExecGenMakefile()
	case GenGithubActions:
		ExecGenGithubActions()
	default:
		fmt.Println("Not implemented.")
	}
}

func ExecGenDocker() {
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

	ExecGeneral(TemplateDocker, appVars)
}

func ExecGenDockerSwagger() {
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

	ExecGeneral(TemplateDockerSwagger, appVars)
}

func ExecGenGithubActions() {
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

	ExecGeneral(TemplateGithubActions, appVars)
}

func ExecGenMakefile() {
	ExecGeneral(TemplateMakefile, nil)
}

func ExecGenK8sDeployment() {
	prompt := promptui.Prompt{
		Label: "Docker Image",
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("error reading Docker image %v", err)
		return
	}

	appVars := struct {
		Image string
	}{
		Image: result,
	}
	ExecGeneral(TemplateK8sDeployment, appVars)
}

func ExecGenK8sService() {
	ExecGeneral(TemplateK8sService, nil)
}

func ExecGeneral(templateType string, vars interface{}) {
	var filePath *os.File
	out := mapping[templateType]

	dirName := filepath.Dir(out)
	err := os.MkdirAll(dirName, os.ModePerm)
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

	if err := templates.ExecuteTemplate(filePath, templateType, vars); err != nil {
		log.Fatal(err)
	}
}
