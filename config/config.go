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
	CurrentServiceName  string
	CurrentService      Service
	ServiceOperationMap map[string][]ApiPath
	ServiceDatabaseMap  map[string]Database
}

type Command struct {
	Name string
	Type string
}

type Database struct {
	Type   string
	Models []Model
}

func (db Database) ModelWithName(name string) *Model {
	for _, model := range db.Models {
		if model.Name == name {
			return &model
		}
	}
	return nil
}

type Model struct {
	Name       string `yaml:"name"`
	Type       string
	Owner      string
	SearchBy   []string `yaml:"searchBy"`
	FindBy     []string `yaml:"findBy"`
	Properties map[string]ModelProperty
}

type ModelProperty struct {
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
	Type         string
	AuthService  string `yaml:"authService"`
	ScaffoldCRUD bool   `yaml:"scaffoldCRUD"`
	Database     string
	Model        string
	Methods      map[string]ServiceMethod
	Services     map[string][]string
}

type ServiceMethod struct {
	UseRepo          bool   `yaml:"useRepo"`
	UseService       bool   `yaml:"useService"`
	UseServiceMethod string `yaml:"useServiceMethod"`
	Arguments        []MethodArgument
	Returns          []string
}

type MethodArgument struct {
	Name string
	Type string
}

type Env struct {
	Name         string
	StructName   string `yaml:"structName"`
	Type         string
	DefaultValue string `yaml:"defaultValue"`
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
