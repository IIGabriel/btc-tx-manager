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

// @Summary Retrieve multiple transactions
// @Description Retrieve transactions with optional filters.
// @Accept  json
// @Produce  json
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Param input_address query string false "Input Address"
// @Param output_address query string false "Output Address"
// @Param sort_field query string false "Field for sorting results"
// @Param asc query boolean false "Sort in ascending order"
// @Param page query int false "Page number for pagination"
// @Param perPage query int false "Number of results per page"
// @Success 200 {object} models.Transaction
// @Router /transactions [get]
func (t transaction) GetMany(ctx *fiber.Ctx) error {
	var mongoFilter interfaces.MongoFilter
	if err := ctx.QueryParser(&mongoFilter); err != nil {
		return err
	}
	filter := bson.D{}

	start, end, err := utils.FilterRangeDate(ctx)
	if err != nil {
		return err
	}
	if !start.IsZero() && !end.IsZero() {
		filter = append(filter, bson.E{Key: "time", Value: bson.M{"$gte": start, "$lte": end}})
	}
	var filters = map[string]string{
		"inputs.address":  ctx.Query("input_address"),
		"outputs.address": ctx.Query("output_address"),
	}
	filter = utils.QueryFiltersString(filter, filters)
	transactions, err := t.repoM.Find(filter, mongoFilter)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusInternalServerError, err, "failed to get transactions from db")
	}

	count, err := t.repoM.Count(filter)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusInternalServerError, err, "failed to count transactions")
	}

	return utils.HTTPSuccess(ctx, transactions, uint64(mongoFilter.Page), uint64(mongoFilter.PerPage), uint64(count))
}

// @Summary Retrieve a single transaction
// @Description Retrieve a specific transaction by hash or id.
// @Accept  json
// @Produce  json
// @Param hashId path string true "Transaction Hash or ID"
// @Success 200 {object} models.Transaction
// @Router /transactions/{hashId} [get]
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

// @Summary Add a new transaction
// @Description Create a new transaction based on provided hash.
// @Accept  json
// @Produce  json
// @Param hash path string true "Transaction Hash"
// @Success 201 {object} models.Transaction
// @Router /transactions/{hash} [post]
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

// @Summary Update a transaction
// @Description Update specific fields of a transaction.
// @Accept  json
// @Produce  json
// @Param hashId path string true "Transaction Hash or ID"
// @Param transaction body models.TransactionToUpdate true "Transaction object"
// @Success 200 {object} models.TransactionToUpdate
// @Router /transactions/{hashId} [put]
func (t transaction) Update(ctx *fiber.Ctx) error {
	hashId := ctx.Params("hashId")
	if hashId == "" {
		return utils.HTTPFail(ctx, http.StatusBadRequest, nil, ErrRequired("hash").Error())
	}

	var updateTx models.TransactionToUpdate
	if err := ctx.BodyParser(&updateTx); err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "failed to parse request body")
	}

	filter := bson.D{{"transaction_hash", hashId}}
	if pId, err := primitive.ObjectIDFromHex(hashId); err == nil {
		filter = bson.D{{"_id", pId}}
	}
	updateFields := bson.M{}

	mapFieldsValue := map[string]struct {
		fieldValue any
		isNotEmpty bool
	}{
		"time":          {updateTx.Time, !updateTx.Time.IsZero()},
		"fee":           {updateTx.Fee, updateTx.Fee != 0},
		"inputs":        {updateTx.Inputs, len(updateTx.Inputs) != 0},
		"outputs":       {updateTx.Outputs, len(updateTx.Outputs) != 0},
		"confirmations": {updateTx.Confirmations, updateTx.Confirmations != 0},
		"block_height":  {updateTx.BlockHeight, updateTx.Confirmations != 0},
		"block_index":   {updateTx.BlockIndex, updateTx.Confirmations != 0},
	}

	for key, value := range mapFieldsValue {
		if value.isNotEmpty {
			updateFields[key] = value.fieldValue
		}
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

// @Summary Delete a transaction
// @Description Remove a transaction based on provided hash or id.
// @Accept  json
// @Produce  json
// @Param hashId path string true "Transaction Hash or ID"
// @Success 200 {string} string "transaction deleted"
// @Router /transactions/{hashId} [delete]
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

	return ctx.SendStatus(http.StatusOK)
}

func (t transaction) Custom(route string) func(ctx *fiber.Ctx) error {
	switch route {
	case constants.CustomUpdateByBlockchain:
		return t.UpdateByBlockchain
	}
	return nil
}

// @Summary Update a transaction using blockchain data
// @Description Fetch data from the blockchain and update specific fields of a transaction.
// @Accept  json
// @Produce  json
// @Param hash path string true "Transaction Hash"
// @Success 200
// @Router /transactions/blockchain/{hash} [put]
func (t transaction) UpdateByBlockchain(ctx *fiber.Ctx) error {
	hash := ctx.Params("hash")
	if hash == "" {
		return utils.HTTPFail(ctx, http.StatusBadRequest, nil, ErrRequired("hash").Error())
	}

	transactionByHash, err := services.GetTransactionByHash(hash)
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusBadRequest, err, "failed to get transaction from blockchain")
	}

	filter := bson.D{{"transaction_hash", hash}}

	updateFields := bson.M{
		"time":          transactionByHash.Time,
		"fee":           transactionByHash.Fee,
		"inputs":        transactionByHash.Inputs,
		"outputs":       transactionByHash.Outputs,
		"confirmations": transactionByHash.Confirmations,
		"block_height":  transactionByHash.BlockHeight,
		"block_index":   transactionByHash.BlockIndex,
	}

	err = t.repoM.Update(filter, bson.D{{"$set", updateFields}})
	if err != nil {
		return utils.HTTPFail(ctx, http.StatusInternalServerError, err, "failed to update the transaction in the database using blockchain data")
	}

	return ctx.SendStatus(http.StatusOK)
}
