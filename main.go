package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/tmcmilla/simplebank/api"
	db "github.com/tmcmilla/simplebank/db/sqlc"
	"github.com/tmcmilla/simplebank/util"
)

// Need to be in a configuration

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDRIVER, config.DBSOURCE)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.SERVERADDRESS)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
