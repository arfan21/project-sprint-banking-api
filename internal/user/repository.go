package user

import (
	"context"
	"database/sql"

	"github.com/arfan21/project-sprint-banking-api/internal/entity"
	userrepo "github.com/arfan21/project-sprint-banking-api/internal/user/repository"
)

type RepositoryStdLib interface {
	Begin(ctx context.Context) (tx *sql.Tx, err error)
	WithTx(tx *sql.Tx) *userrepo.RepositoryStdLib

	Create(ctx context.Context, data entity.User) (err error)
	GetByEmail(ctx context.Context, email string) (data entity.User, err error)
}
