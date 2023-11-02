package main

import (
	"log"

	"github.com/bersennaidoo/eopd/registration-service/application/rest/handler"
	"github.com/bersennaidoo/eopd/registration-service/application/rest/router"
	"github.com/bersennaidoo/eopd/registration-service/application/rest/server"
	"github.com/bersennaidoo/eopd/registration-service/foundation/config"
	"github.com/bersennaidoo/eopd/registration-service/infrastructure/storage"
)

func main() {

	cfg := config.New("registration-service")

	err := cfg.SetupConnectionToDB("mysql", "root:bersen@/eopd")
	if err != nil {
		log.Fatal(err)
	}

	db := cfg.DB()

	store := storage.New(db)

	handle := handler.New(store)

	route := router.New(handle)

	server := server.New(":3000", route)

	server.ListenAndServe()
}
