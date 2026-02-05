package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/hasnatinter/ticket-buyer/conn"
)

var db *sql.DB

type Venue struct {
	ID   int64
	Name string
}

type Performer struct {
	ID   int64
	Name string
}

type Event struct {
	ID           int64
	Name         string
	Description  string
	VenueId      int64
	StartTime    string
	Venue        Venue
	PerformerId  int64
	Performer    Performer
	TotalTickets int64
}

type EventFilter struct {
	StartDate string `validate:"omitempty,datetime=2006-01-02"`
	EndDate   string `validate:"omitempty,datetime=2006-01-02"`
	Venue     string `validate:"omitempty,alphanum"`
	Category  string `validate:"omitempty,alphanum"`
	Limit     string `validate:"required_with=Offset,omitempty,number"`
	Offset    string `validate:"omitempty,number"`
}

func main() {
	db = conn.ConnectDb()
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/events", getEvents)

	http.ListenAndServe("0.0.0.0:8080", r)
}

func getEvents(w http.ResponseWriter, req *http.Request) {
	input, err := ValidateInput(req)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	events, err := EventsQuery(input)
	if err != nil {
		log.Fatal(err)
	}

	response := map[string][]Event{"data": events}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

func ValidateInput(req *http.Request) (*EventFilter, error) {
	start := req.FormValue("start")
	end := req.FormValue("end")
	venue := req.FormValue("venue")
	limit := req.FormValue("limit")
	offset := req.FormValue("offset")
	category := req.FormValue("category")
	input := &EventFilter{
		StartDate: start,
		EndDate:   end,
		Venue:     venue,
		Limit:     limit,
		Offset:    offset,
		Category:  category,
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(input)
	if err != nil {
		return nil, err
	}
	return input, nil
}

func EventsQuery(input *EventFilter) ([]Event, error) {
	sql := "SELECT e.id, e.name, e.description, p.name as performer_name, e.start_time as start_time, v.name as venue_name, " +
		" (select count(*) from ticket WHERE status = 'available' AND event_id = e.id) as total_tickets" +
		" FROM event e" +
		" LEFT JOIN venue v ON v.id = e.venue_id" +
		" LEFT JOIN performer p ON p.id = e.performer_id" +
		" WHERE 1"
	var args []any
	if len(input.StartDate) > 0 {
		args = append(args, input.StartDate)
		sql = sql + " AND start_time >= ?"
	}
	if len(input.EndDate) > 0 {
		args = append(args, input.EndDate)
		sql = sql + " AND start_time <= ?"
	}
	if len(input.Venue) > 0 {
		args = append(args, input.Venue)
		sql = sql + " AND v.name = ?"
	}
	if len(input.Category) > 0 {
		args = append(args, input.Category)
		sql = sql + " AND e.category = ?"
	}
	sql += " ORDER BY e.start_time"
	if len(input.Limit) > 0 {
		args = append(args, input.Limit)
		sql = sql + " LIMIT ?"
	}
	if len(input.Offset) > 0 {
		args = append(args, input.Offset)
		sql = sql + " OFFSET ?"
	}
	stmt, err := db.Prepare(sql)
	rows, err := stmt.Query(args...)
	if err != nil {
		log.Fatal(err)
	}

	var events []Event
	defer rows.Close()
	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Performer.Name, &event.StartTime, &event.Venue.Name, &event.TotalTickets); err != nil {
			log.Fatal(err)
		}
		events = append(events, event)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return events, nil

}
