package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/IIGabriel/btc-tx-manager/constants"
	"github.com/IIGabriel/btc-tx-manager/interfaces"
	"github.com/IIGabriel/btc-tx-manager/models"
	"github.com/IIGabriel/btc-tx-manager/services"
	"github.com/IIGabriel/btc-tx-manager/utils"
)

func NewTransaction() interfaces.Controller {
	return transaction{
		models.NewMongoObject(&models.Transaction{}, services.Mongo().Collection(constants.CollectionTransactions))}
}

type transaction struct {
	repoM interfaces.RepositoryMongo[*models.Transaction]
}

func (t transaction) GetOne(ctx *fiber.Ctx) error {
	hash := ctx.Params("hash")
	if hash == "" {
		return utils.HTTPFail(ctx, http.StatusBadRequest, nil, ErrRequired("hash").Error())
	}

	tx, err := t.repoM.FindOne(bson.D{{"hash", hash}})
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusNotFound, err, "failed to get transaction")
	}

	return utils.HTTPSuccess(ctx, tx)

}
func (t transaction) Create(ctx *fiber.Ctx) error {
	hash := ctx.Params("hash")
	if hash == "" {
		return utils.HTTPFail(ctx, http.StatusBadRequest, nil, ErrRequired("hash").Error())
	}

	transactionByHash, err := services.GetTransactionByHash(hash)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "failed to get transaction")
	}

	if err = t.repoM.Create(transactionByHash); err != nil {
		return utils.HTTPFail(ctx, http.StatusInternalServerError, err, "failed to create transaction")
	}

	return ctx.Status(http.StatusCreated).JSON(transactionByHash)
}
func (t transaction) Update(ctx *fiber.Ctx) error {
	return nil
}
func (t transaction) Delete(ctx *fiber.Ctx) error {
	hash := ctx.Params("hash")
	if hash == "" {
		return utils.HTTPFail(ctx, http.StatusBadRequest, nil, ErrRequired("hash").Error())
	}

	if err := t.repoM.Delete(bson.D{{"hash", hash}}); err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "failed to delete transaction")
	}

	return utils.HTTPSuccess(ctx, "transaction deleted")
}

func (t transaction) Custom(route string) func(ctx *fiber.Ctx) error {
	return nil
}
