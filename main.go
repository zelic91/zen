package main

import (
	"embed"
	"fmt"
	"html/template"

	"github.com/manifoldco/promptui"
	"github.com/zelic91/zen/command"
)

var (
	//go:embed templates/*.tmpl
	rootFs embed.FS

	commands []command.CommandInterface

	choices []string

	templates *template.Template
)

func init() {
	choices = []string{
		command.GenDocker,
		command.GenDockerSwagger,
		command.GenDockerCompose,
		command.GenK8sAll,
		command.GenGithubActions,
		command.GenMakefile,
	}

	commands = []command.CommandInterface{
		command.NewDockerCommand(),
		command.NewDockerGoSwaggerCommand(),
		command.NewDockerComposeCommand(),
		command.NewGithubActionsCommand(),
		command.NewK8sAllCommand(),
		command.NewMakefileCommand(),
	}

	var err error
	templates, err = template.ParseFS(rootFs, "templates/*.tmpl")
	if err != nil {
		fmt.Printf("Error parsing templates: %v\n", err)
	}
}

func main() {

	prompt := promptui.Select{
		Label: "What do you want to generate?",
		Items: choices,
		Size:  10,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		return
	}

	for _, c := range commands {
		if com, ok := c.(*command.BaseCommand); ok && com.Command == result {
			c.Exec(templates)
		}
	}
}
