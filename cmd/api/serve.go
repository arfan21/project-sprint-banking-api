package api

import (
	"github.com/arfan21/project-sprint-banking-api/config"
	"github.com/arfan21/project-sprint-banking-api/internal/server"
	dbpostgres "github.com/arfan21/project-sprint-banking-api/pkg/db/postgres"
)

func Serve() error {
	_, err := config.LoadConfig()
	if err != nil {
		return err
	}

	_, err = config.ParseConfig(config.GetViper())
	if err != nil {
		return err
	}

	// db, err := dbpostgres.NewPgx()
	// if err != nil {
	// 	return err
	// }
	db, err := dbpostgres.NewStdLib()
	if err != nil {
		return err
	}

	server := server.New(
		db,
	)
	return server.Run()
}
