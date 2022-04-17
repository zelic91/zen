package command

import (
	"log"
	"os"
	"strings"
)

type MakefileCommand struct {
	BaseCommand
}

func NewMakefileCommand() *BaseCommand {
	return &BaseCommand{
		Command: GenMakefile,
		Mappings: []TemplateMapping{
			{
				Template:   "make-go.tmpl",
				OutputFile: "Makefile",
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
