package event

import (
	"encoding/json"
	"fmt"
	"net/http"

	errs "app/code/api/resources/common/errors"
	validatorUtil "app/code/validator"

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
//	@success        200 {array}     Event
//	@failure        500 {object}    error.Error
//	@router         /events [get]
func (e *EventsApi) Read(w http.ResponseWriter, req *http.Request) {
	input, err := ValidateInput(req)
	if err != nil {
		fmt.Println(err)
		respBody, err := json.Marshal(validatorUtil.ToErrResponse(err))
		if err != nil {
			errs.ServerError(w, errs.RespJSONDecodeFailure)
		}
		errs.ValidationError(w, respBody)
		return
	}
	ctx := req.Context()
	events, err := e.repo.ListWithTickets(input, ctx)
	if err != nil {
		errs.ServerError(w, errs.RespDBDataAccessFailure)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if len(events) == 0 {
		fmt.Fprint(w, "[]")
		return
	}

	if err := json.NewEncoder(w).Encode(events.ToDTO()); err != nil {
		errs.ServerError(w, errs.RespJSONEncodeFailure)
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
