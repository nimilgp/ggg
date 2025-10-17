package main

import (
	"log"
	"nimilgp/app/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
