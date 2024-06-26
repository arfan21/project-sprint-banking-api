package dbpostgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/arfan21/project-sprint-banking-api/config"
	"github.com/arfan21/project-sprint-banking-api/pkg/logger"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

const (
	maxOpenConnection = 60
	connMaxLifetime   = 120
	maxIdleConns      = 30
	connMaxIdleTime   = 20
)

func NewPgx() (db *pgxpool.Pool, err error) {
	url := config.Get().Database.GetURL()
	if config.Get().Env == "dev" {
		url += "?sslmode=disable"
	} else {
		url += "?sslmode=verify-full&sslrootcert=ap-southeast-1-bundle.pem"
	}

	ctx := context.Background()
	pgConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		err = fmt.Errorf("failed to parse database config: %w", err)
		return nil, err
	}
	pgConfig.MaxConns = maxOpenConnection
	pgConfig.MaxConnIdleTime = connMaxIdleTime * time.Second
	pgConfig.MaxConnLifetime = connMaxLifetime * time.Second

	db, err = pgxpool.NewWithConfig(ctx, pgConfig)
	if err != nil {
		err = fmt.Errorf("failed to connect to database: %w", err)
		return nil, err
	}

	if err = db.Ping(ctx); err != nil {
		err = fmt.Errorf("failed to ping database: %w", err)
		return nil, err
	}

	logger.Log(ctx).Info().Msg("dbpostgres: connection established")
	return db, nil
}

type Queryer interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
}

type QueryerStdLib interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func NewStdLib() (db *sql.DB, err error) {
	url := config.Get().Database.GetURL()
	if config.Get().Env == "dev" {
		url += "?sslmode=disable"
	} else {
		url += fmt.Sprintf("?%s", config.Get().Database.Params)
	}

	db, err = sql.Open("pgx", url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	err = prometheus.Register(collectors.NewDBStatsCollector(db, config.Get().Database.Name))
	if err != nil {
		return nil, err
	}

	return db, nil
}
