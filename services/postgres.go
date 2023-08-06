package services

import (
	"fmt"

	"github.com/verbeux-ai/verbeux-admin/constants"
	"github.com/verbeux-ai/verbeux-admin/utils"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var postgresInstance *gorm.DB

func Postgres() *gorm.DB {
	if postgresInstance == nil {
		connectionURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			utils.EnvString(constants.PostgresHost),
			utils.EnvString(constants.PostgresUser),
			utils.EnvString(constants.PostgresPass),
			utils.EnvString(constants.PostgresDb),
			utils.EnvInt(constants.PostgresPort),
		)

		db, err := gorm.Open(postgres.Open(connectionURL))
		if err != nil {
			zap.L().Panic("failed to connect to postgres", zap.Error(err))
		}

		d, err := db.DB()
		if err != nil {
			zap.L().Panic("failed to get DB", zap.Error(err))
		}

		if err := d.Ping(); err != nil {
			zap.L().Panic("failed to ping DB", zap.Error(err))
		}

		postgresInstance = db
	}

	return postgresInstance
}
