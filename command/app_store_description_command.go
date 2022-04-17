package command

import (
	"log"
	"os"
	"strings"
)

func NewAppStoreDescriptionCommand() *BaseCommand {
	return &BaseCommand{
		Command: GenAppStoreFastlane,
		Mappings: []TemplateMapping{
			{
				Template:   "app_store_description.tmpl",
				OutputFile: "description.txt",
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
