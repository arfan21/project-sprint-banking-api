package entity

import (
	"time"

	"github.com/shopspring/decimal"
	"gopkg.in/guregu/null.v4"
)

type Transaction struct {
	ID                string          `json:"id"`
	UserID            string          `json:"userId"`
	Amount            decimal.Decimal `json:"amount"`
	Currency          string          `json:"currency"`
	BankName          string          `json:"bankName"`
	BankAccountNumber string          `json:"bankAccountNumber"`
	TransferProofImg  null.String     `json:"transferProofImg"`
	CreatedAt         time.Time       `json:"createdAt"`
	UpdatedAt         time.Time       `json:"updatedAt"`
	Total             int             `json:"total"`
}

func (Transaction) TableName() string {
	return "transactions"
}
