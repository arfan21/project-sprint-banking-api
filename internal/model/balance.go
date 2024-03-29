package model

type BalanceGetResponse struct {
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}
