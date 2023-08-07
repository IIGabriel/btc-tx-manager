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
	var transaction models.RawTransaction
	if err := json.Unmarshal(body, &transaction); err != nil {
		return nil, err
	}

	tx := transaction.ConvertToTransaction()
	lastBlockHeight, err := GetLastBlockHeight()
	if err != nil {
		return nil, err
	}
	if transaction.BlockHeight != nil {
		tx.Confirmations = (*lastBlockHeight - *transaction.BlockHeight) + 1
	} else {
		tx.Confirmations = 0
	}
	return tx, nil
}

func GetLastBlockHeight() (*int64, error) {
	body, err := utils.Get(fmt.Sprintf("%s/latestblock", constants.BlockChainURL))
	if err != nil {
		return nil, err
	}
	var height struct {
		Height *int64 `json:"height"`
	}
	if err := json.Unmarshal(body, &height); err != nil {
		return nil, err
	}

	return height.Height, nil
}
