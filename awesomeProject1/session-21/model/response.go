package model

type RateResponse struct {
	Id       int     `json:"id"`
	Currency string  `json:"currency"`
	Price    float64 `json:"price"`
}
