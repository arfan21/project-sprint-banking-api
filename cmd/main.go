package main

import (
	"fmt"
	"os"

	"github.com/arfan21/project-sprint-banking-api/cmd/api"
	migration "github.com/arfan21/project-sprint-banking-api/cmd/migrate"
)

// @title project-sprint-banking-api
// @version 1.0
// @description This is a sample server cell for project-sprint-banking-api.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.synapsis.id
// @contact.email
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {

	args := os.Args

	if len(args) > 2 {
		if args[1] == "migrate" && args[2] == "up" {
			migration.Up()
			return
		}

		if args[1] == "migrate" && args[2] == "down" {
			migration.Down()
			return
		}

		if args[1] == "migrate" && args[2] == "fresh" {
			migration.Fresh()
			return

		}

		if args[1] == "migrate" && args[2] == "drop" {
			migration.Drop()
			return
		}
	}

	err := api.Serve()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
