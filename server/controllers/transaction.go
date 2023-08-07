package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

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

func (t transaction) GetMany(ctx *fiber.Ctx) error {
	var pagination interfaces.Pagination
	var sort interfaces.Sort

	if err := ctx.QueryParser(&pagination); err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "failed to parse params")
	}
	if err := ctx.QueryParser(&sort); err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "failed to parse params")
	}

	transactions, err := t.repoM.Find(bson.D{}, interfaces.MongoFilter{Pagination: pagination, Sort: sort})
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusInternalServerError, err, "failed to get transactions from db")
	}

	count, err := t.repoM.Count(bson.D{})
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusInternalServerError, err, "failed to count transactions")
	}

	return utils.HTTPSuccess(ctx, transactions, uint64(pagination.Page), uint64(pagination.PerPage), uint64(count))
}

func (t transaction) GetOne(ctx *fiber.Ctx) error {
	hashId := ctx.Params("hashId")
	if hashId == "" {
		return utils.HTTPFail(ctx, http.StatusBadRequest, nil, ErrRequired("hash or id").Error())
	}

	filter := bson.D{{"transaction_hash", hashId}}
	if pId, err := primitive.ObjectIDFromHex(hashId); err == nil {
		filter = bson.D{{"_id", pId}}
	}

	tx, err := t.repoM.FindOne(filter)
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

	id, err := t.repoM.Create(transactionByHash)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusInternalServerError, err, "failed to create transaction")
	}
	transactionByHash.ID = id
	return ctx.Status(http.StatusCreated).JSON(transactionByHash)
}

func (t transaction) Update(ctx *fiber.Ctx) error {
	hashId := ctx.Params("hashId")
	if hashId == "" {
		return utils.HTTPFail(ctx, http.StatusBadRequest, nil, ErrRequired("hash").Error())
	}

	var updateTx models.Transaction
	if err := ctx.BodyParser(&updateTx); err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "failed to parse request body")
	}

	filter := bson.D{{"transaction_hash", hashId}}
	updateFields := bson.M{}

	// Check each field individually
	if !updateTx.Time.IsZero() {
		updateFields["time"] = updateTx.Time
	}
	if updateTx.Fee != 0 {
		updateFields["fee"] = updateTx.Fee
	}
	if len(updateTx.Inputs) > 0 {
		updateFields["inputs"] = updateTx.Inputs
	}
	if len(updateTx.Outputs) > 0 {
		updateFields["outputs"] = updateTx.Outputs
	}

	if len(updateFields) == 0 {
		return utils.HTTPFail(ctx, http.StatusBadRequest, nil, "no fields provided for update")
	}

	err := t.repoM.Update(filter, bson.D{{"$set", updateFields}})
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusInternalServerError, err, "failed to update the transaction in the database")
	}

	return ctx.SendStatus(http.StatusOK)
}

func (t transaction) Delete(ctx *fiber.Ctx) error {
	hashId := ctx.Params("hashId")
	if hashId == "" {
		return utils.HTTPFail(ctx, http.StatusBadRequest, nil, ErrRequired("hash or id").Error())
	}
	filter := bson.D{{"transaction_hash", hashId}}
	if pId, err := primitive.ObjectIDFromHex(hashId); err == nil {
		filter = bson.D{{"_id", pId}}
	}
	if err := t.repoM.Delete(filter); err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "failed to delete transaction")
	}

	return utils.HTTPSuccess(ctx, "transaction deleted")
}

func (t transaction) Custom(route string) func(ctx *fiber.Ctx) error {
	return nil
}
