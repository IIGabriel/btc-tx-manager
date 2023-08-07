package utils

import (
	"os"
	"strconv"
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
	constants.MongoEnvKey:      "mongodb://localhost:27017",
	constants.MongoDataBaseKey: "verbeux",
	constants.InfuraApiKey:     "41ece36efb6d4a36ba997e43494216fd",
}

func EnvInt(key constants.EnvKey) int64 {
	valueStr := os.Getenv(string(key))
	value, err := strconv.ParseInt(valueStr, 10, 64)
	if valueStr == "" || err != nil {
		if err != nil && valueStr != "" {
			zap.L().Panic("failed to convert integer from env", zap.Error(err))
		}

		valueInterface, ok := defaults[key]
		if !ok {
			zap.L().Fatal("missing env", zap.Error(ErrMissingEnv), zap.String("env", string(key)))
		}

		valueInt, ok := valueInterface.(int)
		if !ok {
			zap.L().Fatal("wrong type", zap.Error(ErrEnvType), zap.String("env", string(key)))
		}
		value = int64(valueInt)
	}

	return value
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
