package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/IIGabriel/eth-tx-manager/constants"
	"github.com/IIGabriel/eth-tx-manager/server/controllers"
)

func Transactions(app *fiber.App) {
	group := app.Group(constants.RouteTransaction)
	controller := controllers.NewTransaction()

	group.Get(":id", controller.GetOne)
	group.Post("", controller.Create)
	group.Delete(":id", controller.Delete)
}
