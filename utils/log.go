package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/IIGabriel/btc-tx-manager/constants"
)

func createDirectoryIfNotExist() {
	path, err := os.Getwd()
	if err != nil {
		log.Panicln(err)
	}

	if _, err = os.Stat(fmt.Sprintf("%s/logs", path)); os.IsNotExist(err) {
		_ = os.Mkdir("logs", os.ModePerm)
	}
}

func getLogWriter() zapcore.WriteSyncer {
	path, err := os.Getwd()
	if err != nil {
		log.Panicln(err)
	}
	file, err := os.OpenFile(fmt.Sprintf("/%s/logs/%s.txt", path, time.Now().Format("2006-01-02")),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panicln(err)
	}
	return zapcore.AddSync(file)
}

func getLogFormat() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Local().Format("2006-01-02 15:04:05"))
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func InitLogger() {
	logg, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalln(err)
	}
	if !EnvBool(constants.DebugMode) {
		createDirectoryIfNotExist()
		core := zapcore.NewCore(getLogFormat(), getLogWriter(), zapcore.DebugLevel)
		logg = zap.New(core, zap.AddCaller())
	}

	zap.ReplaceGlobals(logg)
}
