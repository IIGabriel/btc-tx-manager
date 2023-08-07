package services

import (
	"encoding/json"
	"fmt"

	"github.com/IIGabriel/eth-tx-manager/models"
	"github.com/IIGabriel/eth-tx-manager/utils"
)

func GetTransactionByHash(hash string) (*models.Transaction, error) {

	body, err := utils.Get(fmt.Sprintf("https://blockchain.info/rawtx/%s", hash))
	if err != nil {
		return nil, err
	}
	var transaction models.Transaction
	if err := json.Unmarshal(body, &transaction); err != nil {
		return nil, err
	}

	return &transaction, nil
}
