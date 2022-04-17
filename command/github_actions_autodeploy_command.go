package command

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func NewGithubActionsAutoDeployCommand() *BaseCommand {
	return &BaseCommand{
		Command: GenGithubActionsAutoDeploy,
		Mappings: []TemplateMapping{
			{
				Template:   "github-action-go-deploy.tmpl",
				OutputFile: "./.github/workflows/main.yaml",
			},
		},
		ReadArgsFunc: func(b *BaseCommand) {
			prompt := promptui.Prompt{
				Label: "Docker Repository",
			}

			repo, err := prompt.Run()

			if err != nil {
				fmt.Printf("error reading Docker repository %v", err)
				return
			}

			prompt = promptui.Prompt{
				Label: "Target Remote Folder",
			}

			folder, err := prompt.Run()

			if err != nil {
				fmt.Printf("error reading Target Remote Folder %v", err)
				return
			}

			b.Args = struct {
				DockerRepo string
				Folder     string
			}{
				DockerRepo: repo,
				Folder:     folder,
			}
		},
	}
}
