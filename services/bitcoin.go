package services

import (
	"encoding/json"
	"fmt"

	"github.com/IIGabriel/btc-tx-manager/constants"
	"github.com/IIGabriel/btc-tx-manager/models"
	"github.com/IIGabriel/btc-tx-manager/utils"
)

func GetTransactionByHash(hash string) (*models.Transaction, error) {

	body, err := utils.Get(fmt.Sprintf("%s/rawtx/%s", constants.BlockChainURL, hash))
	if err != nil {
		return nil, err
	}
	var transaction models.Transaction
	if err := json.Unmarshal(body, &transaction); err != nil {
		return nil, err
	}

	return &transaction, nil
}
