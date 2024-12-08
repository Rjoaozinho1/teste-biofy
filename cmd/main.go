package main

import (
	"log"

	"github.com/Rjoaozinho1/teste-biofy/api"
	"github.com/Rjoaozinho1/teste-biofy/storage"
)

func main() {
	storage, err := storage.NewStorageConnection()
	if err != nil {
		log.Fatal(err)
	}
	if err := storage.Init(); err != nil {
		log.Fatal(err)
	}
	server := api.NewApiServer(":2026", storage)
	server.Run()
}
