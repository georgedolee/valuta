package apiclient

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/georgedolee/valuta/internal/model"
)

var (
	baseUrl = "https://nbg.gov.ge/gw/api/ct/monetarypolicy/currencies"
	client  = &http.Client{Timeout: 10 * time.Second}
)

func FetchConversionApi(amount float64, from, to string) (float64, error) {
	url, err := url.Parse(baseUrl)
	if err != nil {
		return 0, fmt.Errorf("invalid base URL: %w", err)
	}

	url.Path += "/calculator"
	setQueries(url, map[string]string{
		"codeFrom": from,
		"codeTo":   to,
		"quantity": fmt.Sprint(amount),
	})

	res, err := fetch(url.String())
	if err != nil {
		return 0, err
	}

	var result float64
	err = decodeJSONBody(res.Body, &result)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func FetchRatesApi(currencies []string) ([]model.Currencies, error) {
	url, err := url.Parse(baseUrl)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	url.Path += "/en/json/"

	q := url.Query()
	for _, currency := range currencies {
		q.Add("currencies", currency)
	}
	url.RawQuery = q.Encode()

	res, err := fetch(url.String())
	if err != nil {
		return nil, err
	}

	var response model.RatesResponse
	err = decodeJSONBody(res.Body, &response)
	if err != nil {
		return nil, err
	}

	return response[0].Currencies, nil
}
