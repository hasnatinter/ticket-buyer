package router

import (
	"app/code/api"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/healthcheck", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("-"))
	})
	r.Route("/v1", func(r chi.Router) {
		events := api.New(db)
		r.Get("/events", events.GetEvents)
	})
	return r
}
