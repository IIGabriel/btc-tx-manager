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

var defaults = map[string]interface{}{
	constants.DebugMode:    true,
	constants.Port:         "8080",
	constants.PostgresHost: "5.161.56.166",
	constants.PostgresDb:   "verbeux",
	constants.PostgresPort: 5432,
	constants.PostgresUser: "verbeux",
	constants.PostgresPass: "r+w12DbCZkDu9ehqBZ83+x2bVGyxlRiA",
	constants.InfuraApiKey: "41ece36efb6d4a36ba997e43494216fd",
}

func EnvInt(key string) int64 {
	valueStr := os.Getenv(key)
	value, err := strconv.ParseInt(valueStr, 10, 64)
	if valueStr == "" || err != nil {
		if err != nil && valueStr != "" {
			zap.L().Panic("failed to convert integer from env", zap.Error(err))
		}

		valueInterface, ok := defaults[key]
		if !ok {
			zap.L().Fatal("missing env", zap.Error(ErrMissingEnv), zap.String("env", key))
		}

		valueInt, ok := valueInterface.(int)
		if !ok {
			zap.L().Fatal("wrong type", zap.Error(ErrEnvType), zap.String("env", key))
		}
		value = int64(valueInt)
	}

	return value
}

func EnvBool(key string) bool {
	env := os.Getenv(key)
	value := strings.EqualFold(env, "true")
	if env == "" {
		valueBool, ok := defaults[key]
		if !ok {
			zap.L().Fatal("missing env", zap.Error(ErrMissingEnv), zap.String("env", key))
		}

		value, ok = valueBool.(bool)
		if !ok {
			zap.L().Fatal("wrong type", zap.Error(ErrEnvType), zap.String("env", key))
		}
	}

	return value
}

func EnvString(key string) string {
	env := os.Getenv(key)
	if env == "" {
		valueInterface, ok := defaults[key]
		if !ok {
			zap.L().Fatal("missing env", zap.Error(ErrMissingEnv), zap.String("env", key))
		}

		env, ok = valueInterface.(string)
		if !ok {
			zap.L().Fatal("wrong type", zap.Error(ErrEnvType), zap.String("env", key))
		}

	}

	return env
}
