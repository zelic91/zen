package config

type Config struct {
	Commands  []Command
	Databases []Database
	Services  []Service
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
