//go:generate oapi-codegen --config=cfg.yaml ./api.yaml

package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	middleware "github.com/oapi-codegen/echo-middleware"

	"github.com/robertjshirts/ratioed/backend/accounts/api"
	"github.com/robertjshirts/ratioed/backend/accounts/db"
	"github.com/robertjshirts/ratioed/backend/accounts/utils"
)

func main() {

	database := db.NewDatabase()
	defer database.Close()

	server := api.NewServer(database)

	e := echo.New()

	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatal(err)
	}

	e.Use(middleware.OapiRequestValidator(swagger))

	api.RegisterHandlers(e, server)

	port := utils.GetEnv("PORT")
	address := "0.0.0.0:" + port
	log.Println("Listening on:", address)
	log.Fatal(e.Start(address))
}
