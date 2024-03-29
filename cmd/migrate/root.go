package migration

import (
	"github.com/arfan21/project-sprint-banking-api/migration"
	dbpostgres "github.com/arfan21/project-sprint-banking-api/pkg/db/postgres"
)

func initMigration() (*migration.Migration, error) {
	db, err := dbpostgres.NewStdLib()
	if err != nil {
		return nil, err
	}

	// sqlDB := stdlib.OpenDBFromPool(db)

	migration, err := migration.New(db)
	if err != nil {
		return nil, err
	}

	return migration, nil

}
