package event

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
)

type EventFilter struct {
	StartDate string `validate:"omitempty,datetime=2006-01-02"`
	EndDate   string `validate:"omitempty,datetime=2006-01-02"`
	Venue     string `validate:"omitempty,alphanum"`
	Category  string `validate:"omitempty,alphanum"`
	Limit     string `validate:"required_with=Offset,omitempty,number"`
	Offset    string `validate:"omitempty,number"`
}

type EventsApi struct {
	db *pgx.Conn
}

func New(db *pgx.Conn) *EventsApi {
	return &EventsApi{
		db,
	}
}

// List godoc
//
//	@summary        List events
//	@description    List events
//	@tags           events
//	@accept         json
//	@produce        json
//	@success        200 {array}     Event
//	@failure        500 {object}    error.Error
//	@router         /events [get]
func (e *EventsApi) Read(w http.ResponseWriter, req *http.Request) {
	input, err := ValidateInput(req)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	events, err := EventsQuery(input, e.db)
	if err != nil {
		log.Fatal(err)
	}

	response := map[string][]Event{"data": events}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		log.Fatal(err.Error())
	}
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

func EventsQuery(input *EventFilter, conn *pgx.Conn) ([]Event, error) {
	sql := "SELECT e.id, e.name, e.description, p.name as performer_name, e.start_time as start_time, v.name as venue_name, " +
		" (select count(*) from ticket WHERE status = 'available' AND event_id = e.id) as total_tickets" +
		" FROM event e" +
		" LEFT JOIN venue v ON v.id = e.venue_id" +
		" LEFT JOIN performer p ON p.id = e.performer_id" +
		" WHERE 1=1"
	args := make(map[string]any, 0)
	if len(input.StartDate) > 0 {
		args["start"] = input.StartDate
		sql = sql + " AND start_time >= @start"
	}
	if len(input.EndDate) > 0 {
		args["end"] = input.EndDate
		sql = sql + " AND start_time <= @end"
	}
	if len(input.Venue) > 0 {
		args["venue"] = input.Venue
		sql = sql + ` AND v.name = @venue`
	}
	if len(input.Category) > 0 {
		args["category"] = input.EndDate
		sql = sql + " AND e.category = @category"
	}
	sql += " ORDER BY e.start_time"
	if len(input.Limit) > 0 {
		args["limit"] = input.Limit
		sql = sql + " LIMIT @limit"
	}
	if len(input.Offset) > 0 {
		args["offset"] = input.Offset
		sql = sql + " OFFSET @offset"
	}

	rows, err := conn.Query(context.Background(), sql, pgx.NamedArgs(args))
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
