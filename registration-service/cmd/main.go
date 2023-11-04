package main

import (
	"log"

	"github.com/bersennaidoo/eopd/registration-service/application/rest/handler"
	"github.com/bersennaidoo/eopd/registration-service/application/rest/router"
	"github.com/bersennaidoo/eopd/registration-service/application/rest/server"
	"github.com/bersennaidoo/eopd/registration-service/infrastructure/msgbroker"
	"github.com/bersennaidoo/eopd/registration-service/infrastructure/storage"
	"github.com/bersennaidoo/eopd/registration-service/physical/config"
)

func main() {

	cfgdata := config.NewCFGData()

	cfg := config.New("registration-service", cfgdata)

	err := cfg.SetupConnectionToDB()
	if err != nil {
		log.Fatal(err)
	}

	err = cfg.SetupConnectionToNATS()
	if err != nil {
		log.Fatal(err)
	}

	db := cfg.DB()

	nc := cfg.NATS()

	natsmsgbroker := msgbroker.New(nc)

	store := storage.New(db)

	handle := handler.New(store, natsmsgbroker)

	route := router.New(handle)

	server := server.New(cfg.Cfgdata.Port, route)

	server.ListenAndServe()
}
