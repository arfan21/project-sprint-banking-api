package transactionsvc

import (
	"context"
	"fmt"

	"github.com/arfan21/project-sprint-banking-api/internal/entity"
	"github.com/arfan21/project-sprint-banking-api/internal/model"
	transactionrepo "github.com/arfan21/project-sprint-banking-api/internal/transaction/repository"
	"github.com/arfan21/project-sprint-banking-api/pkg/pkgutil"
	"github.com/arfan21/project-sprint-banking-api/pkg/validation"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gopkg.in/guregu/null.v4"
)

type Service struct {
	repo transactionrepo.RepositoryStdLib
}

func New(repo transactionrepo.RepositoryStdLib) Service {
	return Service{repo: repo}
}

func (s Service) AddBalance(ctx context.Context, req model.TransactionAddBalanceRequest) (err error) {
	err = validation.Validate(req)
	if err != nil {
		err = fmt.Errorf("transaction.service.AddBalance: failed to validate request: %w", err)
		return
	}
	reqBalanceDecimal := decimal.NewFromFloat(req.AddedBalance)

	id, err := uuid.NewV7()
	if err != nil {
		err = fmt.Errorf("user.service.Create: failed to generate product id: %w", err)
		return err
	}

	balanceRecord := entity.Balance{
		UserID:   req.UserID,
		Currency: req.Currency,
		Balance:  reqBalanceDecimal,
	}

	transactionRecord := entity.Transaction{
		ID:                id.String(),
		UserID:            req.UserID,
		Amount:            reqBalanceDecimal,
		Currency:          req.Currency,
		BankName:          req.SenderBankName,
		BankAccountNumber: req.SenderBankAccountNumber,
		TransferProofImg:  null.StringFrom(req.TransferProofImg),
	}

	err = s.repo.UpsertBalanceWithRecordTx(ctx, balanceRecord, transactionRecord)
	if err != nil {
		err = fmt.Errorf("transaction.service.AddBalance: failed to record transaction: %w", err)
		return err
	}

	return nil
}

func (s Service) TransferBalance(ctx context.Context, req model.TransactionTransferBalanceRequest) (err error) {
	err = validation.Validate(req)
	if err != nil {
		err = fmt.Errorf("transaction.service.TransferBalance: failed to validate request: %w", err)
		return
	}

	reqBalanceDecimal := decimal.NewFromFloat(req.Balances).Neg()

	id, err := uuid.NewV7()
	if err != nil {
		err = fmt.Errorf("user.service.Create: failed to generate product id: %w", err)
		return err
	}

	balanceRecord := entity.Balance{
		UserID:   req.UserID,
		Currency: req.FromCurrency,
		Balance:  reqBalanceDecimal,
	}

	transactionRecord := entity.Transaction{
		ID:                id.String(),
		UserID:            req.UserID,
		Amount:            reqBalanceDecimal,
		Currency:          req.FromCurrency,
		BankName:          req.RecipientBankName,
		BankAccountNumber: req.RecipientBankAccountNumber,
	}

	err = s.repo.UpdateBalanceWithREcordTx(ctx, balanceRecord, transactionRecord)
	if err != nil {
		err = fmt.Errorf("transaction.service.AddBalance: failed to record transaction: %w", err)
		return err
	}

	return nil
}

func (s Service) GetListByUserID(ctx context.Context, req model.TransactionGetListRequest) (res []model.TransactionGetResponse, meta pkgutil.MetaResponse, err error) {
	if req.Limit == 0 {
		req.Limit = 5
	}

	resDB, err := s.repo.GetList(ctx, req)
	if err != nil {
		err = fmt.Errorf("transaction.service.GetListByUserID: failed to get list record transaction: %w", err)
		return
	}

	res = make([]model.TransactionGetResponse, len(resDB))

	for i, val := range resDB {
		res[i] = model.TransactionGetResponse{
			TransactionID:    val.ID,
			Balance:          val.Amount.InexactFloat64(),
			Currency:         val.Currency,
			TransferProofImg: val.TransferProofImg.ValueOrZero(),
			CreatedAt:        val.CreatedAt.Unix(),
			Source: model.TransactionSourceResponse{
				BankAcccountNumber: val.BankAccountNumber,
				BankName:           val.BankName,
			},
		}

		meta.Total = val.Total
	}

	meta.Limit = req.Limit
	meta.Offset = req.Offset

	return
}
