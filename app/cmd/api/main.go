package main

import (
	"nimilgp/app/internal/env"
	"nimilgp/app/internal/store"

	"go.uber.org/zap"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description

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

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

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
		logger: logger,
	}

	mux := app.mount()
	logger.Fatal(app.run(mux))
}
