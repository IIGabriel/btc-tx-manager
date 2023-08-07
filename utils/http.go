package utils

import (
	"io"
	"net/http"

	"github.com/IIGabriel/eth-tx-manager/constants"
)

var client = &http.Client{
	Timeout: constants.HTTPTimeout,
}

func Get(url string) ([]byte, error) {
	res, err := client.Get(url)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, ErrNotReturnedOk
	}
	return io.ReadAll(res.Body)
}
