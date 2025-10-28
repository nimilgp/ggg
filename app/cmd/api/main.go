package main

import (
	"log"
	"nimilgp/app/internal/env"
	"nimilgp/app/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewPostgresStorage(nil)

	app := application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
