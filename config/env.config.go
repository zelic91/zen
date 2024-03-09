package config

type Env struct {
	Name         string
	StructName   string `yaml:"structName"`
	Type         string
	DefaultValue string `yaml:"defaultValue"`
	Secret       bool
}
