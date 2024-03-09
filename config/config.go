package config

import "embed"

type Config struct {
	// Root template path points to the directory which hosts
	// every templates.
	// The templates will be organized into subfolders
	// by their functionalities. For example:
	// - root: all necessary files for the project roots: Makefile, sample env, etc.
	RootTemplatePath string
	RootFs           embed.FS
	ModuleName       string `yaml:"moduleName"`
	Title            string
	Description      string
	Env              map[string]Env

	// Functionalities
	Models     []Model
	Databases  []Database
	Api        Api
	Commands   map[string]Command
	Deployment Deployment
	Crawler    Crawler

	// Access Utils
	ModelMap    map[string]Model
	DatabaseMap map[string]Database
}
