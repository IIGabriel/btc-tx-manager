package services

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"

	"github.com/IIGabriel/eth-tx-manager/constants"
	"github.com/IIGabriel/eth-tx-manager/utils"
)

var postgresInstance *gorm.DB

func Postgres() *gorm.DB {
	if postgresInstance == nil {
		dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			utils.EnvString(constants.PostgresHost),
			utils.EnvString(constants.PostgresUser),
			utils.EnvString(constants.PostgresPass),
			utils.EnvString(constants.PostgresDb),
			utils.EnvInt(constants.PostgresPort),
		)

		db, err := gorm.Open(postgres.Open(dns))
		if err != nil {
			zap.L().Panic("failed to connect to postgres", zap.Error(err))
		}

		d, err := db.DB()
		if err != nil {
			zap.L().Panic("failed to get DB", zap.Error(err))
		}

		if err = d.Ping(); err != nil {
			zap.L().Panic("failed to ping DB", zap.Error(err))
		}

		postgresInstance = db
	}

	return postgresInstance
}
