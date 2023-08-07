package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/IIGabriel/eth-tx-manager/constants"
	"github.com/IIGabriel/eth-tx-manager/interfaces"
	"github.com/IIGabriel/eth-tx-manager/models"
	"github.com/IIGabriel/eth-tx-manager/services"
)

func NewTransaction() interfaces.Controller {
	return transaction{
		models.NewMongoObject(models.Transaction{}, services.Mongo().Collection(constants.CollectionTransactions))}
}

type transaction struct {
	repoM interfaces.RepositoryMongo[models.Transaction]
}

func (t transaction) GetOne(ctx *fiber.Ctx) error {
	a := ctx.Params("id")
	b := ctx.Get("id")
	c := ctx.Query("id")
	fmt.Println(a, b, c)
	return ctx.Status(200).JSON("oi")

}
func (t transaction) Create(ctx *fiber.Ctx) error {
	return nil
}
func (t transaction) Update(ctx *fiber.Ctx) error {
	return nil
}
func (t transaction) Delete(ctx *fiber.Ctx) error {
	return nil
}

func (t transaction) Custom(route string) func(ctx *fiber.Ctx) error {
	return nil
}
