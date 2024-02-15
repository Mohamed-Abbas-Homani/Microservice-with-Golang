package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Mohamed-Abbas-Homani/microservice/types"
)

// PriceFetcher is an interface that can fetch a price.
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

// priceFetcher implements the PriceFetcher interface.
type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	apiKey := "gfnYJdIavCcbaPYCANqwYTXbEpiFf8L4"
	baseCurrency := "USD"
	targetCurrency := ticker

	url := fmt.Sprintf("https://api.apilayer.com/exchangerates_data/latest?symbols=%s&base=%s", targetCurrency, baseCurrency)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("apikey", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	apiResp := new(types.APIResponse)
	if err := json.NewDecoder(resp.Body).Decode(apiResp); err != nil {
		return 0, fmt.Errorf("failed to decode response: %w", err)
	}

	rate, ok := apiResp.Rates[ticker]
	if !ok {
		return 0, fmt.Errorf("currency %s not found in response", ticker)
	}

	return float64(rate), nil
}

// var priceMocker = map[string]float64{
// 	"BTC": 20000.0,
// 	"ETH": 200.0,
// 	"GG":  1000000.0,
// }

// func MocPriceFetcher(ctx context.Context, ticker string) (float64, error) {
// 	price, ok := priceMocker[ticker]
// 	if !ok {
// 		return price, fmt.Errorf("the given ticker %s is not supported", ticker)
// 	}

// 	return price, nil
// }
