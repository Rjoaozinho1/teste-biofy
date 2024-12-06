package main

import "log"

func main() {
	storage, err := newStorageConnection()
	if err != nil {
		log.Fatal(err)
	}
	if err := storage.Init(); err != nil {
		log.Fatal(err)
	}
	server := newApiServer(":2026", storage)
	server.Run()
}
