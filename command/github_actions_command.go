package command

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func NewGithubActionsCommand() *BaseCommand {
	return &BaseCommand{
		Command: GenGithubActions,
		Mappings: []TemplateMapping{
			{
				Template:   "github-action-go.tmpl",
				OutputFile: "./.github/workflows/main.yaml",
			},
		},
		ReadArgsFunc: func(b *BaseCommand) {
			prompt := promptui.Prompt{
				Label: "Docker Repository",
			}

			result, err := prompt.Run()

			if err != nil {
				fmt.Printf("error reading Docker repository %v", err)
				return
			}

			b.Args = struct {
				DockerRepo string
			}{
				DockerRepo: result,
			}
		},
	}
}
