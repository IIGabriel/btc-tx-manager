package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"

	"github.com/IIGabriel/eth-tx-manager/constants"
	"github.com/IIGabriel/eth-tx-manager/utils"
)

func main() {

	utils.InitLogger()

	//utils.Migrate(services.Postgres())

	app := fiber.New()
	app.Use(cors.New())
	//routes.Load(app)

	port := ":" + utils.EnvString(constants.Port)
	zap.L().Info("starting server...", zap.String("port", port))
	if err := app.Listen(port); err != nil {
		zap.L().Fatal("failed to start server", zap.Error(err))
	}
}
