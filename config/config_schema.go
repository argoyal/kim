package config

// Configstructure is the structure for the config file and
// maintains all the config options
type Configstructure struct {
	InfrastructureEnvironment    string
	DefaultDatabaseName          string
	DatabaseEngine               string
	DatabaseConfig               map[string]string
	EmptyTemplates               []string
	IncludeTemplateVariableFile  bool
	TemplateVariableFileName     string
	ImageRepositoryProvider      string
	LoadAwsCredentialsFromConfig bool
	LoadAwsCredentialsFromEnv    bool
}
