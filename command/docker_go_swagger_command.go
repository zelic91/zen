package command

import (
	"log"
	"os"
	"strings"
)

func NewDockerGoSwaggerCommand() *BaseCommand {
	return &BaseCommand{
		Command: GenDockerSwagger,
		Mappings: []TemplateMapping{
			{
				Template:   "docker-go-swagger.tmpl",
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
