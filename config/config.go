package config

type Config struct {
	ModuleName string `yaml:"moduleName"`
	Commands   []Command
	Databases  []Database
	Services   []Service
	Env        []Env
}

type Command struct {
	Name string
	Type string
}

type Database struct {
	Name string
	Type string
}

type Model struct {
	Name       string
	Properties []ModelProperties
}

type ModelProperties struct {
	Name string
	Type string
	Ref  string
}

type Service struct {
	Name     string
	Database string
}

type Env struct {
	Name       string
	StructName string `yaml:"structName"`
	Type       string
}
