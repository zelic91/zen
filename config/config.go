package config

type Config struct {
	Title               string
	Description         string
	ModuleName          string `yaml:"moduleName"`
	Api                 Api
	Commands            map[string]Command
	Databases           map[string]Database
	Services            map[string]Service
	Env                 map[string]Env
	CurrentPackage      string
	CurrentModelName    string
	CurrentModel        Model
	CurrentCommand      string
	ServiceOperationMap map[string][]ApiPath
	ServiceDatabaseMap  map[string]Database
}

type Command struct {
	Name string
	Type string
}

type Database struct {
	Type   string
	Models map[string]Model
}

type Model struct {
	Name       string
	Type       string
	Owner      string
	Properties map[string]ModelProperties
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
	Database string
	Model    string
	Methods  []string
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
	Service         string
	OperationID     string `yaml:"operationId"`
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
