package router

import (
	"app/internal/api/resources/event"
	"app/internal/api/resources/health"
	"app/internal/server"
	"app/pkg/middleware/requestlog"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New(server *server.Server) *chi.Mux {
	l := server.Logger()
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/healthcheck", health.Read)

	r.Route("/v1", func(r chi.Router) {
		events := event.New(server.DB())
		r.Method("GET", "/events", requestlog.NewHandler(events.List, l))
		r.Method("GET", "/events/{id}", requestlog.NewHandler(events.Read, l))
	})
	return r
}
