package command

import (
	"log"
	"os"
	"strings"
)

func NewDockerCommand() *BaseCommand {
	return &BaseCommand{
		Command: GenDocker,
		Mappings: []TemplateMapping{
			{
				Template:   "docker-go.tmpl",
				OutputFile: "Dockerfile",
			},
		},
		ReadArgsFunc: func(b *BaseCommand) {
			path, err := os.Getwd()
			if err != nil {
				panic(err)
			}

			elements := strings.Split(path, "/")
			folder := elements[len(elements)-1]
			log.Println(folder)

			b.Args = struct {
				Folder string
			}{
				Folder: folder,
			}
		},
	}
}
