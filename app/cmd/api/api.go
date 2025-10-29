package main

import (
	"log"
	"net/http"
	"nimilgp/app/internal/store"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	store  store.Storage
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

	r.Mount("/v1", vr)

	return r
}

func (app *application) run(mux http.Handler) error {
	srv := http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Print("Starting server on ", app.config.addr)

	return srv.ListenAndServe()
}
