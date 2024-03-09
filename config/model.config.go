package config

type Model struct {
	Name       string `yaml:"name"`
	Type       string
	Owner      string
	Properties map[string]ModelProperty
}

type ModelProperty struct {
	Type       string
	Owner      string
	NotNull    bool `default:"true" yaml:"notNull"`
	Unique     bool
	Ref        string
	Searchable bool
}
