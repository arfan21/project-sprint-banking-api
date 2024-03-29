package userrepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/arfan21/project-sprint-banking-api/internal/entity"
	"github.com/arfan21/project-sprint-banking-api/pkg/constant"
	dbpostgres "github.com/arfan21/project-sprint-banking-api/pkg/db/postgres"
	"github.com/jackc/pgx/v5/pgconn"
)

type RepositoryStdLib struct {
	db    dbpostgres.QueryerStdLib
	rawDb *sql.DB
}

func NewStdLib(db *sql.DB) *RepositoryStdLib {
	return &RepositoryStdLib{db: db, rawDb: db}
}

func (r RepositoryStdLib) Begin(ctx context.Context) (tx *sql.Tx, err error) {
	return r.rawDb.BeginTx(ctx, nil)
}

func (r RepositoryStdLib) WithTx(tx *sql.Tx) *RepositoryStdLib {
	r.db = tx
	return &r
}

func (r RepositoryStdLib) Create(ctx context.Context, data entity.User) (err error) {
	query := `
		INSERT INTO users (id, email, name, password)
		VALUES ($1, $2, $3, $4)
	`

	_, err = r.db.ExecContext(ctx, query, data.ID, data.Email, data.Name, data.Password)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == constant.ErrSQLUniqueViolation {
				err = constant.ErrEmailAlreadyRegistered
			}
		}

		err = fmt.Errorf("user.repository.Create: failed to create user: %w", err)
		return
	}

	return
}

func (r RepositoryStdLib) GetByEmail(ctx context.Context, email string) (data entity.User, err error) {
	query := `
		SELECT id, email, name, password
		FROM users
		WHERE email = $1
	`
	err = r.db.QueryRowContext(ctx, query, email).Scan(
		&data.ID,
		&data.Email,
		&data.Name,
		&data.Password,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = constant.ErrUserNotFound
		}

		err = fmt.Errorf("user.repository.GetByEmail: failed to get user by email %s: %w", email, err)

		return
	}

	return
}
