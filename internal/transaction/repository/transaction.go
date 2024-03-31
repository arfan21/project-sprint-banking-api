package transactionrepo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/arfan21/project-sprint-banking-api/internal/entity"
	"github.com/arfan21/project-sprint-banking-api/internal/model"
	dbpostgres "github.com/arfan21/project-sprint-banking-api/pkg/db/postgres"
)

type RepositoryStdLib struct {
	db    dbpostgres.QueryerStdLib
	rawDb *sql.DB
}

func NewStdLib(db *sql.DB) RepositoryStdLib {
	return RepositoryStdLib{db: db, rawDb: db}
}

func (r RepositoryStdLib) Begin(ctx context.Context) (tx *sql.Tx, err error) {
	return r.rawDb.BeginTx(ctx, nil)
}

func (r RepositoryStdLib) WithTx(tx *sql.Tx) *RepositoryStdLib {
	r.db = tx
	return &r
}

func (r RepositoryStdLib) GetList(ctx context.Context, filter model.TransactionGetListRequest) (data []entity.Transaction, err error) {
	query := `
		SELECT COUNT(id) OVER() as total,  id, userId, amount, currency, bankName, bankAccountNumber, createdAt, transferProofImg
		FROM transactions
		WHERE userId = $1
		ORDER BY id DESC
		LIMIT $2 
		OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, query, filter.UserID, filter.Limit, filter.Offset)
	if err != nil {
		err = fmt.Errorf("transaction.repository.GetList: failed to get list transaction: %w", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction entity.Transaction
		err = rows.Scan(&transaction.Total, &transaction.ID, &transaction.UserID, &transaction.Amount, &transaction.Currency, &transaction.BankName, &transaction.BankAccountNumber, &transaction.CreatedAt, &transaction.TransferProofImg)
		if err != nil {
			err = fmt.Errorf("transaction.repository.GetList: failed to scan transaction: %w", err)
			return nil, err
		}

		data = append(data, transaction)
	}

	return data, nil
}
