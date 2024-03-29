package transactionrepo

import (
	"context"
	"fmt"

	"github.com/arfan21/project-sprint-banking-api/internal/entity"
	"github.com/arfan21/project-sprint-banking-api/pkg/constant"
)

func (r RepositoryStdLib) GetBalanceByUserID(ctx context.Context, userId string) (data []entity.Balance, err error) {
	query := `
		SELECT  balance, currency
		FROM balances
		WHERE userId = $1
		ORDER BY balance DESC
	`

	rows, err := r.db.QueryContext(ctx, query, userId)
	if err != nil {
		err = fmt.Errorf("transaction.repository.GetBalanceByUserID: failed to get balance by user id: %w", err)
		return data, err
	}

	for rows.Next() {
		var balance entity.Balance
		err = rows.Scan(&balance.Balance, &balance.Currency)
		if err != nil {
			return data, err
		}

		data = append(data, balance)
	}

	return data, nil
}

func (r RepositoryStdLib) UpsertBalanceWithRecordTx(ctx context.Context, data entity.Balance, dataTx entity.Transaction) (err error) {
	query := `
		WITH
			updated_row AS (
				INSERT INTO
					balances (userId, currency, balance)
				VALUES
					($1, $2, $3) ON CONFLICT (userId, currency)
				DO
					UPDATE
						SET
						balance = balances.balance + $3
					WHERE
						balances.balance + $3 >= 0
						RETURNING *
		)
		-- insert to transaction
		INSERT INTO
		transactions (
			id,
			userId,
			currency,
			amount,
			bankName,
			bankAccountNumber,
			transferProofImg
		)
		SELECT
			$7 AS id,
			userId,
			currency,
			$3 AS amount,
			$4 AS bankName,
			$5 AS bankAccountNumber,
			$6 AS transferProofImg
		FROM updated_row
	`

	_, err = r.db.ExecContext(ctx, query,
		data.UserID, data.Currency, data.Balance,
		dataTx.BankName, dataTx.BankAccountNumber, dataTx.TransferProofImg,
		dataTx.ID,
	)

	if err != nil {
		err = fmt.Errorf("transaction.repository.InsertOrUpdateBalanceWithRecordTx: failed to insert or update balance with record transaction: %w", err)
		return err
	}

	return nil
}

func (r RepositoryStdLib) UpdateBalanceWithREcordTx(ctx context.Context, data entity.Balance, dataTx entity.Transaction) (err error) {
	query := `
		WITH
			updated_row AS (
			UPDATE 
				balances
			SET
				balance = balances.balance + $3
			WHERE
				balances.userId = $1 
				AND balances.currency = $2 
				AND balances.balance + $3 >= 0
			RETURNING *
		)
		-- insert to transaction
		INSERT INTO
		transactions (
			id,
			userId,
			currency,
			amount,
			bankName,
			bankAccountNumber,
			transferProofImg
		)
		SELECT
			$7 AS id,
			userId,
			currency,
			$3 AS amount,
			$4 AS bankName,
			$5 AS bankAccountNumber,
			$6 AS transferProofImg
		FROM updated_row
	`

	result, err := r.db.ExecContext(ctx, query,
		data.UserID, data.Currency, data.Balance,
		dataTx.BankName, dataTx.BankAccountNumber, dataTx.TransferProofImg,
		dataTx.ID,
	)

	if err != nil {
		err = fmt.Errorf("transaction.repository.InsertOrUpdateBalanceWithRecordTx: failed to insert or update balance with record transaction: %w", err)
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		err = fmt.Errorf("transaction.repository.InsertOrUpdateBalanceWithRecordTx: failed to get affected rows: %w", err)
		return
	}

	if affected == 0 {
		err = fmt.Errorf("transaction.repository.InsertOrUpdateBalanceWithRecordTx: failed to insert or update balance with record transaction: %w", constant.ErrBalanceInsufficient)
		return err
	}

	return nil
}
