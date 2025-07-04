package model

type Currencies struct {
	Code     string  `json:"code"`
	Quantity float64 `json:"quantity"`
	Rate     float64 `json:"rate"`
}

type RatesResponse []struct {
	Currencies []Currencies `json:"currencies"`
}
