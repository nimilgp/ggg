package main

import (
	"log"
	"nimilgp/app/internal/env"
	"nimilgp/app/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			dsn:             env.GetString("DSN", "postgres://user:pass@localhost:5432/dbname?sslmode=disable"),
			maxOpenConns:    env.GetInt("DB_MAX_OPEN_CONNS", 25),
			maxIdleConns:    env.GetInt("DB_MAX_IDLE_CONNS", 25),
			maxIdleTimeSecs: env.GetString("DB_MAX_IDLE_TIME_SECS", "15min"),
		},
	}

	// db, err := db.New(
	// 	cfg.db.dsn,
	// 	cfg.db.maxOpenConns,
	// 	cfg.db.maxIdleConns,
	// 	cfg.db.maxIdleTimeSecs,
	// )
	// if err != nil {
	// 	log.Panic(err)
	// }
	// defer db.Close()

	store := store.NewPostgresStorage(nil)

	app := application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
