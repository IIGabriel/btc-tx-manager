package constants

type EnvKey string

const (
	MongoEnvKey      EnvKey = "MONGO_URI"
	MongoDataBaseKey EnvKey = "MONGO_DATABASE"
	DebugMode        EnvKey = "DEBUG"
	Port             EnvKey = "PORT"
)
