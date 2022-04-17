package command

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/zelic91/zen/rand"
)

func NewDockerComposeCommand() *BaseCommand {
	return &BaseCommand{
		Command: GenDockerCompose,
		Mappings: []TemplateMapping{
			{
				Template:   "docker-compose.tmpl",
				OutputFile: "docker-compose.yaml",
			},
		},
		ReadArgsFunc: func(b *BaseCommand) {
			prompt := promptui.Prompt{
				Label: "Docker Image for your backend",
			}

			result, err := prompt.Run()

			if err != nil {
				fmt.Printf("error reading Docker image %v", err)
				return
			}

			b.Args = struct {
				Image            string
				PostgresUser     string
				PostgresDB       string
				PostgresPassword string
			}{
				Image:            result,
				PostgresUser:     rand.RandomAlphabet(32),
				PostgresDB:       rand.RandomAlphabet(32),
				PostgresPassword: rand.RandomAlphaNumeric(64),
			}
		},
	}
}
