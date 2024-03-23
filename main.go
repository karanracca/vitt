package main

import (
	"log"
	"vitt/pkg/cmd"
	"vitt/pkg/store"
)

func main() {
	store, err := store.Init("./vitt.db")
	if err != nil {
		log.Fatal(err)
	}

	cmd.Init(store)

	defer store.DB.Close()
}
