package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"go.uber.org/zap"

	_ "github.com/IIGabriel/btc-tx-manager/docs"

	"github.com/IIGabriel/btc-tx-manager/constants"
	"github.com/IIGabriel/btc-tx-manager/server/routes"
	"github.com/IIGabriel/btc-tx-manager/services"
	"github.com/IIGabriel/btc-tx-manager/utils"
)

// @title BTC Transaction Manager
// @description API for managing Bitcoin transactions.
// @version 1.0
// @host localhost:8080
// @BasePath /
func main() {

	utils.InitLogger()
	services.SetupMongo()

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	routes.Load(app)

	port := ":" + utils.EnvString(constants.Port)
	zap.L().Info("starting server...", zap.String("port", port))
	if err := app.Listen(port); err != nil {
		zap.L().Fatal("failed to start server", zap.Error(err))
	}
}
