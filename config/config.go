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
	Name   string
	Type   string
	Models []Model
}

func (d *Database) IsPostgres() bool {
	return d.Name == "postgres"
}

type Model struct {
	Name       string
	Type       string
	Owner      string
	Properties []ModelProperties
}

type ModelProperties struct {
	Name       string
	Type       string
	Owner      string
	NotNull    bool `yaml:"notNull"`
	References string
	Unique     bool
}

type ModelReference struct {
	ForeignKey string
	From       string
	To         string
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
