package main

import "log"

func main() {
	storage, err := NewStorageConnection()
	if err != nil {
		log.Fatal(err)
	}
	if err := storage.Init(); err != nil {
		log.Fatal(err)
	}
	server := NewApiServer(":2026", storage)
	server.Run()
}
