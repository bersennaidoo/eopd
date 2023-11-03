package main

import (
	"log"

	"github.com/bersennaidoo/eopd/registration-service/application/rest/handler"
	"github.com/bersennaidoo/eopd/registration-service/application/rest/router"
	"github.com/bersennaidoo/eopd/registration-service/application/rest/server"
	"github.com/bersennaidoo/eopd/registration-service/infrastructure/msgbroker"
	"github.com/bersennaidoo/eopd/registration-service/infrastructure/storage"
	"github.com/bersennaidoo/eopd/registration-service/physical/config"
	"github.com/nats-io/nats.go"
)

func main() {

	cfg := config.New("registration-service")

	err := cfg.SetupConnectionToDB("mysql", "root:bersen@/eopd")
	if err != nil {
		log.Fatal(err)
	}

	err = cfg.SetupConnectionToNATS(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	db := cfg.DB()

	nc := cfg.NATS()

	natsmsgbroker := msgbroker.New(nc)

	store := storage.New(db)

	handle := handler.New(store, natsmsgbroker)

	route := router.New(handle)

	server := server.New(":3000", route)

	server.ListenAndServe()
}
