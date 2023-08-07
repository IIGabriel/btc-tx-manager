package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/IIGabriel/btc-tx-manager/constants"
	"github.com/IIGabriel/btc-tx-manager/models"
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

func HTTPSuccess(ctx *fiber.Ctx, data interface{}, pageData ...uint64) error {
	gotPage := uint64(0)
	gotPerPage := uint64(0)
	total := uint64(0)
	if len(pageData) > 0 {
		gotPage = pageData[0]
	}

	if len(pageData) > 1 {
		gotPerPage = pageData[1]
	}

	if len(pageData) > 2 {
		total = pageData[2]
	}

	dataReflect := reflect.ValueOf(data)
	isSlice := dataReflect.Kind() == reflect.Slice

	if data == nil || (isSlice && dataReflect.Len() < 1) {
		data = []string{}
	}

	return ctx.Status(http.StatusOK).JSON(&models.HTTPResponse{
		Data:    data,
		Page:    gotPage,
		PerPage: gotPerPage,
		Total:   total,
	})
}

func HTTPFail(ctx *fiber.Ctx, code int, err error, message string) error {
	errJson, _ := json.Marshal(err)

	result := &models.HTTPErrorResponse{
		Error:   errJson,
		Message: message,
	}

	if err != nil {
		result.ErrorMessage = err.Error()
	}

	return ctx.Status(code).JSON(result)
}

func FilterRangeDate(ctx *fiber.Ctx) (start time.Time, end time.Time, err error) {
	startDate := ctx.Query(constants.QueryInitialDate)
	endDate := ctx.Query(constants.QueryFinalDate)
	if startDate != "" {
		start, err = time.Parse(time.RFC3339, startDate)
		if err != nil {
			return time.Time{}, time.Time{}, HTTPFail(ctx, http.StatusBadRequest, err, fmt.Sprintf("invalid %s format", constants.QueryInitialDate))
		}
	}

	if endDate != "" {
		end, err = time.Parse(time.RFC3339, endDate)
		if err != nil {
			return time.Time{}, time.Time{}, HTTPFail(ctx, http.StatusBadRequest, err, fmt.Sprintf("invalid %s format", constants.QueryFinalDate))
		}
	}

	return start, end, nil

}
