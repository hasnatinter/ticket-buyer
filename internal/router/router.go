package router

import (
	"app/internal/api/booking"
	"app/internal/api/event"
	"app/internal/api/health"
	"app/internal/server"
	"app/pkg/middleware"
	"app/pkg/middleware/requestlog"

	"github.com/go-chi/chi/v5"
	m "github.com/go-chi/chi/v5/middleware"
)

func New(server *server.Server) *chi.Mux {
	l := server.Logger()
	r := chi.NewRouter()

	r.Use(m.RequestID)
	r.Use(m.Logger)
	r.Use(m.Recoverer)
	r.Use(middleware.ContentTypeJson)

	r.Get("/healthcheck", health.Read)

	r.Route("/v1", func(r chi.Router) {
		events := event.New(server.DB())
		r.Method("GET", "/events", requestlog.NewHandler(events.List, l))
		r.Method("GET", "/events/{id}", requestlog.NewHandler(events.Read, l))

		booking := booking.New(server.DB(), l)
		r.Method("POST", "/bookings", requestlog.NewHandler(booking.Create, l))
	})

	return r
}
