package config

type Api struct {
	Title       string
	Description string
	Resources   []Resource
}

type Resource struct {
	Security    string
	ModelRef    string `yaml:"modelRef"`
	DatabaseRef string `yaml:"databaseRef"`
	Model       Model
	Database    Database
}
