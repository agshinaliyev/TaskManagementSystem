package main

import (
	"github.com/agshinaliyev/ms-app/db"
	"github.com/agshinaliyev/ms-app/server"
	"log"
)

func main() {

	if err := db.Init(); err != nil {
		log.Fatal("Error initializing database", err)
	}
	if err := server.Init(); err != nil {
		log.Fatal("Error initializing server", err)
	}

}
