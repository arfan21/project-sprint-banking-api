package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Balance struct {
	UserID    string          `json:"userId"`
	Currency  string          `json:"currency"`
	Balance   decimal.Decimal `json:"balance"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
}

func (Balance) TableName() string {
	return "balances"
}
