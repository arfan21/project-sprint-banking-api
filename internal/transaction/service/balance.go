package transactionsvc

import (
	"context"
	"fmt"

	"github.com/arfan21/project-sprint-banking-api/internal/model"
)

func (s Service) GetBalanceByUserID(ctx context.Context, userId string) (res []model.BalanceGetResponse, err error) {
	data, err := s.repo.GetBalanceByUserID(ctx, userId)
	if err != nil {
		err = fmt.Errorf("transaction.service.GetBalanceByUserID: failed to get balance by user id: %w", err)
		return
	}

	for _, item := range data {
		res = append(res, model.BalanceGetResponse{

			Currency: item.Currency,
			Balance:  item.Balance.InexactFloat64(),
		})
	}

	return
}
