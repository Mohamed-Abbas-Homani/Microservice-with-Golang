package types

type PriceResponse struct {
	Ticker string  `json:"ticker"`
	Price  float64 `json:"price"`
}

type APIResponse struct {
	Success   bool               `json:"success"`
	Timestamp int64              `json:"timestamp"`
	Date      string             `json:"date"`
	Base      string             `json:"base"`
	Rates     map[string]float32 `json:"rates"`
}
