package main

import (
	"fmt"
	"net/http"
	"nimilgp/app/docs"
	"nimilgp/app/internal/store"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

type application struct {
	config config
	store  store.Storage
	logger *zap.SugaredLogger
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn             string
	maxOpenConns    int
	maxIdleConns    int
	maxIdleTimeSecs string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	vr := chi.NewRouter()
	vr.Get("/healthcheck", app.healthcheckHandler)

	docURL := fmt.Sprintf("%s/swagger/doc.json", app.config.addr)
	vr.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(docURL)))

	r.Mount("/v1", vr)

	return r
}

func (app *application) run(mux http.Handler) error {
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/v1"
	srv := http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	app.logger.Infow("Starting has started", "addr", app.config.addr)

	return srv.ListenAndServe()
}
