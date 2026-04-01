package event_test

import (
	"app/internal/api/event"
	"app/internal/api/performer"
	"app/internal/api/venue"
	"app/pkg/helpers"
	"app/pkg/middleware"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func TestListEvents(t *testing.T) {
	db := helpers.SetUp(t)

	e := CreateEvent(t, db)

	r, err := http.NewRequest("GET", "/v1/events", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	eHanlder := event.New(db)
	middleware.ContentTypeJson(http.HandlerFunc(eHanlder.List)).ServeHTTP(w, r)
	res := w.Result()

	if res.StatusCode != http.StatusOK {
		t.Fatal("Error occurred during events fetch")
	}

	var eDto []event.EventDTO
	err = json.NewDecoder(res.Body).Decode(&eDto)
	if err != nil {
		t.Fatal(err)
	}

	if len(eDto) == 0 {
		t.Fatal("Empty events returned")
	}

	firstEvent := eDto[0]
	if firstEvent.Name != e.Name {
		t.Fatal("Event's name does not match")
	}
	if firstEvent.Category != e.Category.String {
		t.Fatal("Event's category does not match")
	}
}

func TestReadEvent(t *testing.T) {
	db := helpers.SetUp(t)

	e := CreateEvent(t, db)

	id := fmt.Sprintf("%d", e.ID)
	r, err := http.NewRequest("GET", "/v1/events/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", id)
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()

	eHanlder := event.New(db)
	middleware.ContentTypeJson(http.HandlerFunc(eHanlder.Read)).ServeHTTP(w, r)
	res := w.Result()

	if res.StatusCode != http.StatusOK {
		t.Fatal("Error occurred during events fetch")
	}

	var eDto event.EventDTO
	err = json.NewDecoder(res.Body).Decode(&eDto)
	if err != nil {
		t.Fatal(err)
	}

	if eDto.Name != e.Name {
		t.Errorf("Expected name %s, got %s", e.Name, eDto.Name)
	}
	if eDto.Category != e.Category.String {
		t.Errorf("Expected category %s, got %s", e.Category.String, eDto.Category)
	}
}

func CreateEvent(t *testing.T, db *gorm.DB) *event.Event {
	performerRepo := performer.NewRepository(db)
	p := &performer.Performer{
		Name: "New performer",
	}
	_, err := performerRepo.Create(t.Context(), p)
	if err != nil {
		t.Fatal(err)
	}

	venueRepo := venue.NewRepository(db)
	v := &venue.Venue{
		Name: "New venue",
	}
	_, err = venueRepo.Create(t.Context(), v)
	if err != nil {
		t.Fatal(err)
	}

	eventRepo := event.NewRepository(db)
	tm := time.Now().Add(1 * time.Hour).UTC()
	e := &event.Event{
		Name:        "New event",
		Description: sql.NullString{String: "New description", Valid: true},
		Category:    sql.NullString{String: "Musik", Valid: true},
		VenueId:     v.ID,
		PerformerId: p.ID,
		StartTime:   &tm,
	}
	_, err = eventRepo.Create(t.Context(), e)
	if err != nil {
		t.Fatal(err)
	}

	return e
}
