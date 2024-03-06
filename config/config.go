package config

import "embed"

type Config struct {
	// Root template path points to the directory which hosts
	// every templates.
	// The templates will be organized into subfolders
	// by their functionalities. For example:
	// - root: all necessary files for the project roots: Makefile, sample env, etc.
	RootTemplatePath string
	RootFs           embed.FS
	Title            string
	Description      string
	Env              map[string]Env
	ModuleName       string `yaml:"moduleName"`
	Api              Api
	Commands         map[string]Command
	Databases        map[string]Database
	Deployment       Deployment
	Crawler          Crawler

	// TODO: Deprecating below code
	Services []Service

	CurrentPackage          string
	CurrentModelName        string
	CurrentModel            Model
	CurrentCommand          string
	CurrentServiceName      string
	CurrentService          Service
	ServiceOperationMap     map[string][]ApiPath
	ServiceDatabaseMap      map[string]Database
	ServiceCrawlerTargetMap map[string][]CrawlerTarget
}

func (c Config) ServiceWithName(name string) *Service {
	for _, service := range c.Services {
		if service.Name == name {
			return &service
		}
	}
	return nil
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
	NotNull    bool `default:"true" yaml:"notNull"`
	References string
	Unique     bool
}

type ModelReference struct {
	ForeignKey string
	From       string
	To         string
}

type Service struct {
	Name        string
	Type        string
	AuthService string `yaml:"authService"`
	Database    string
	Model       string
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
	Secret       bool
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
	Params          []ApiParam
}

type ApiParam struct {
	Name     string
	In       string
	Required bool
	Type     string
	Format   string
}

type Schema struct {
	Name       string
	Required   []string
	Properties map[string]SchemaProperty
}

type SchemaProperty struct {
	Type string
}

type Deployment struct {
	Host              string
	Email             string
	SecretName        string `yaml:"secretName"`
	DockerHubUsername string `yaml:"dockerHubUsername"`
	DockerHubRepo     string `yaml:"dockerHubRepo"`
	TargetPort        string `yaml:"targetPort"`
}

type Crawler struct {
	PostDelayPerCrawl   int             `yaml:"postDelayPerCrawl"`
	PostDelayPerRequest int             `yaml:"postDelayPerRequest"`
	WorkerCount         int             `yaml:"workerCount"`
	BaseURL             string          `yaml:"baseURL"`
	Targets             []CrawlerTarget `yaml:"targets"`
}

type CrawlerTarget struct {
	Name        string
	Service     string
	OperationID string `yaml:"operationId"`
	Properties  []TargetProperty
}

type TargetProperty struct {
	Name  string
	Type  string
	XPath string `yaml:"xpath"`
}
