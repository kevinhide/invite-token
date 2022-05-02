package main

import (
	"log"
	"net/http"
	"os"

	"invite-token/config"
	"invite-token/src/handlers"
	"invite-token/utils/drivers/db"
)

func main() {
	cnf := &db.Config{
		URL:       config.DatabaseURL,
		MaxDBConn: config.MaxDBConn,
	}

	err := db.Init(cnf)
	if err != nil {
		log.Println("Unable to connect to database. Err: ", err)
		os.Exit(1)
	}

	log.Println("Listening to Port: " + config.Port)
	http.ListenAndServe(":"+config.Port, handlers.GetRouter())
}
