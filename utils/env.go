package utils

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"github.com/IIGabriel/eth-tx-manager/constants"
)

func init() {
	_ = godotenv.Load(".env")
}

var defaults = map[constants.EnvKey]interface{}{
	constants.DebugMode:        true,
	constants.Port:             "8080",
	constants.MongoEnvKey:      "",
	constants.MongoDataBaseKey: "",
}

func EnvBool(key constants.EnvKey) bool {
	env := os.Getenv(string(key))
	value := strings.EqualFold(env, "true")
	if env == "" {
		valueBool, ok := defaults[key]
		if !ok {
			zap.L().Fatal("missing env", zap.Error(ErrMissingEnv), zap.String("env", string(key)))
		}

		value, ok = valueBool.(bool)
		if !ok {
			zap.L().Fatal("wrong type", zap.Error(ErrEnvType), zap.String("env", string(key)))
		}
	}

	return value
}

func EnvString(key constants.EnvKey) string {
	env := os.Getenv(string(key))
	if env == "" {
		valueInterface, ok := defaults[key]
		if !ok {
			zap.L().Fatal("missing env", zap.Error(ErrMissingEnv), zap.String("env", string(key)))
		}

		env, ok = valueInterface.(string)
		if !ok {
			zap.L().Fatal("wrong type", zap.Error(ErrEnvType), zap.String("env", string(key)))
		}

	}

	return env
}
