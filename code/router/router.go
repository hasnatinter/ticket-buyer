package router

import (
	"app/code/api/resources/event"
	"app/code/api/resources/health"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

func New(conn *gorm.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/healthcheck", health.Read)

	r.Route("/v1", func(r chi.Router) {
		events := event.New(conn)
		r.Get("/events", events.Read)
	})
	return r
}
