package services

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"

	"github.com/IIGabriel/eth-tx-manager/constants"
	"github.com/IIGabriel/eth-tx-manager/utils"
)

var ethInstance *ethclient.Client

func Ethereum() *ethclient.Client {
	if ethInstance == nil {
		var err error
		ethInstance, err = ethclient.Dial(fmt.Sprintf("https://mainnet.infura.io/v3/%s", utils.EnvString(constants.InfuraApiKey)))
		if err != nil {
			zap.L().Error("Failed to connect to the Ethereum client", zap.Error(err))
		}

	}

	return ethInstance
}
