package model

type TransactionGetListRequest struct {
	UserID string `json:"-"`
	Limit  int    `query:"limit"`
	Offset int    `query:"offset"`
}

//	{
//		"senderBankAccountNumber": "" // not null, minLength 5 maxLength 30
//		"senderBankName": "" // not null, minLength 5 maxLength 30
//		"addedBalance": 1, // not null, can't be negative
//		"currency":"USD", // not null, should be a valid ISO 4217 currency codes
//		"transferProofImg": "", // not null, should be an image url
//	}
type TransactionAddBalanceRequest struct {
	SenderBankAccountNumber string  `json:"senderBankAccountNumber" validate:"required,min=5,max=30"`
	SenderBankName          string  `json:"senderBankName" validate:"required,min=5,max=30"`
	AddedBalance            float64 `json:"addedBalance" validate:"required,gte=0"`
	Currency                string  `json:"currency" validate:"required,iso4217"`
	TransferProofImg        string  `json:"transferProofImg" validate:"required,customurl"`
	UserID                  string  `json:"-"`
}

//	{
//		"recipientBankAccountNumber": "" // not null, minLength 5 maxLength 30
//		"recipientBankName": "" // not null, minLength 5 maxLength 30
//		"fromCurrency":"", // not null, should be a valid ISO 4217 currency codes
//		"balances":1,  // not null
//	}
type TransactionTransferBalanceRequest struct {
	RecipientBankAccountNumber string  `json:"recipientBankAccountNumber" validate:"required,min=5,max=30"`
	RecipientBankName          string  `json:"recipientBankName" validate:"required,min=5,max=30"`
	FromCurrency               string  `json:"fromCurrency" validate:"required,iso4217"`
	Balances                   float64 `json:"balances" validate:"required,gte=0"`
	UserID                     string  `json:"-"`
}

// {
// 	"transactionId":"",
// 	"balance":1,
// 	// ☝️ can be negative if there's transaction
// 	// or positive it there's top up balance
// 	"currency":"",
// 	"transferProofImg": "",
// 	"createdAt": 1582605077000
// 	// ☝️ should in
// 	// UNIX Timestamp with milliseconds
// 	// (ex: 1582605077000)
// 	"source": {
// 		// if it's a transaction, fill with recipient information
// 		// if it's a top up balance, fill it with sender information
// 		"bankAccountNumber":"",
// 		"bankName":""
// 	}
// }

type TransactionGetResponse struct {
	TransactionID    string                    `json:"transactionId"`
	Balance          float64                   `json:"balance"`
	Currency         string                    `json:"currency"`
	TransferProofImg string                    `json:"transferProofImg"`
	CreatedAt        int64                     `json:"createdAt"`
	Source           TransactionSourceResponse `json:"source"`
}

type TransactionSourceResponse struct {
	BankAcccountNumber string `json:"bankAccountNumber"`
	BankName           string `json:"bankName"`
}
