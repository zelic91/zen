package config

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
