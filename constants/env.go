package constants

type EnvKey string

const (
	MongoEnvKey      EnvKey = "MONGO_URI"
	MongoDataBaseKey EnvKey = "MONGO_DATABASE"
	InfuraApiKey     EnvKey = "INFURA_API_KEY"
	DebugMode        EnvKey = "DEBUG"
	Port             EnvKey = "PORT"
)
