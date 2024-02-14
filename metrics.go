package main

import (
	"context"
	"fmt"
)

type MetricsService struct {
	next PriceFetcher
}

func (s *MetricsService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	fmt.Println("Pushing to prometheus")
	// your metrics storage. Push to prometheus(gauge, counters)
	return s.FetchPrice(ctx, ticker)
}
