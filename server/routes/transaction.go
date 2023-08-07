package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/IIGabriel/btc-tx-manager/constants"
	"github.com/IIGabriel/btc-tx-manager/server/controllers"
)

func Transactions(app *fiber.App) {
	group := app.Group(constants.RouteTransaction)
	controller := controllers.NewTransaction()

	group.Get(":hash", controller.GetOne)
	group.Post(":hash", controller.Create)
	group.Delete(":hash", controller.Delete)
}
