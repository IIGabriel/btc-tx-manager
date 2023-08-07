package controllers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/IIGabriel/eth-tx-manager/constants"
	"github.com/IIGabriel/eth-tx-manager/interfaces"
	"github.com/IIGabriel/eth-tx-manager/models"
	"github.com/IIGabriel/eth-tx-manager/server/repository"
)

func NewTransaction() interfaces.Controller {
	return transaction{
		repository.NewMongoObject(models.Transaction{}, constants.CollectionTransactions)}
}

type transaction struct {
	repoM interfaces.RepositoryMongo[models.Transaction]
}

func (t transaction) GetOne(ctx *fiber.Ctx) error {
	return nil

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
	panic("implement me")
}
