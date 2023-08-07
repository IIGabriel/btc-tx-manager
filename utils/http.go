package utils

import (
	"net/http"

	"github.com/IIGabriel/eth-tx-manager/constants"
)

var client = &http.Client{
	Timeout: constants.HTTPTimeout,
}

func Get(url string) error {
	res, err := client.Get(url)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return ErrNotReturnedOk
	}

	return nil
}
