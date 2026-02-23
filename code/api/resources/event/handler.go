package event

import (
	"encoding/json"
	"fmt"
	"net/http"

	"app/code/api/resources/common/errors"
	"app/code/api/resources/ticket"
	validatorUtil "app/code/validator"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
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
	repo *Repository
}

func New(db *gorm.DB) *EventsApi {
	return &EventsApi{
		repo: NewRepository(db),
	}
}

// List godoc
//
//	@summary        List events
//	@description    List events
//	@tags           events
//	@accept         json
//	@produce        json
//	@success        200 {array}     event.EventDTO
//	@failure        500 {object}    errors.Error
//	@router         /events [get]
func (e *EventsApi) List(w http.ResponseWriter, req *http.Request) {
	input, err := ValidateInput(req)
	if err != nil {
		fmt.Println(err)
		respBody, err := json.Marshal(validatorUtil.ToErrResponse(err))
		if err != nil {
			errors.ServerError(w, errors.RespJSONDecodeFailure)
		}
		errors.ValidationError(w, respBody)
		return
	}
	ctx := req.Context()
	events, err := e.repo.ListWithTickets(input, ctx)
	if err != nil {
		errors.ServerError(w, errors.RespDBDataAccessFailure)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if len(events) == 0 {
		fmt.Fprint(w, "[]")
		return
	}

	if err := json.NewEncoder(w).Encode(events.ToDTO()); err != nil {
		errors.ServerError(w, errors.RespJSONEncodeFailure)
		return
	}
}

// Read godoc
//
//	@summary        Read event
//	@description    Read event
//	@tags           events
//	@accept         json
//	@produce        json
//	@param			id	path		string	true	"Event ID"
//	@success        200 {array}     event.EventDTO
//	@failure        500 {object}    errors.Error
//	@router         /events/{id} [get]
func (e *EventsApi) Read(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	if len(id) == 0 {
		errors.BadRequest(w, errors.RespInvalidURLParamID)
	}
	ctx := req.Context()
	event, err := e.repo.ReadWithTickets(id, ctx)
	if err != nil {
		errors.ServerError(w, errors.RespDBDataAccessFailure)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if event == nil {
		fmt.Fprint(w, "[]")
		return
	}
	tRepo := ticket.NewRepository(e.repo.db)
	tickets, _ := tRepo.ListForEvent(id, ctx)

	if err := json.NewEncoder(w).Encode(event.ToDTO(tickets)); err != nil {
		errors.ServerError(w, errors.RespJSONEncodeFailure)
		return
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
