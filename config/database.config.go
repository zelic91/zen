package config

type Database struct {
	Name      string
	Type      string
	ModelRefs []string `yaml:"modelRefs"`
	Models    []Model
}
