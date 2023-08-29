package config

type Config struct {
	Title       string
	Description string
	ModuleName  string `yaml:"moduleName"`
	Api         Api
	Commands    map[string]Command
	Databases   map[string]Database
	Services    []Service
	Env         map[string]Env
}

type Command struct {
	Name string
	Type string
}

type Database struct {
	Type   string
	Models []Model
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

type Api struct {
	Title       string
	Description string
	Paths       map[string]map[string]ApiPath
	Schemas     map[string]Schema
}

type ApiPath struct {
	Method          string
	Summary         string
	OperationId     string `yaml:"operationId"`
	RequestBody     string `yaml:"requestBody"`
	SuccessResponse int    `yaml:"successResponse"`
	ResponseBody    string `yaml:"responseBody"`
	Security        string
}

type Schema struct {
	Name       string
	Required   []string
	Properties map[string]SchemaProperty
}

type SchemaProperty struct {
	Type string
}
