package command

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

const (
	GenDocker        = "🐳 Generate Dockerfile for Go"
	GenDockerSwagger = "🐳 Generate Dockerfile for Go with Swagger"
	GenDockerCompose = "🐳 Generate docker-compose.yaml for typical apps"
	GenK8sAll        = "☸️  Generate k8s deployment & service"
	GenGithubActions = "Generate Github Actions Config for Go"
	GenMakefile      = "Generate Makefile for Go"
)

type CommandInterface interface {
	Exec(templates *template.Template)
	ExecGeneral(templates *template.Template)
}

type TemplateMapping struct {
	Template   string
	OutputFile string
}

type BaseCommand struct {
	Command      string
	Mappings     []TemplateMapping
	ReadArgsFunc func(*BaseCommand)
	Args         interface{}
}

func (b BaseCommand) Exec(templates *template.Template) {
	b.ReadArgsFunc(&b)
	b.ExecGeneral(templates)
}

func (b BaseCommand) ExecGeneral(templates *template.Template) {
	for _, mapping := range b.Mappings {
		var filePath *os.File
		dirName := filepath.Dir(mapping.OutputFile)
		err := os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			fmt.Printf("error creating parent dirs %v", err)
			return
		}

		filePath, err = os.Create(mapping.OutputFile)
		if err != nil {
			fmt.Printf("error creating output file %v", err)
			return
		}

		defer filePath.Close()

		if err := templates.ExecuteTemplate(filePath, mapping.Template, b.Args); err != nil {
			log.Fatal(err)
		}
	}
}
