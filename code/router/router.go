package router

import (
	"app/code/api/resources/event"
	"app/code/api/resources/health"
	"app/code/middleware/requestlog"
	"app/code/server"

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
		r.Method("GET", "/events", requestlog.NewHandler(events.Read, l))
	})
	return r
}
