package config

type Deployment struct {
	Host              string
	Email             string
	SecretName        string `yaml:"secretName"`
	DockerHubUsername string `yaml:"dockerHubUsername"`
	DockerHubRepo     string `yaml:"dockerHubRepo"`
	TargetPort        string `yaml:"targetPort"`
}
