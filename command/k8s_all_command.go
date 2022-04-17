package command

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type K8sAllCommand struct {
	BaseCommand
}

func NewK8sAllCommand() *BaseCommand {
	return &BaseCommand{
		Command: GenK8sAll,
		Mappings: []TemplateMapping{
			{
				Template:   "k8s-deployment.tmpl",
				OutputFile: "deployment/deployment.yaml",
			},
			{
				Template:   "k8s-service.tmpl",
				OutputFile: "deployment/service.yaml",
			},
		},
		ReadArgsFunc: func(b *BaseCommand) {
			prompt := promptui.Prompt{
				Label: "Docker Image",
			}

			result, err := prompt.Run()

			if err != nil {
				fmt.Printf("error reading Docker image %v", err)
				return
			}

			b.Args = struct {
				Image string
			}{
				Image: result,
			}
		},
	}
}
