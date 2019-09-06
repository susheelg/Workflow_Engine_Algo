package model

var (
	//Config - Global obj for accessing application config
	Config AppConfig
)

//AppConfig - Config loaded from config.toml
type AppConfig struct {
	AppPort       string
	LogDir        string
	MongoURL      string
	MongoPort     string
	MongoDBName   string
	MongoUsername string
	MongoPass     string
	HostName      string
}

// GetConfigFilePath return config filepath
func GetConfigFilePath() string {
	return "config/config.toml"
}
