package main

import (
	"context"
	"fmt"
)

// PriceFetcher is an interface that can fetch a price.
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

// priceFetcher implements the PriceFetcher interface.
type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MocPriceFetcher(ctx, ticker)
}

var priceMocker = map[string]float64{
	"BTC": 20000.0,
	"ETH": 200.0,
	"GG":  1000000.0,
}

func MocPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	price, ok := priceMocker[ticker]
	if !ok {
		return price, fmt.Errorf("the given ticker %s is not supported", ticker)
	}

	return price, nil
}
