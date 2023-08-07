package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/IIGabriel/btc-tx-manager/constants"
	"github.com/IIGabriel/btc-tx-manager/server/controllers"
)

func Transactions(app *fiber.App) {
	group := app.Group(constants.RouteTransaction)
	controller := controllers.NewTransaction()

	group.Get(":hashId", controller.GetOne)
	group.Post(":hash", controller.Create)
	group.Delete(":hashId", controller.Delete)
	group.Get("", controller.GetMany)
	group.Put(":hashId", controller.Update)
	group.Put("/blockchain/:hash", controller.Custom(constants.CustomUpdateByBlockchain))
}
