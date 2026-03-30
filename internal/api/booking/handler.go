package booking

import (
	"app/pkg/errors"
	"encoding/json"
	"fmt"
	"net/http"

	l "app/pkg/logger"
	validatorUtil "app/pkg/validator"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CreateFilter struct {
	UserName    string `json:"user_name" validate:"required,max=255"`
	UserAddress string `json:"user_address" validate:"required,max=255"`
	Tickets     []int  `json:"tickets" validate:"required,dive"`
}

type BookingApi struct {
	repo   *Repository
	logger *l.Logger
}

func New(db *gorm.DB, logger *l.Logger) *BookingApi {
	return &BookingApi{
		repo:   NewRepository(db),
		logger: logger,
	}
}

// Create godoc
//
//	@summary        Create booking
//	@description    Create booking
//	@tags           bookings
//	@accept         json
//	@produce        json
//	@param			body body		CreateFilter true "Booking form"
//	@success        201 {array}     booking.BookingDTO
//	@failure		400 {object}	errors.Error
//	@failure		422	{object}	errors.Errors
//	@failure        500 {object}    errors.Error
//	@router         /bookings/ [post]
func (a *BookingApi) Create(w http.ResponseWriter, r *http.Request) {
	input := &CreateFilter{}
	json.NewDecoder(r.Body).Decode(&input)

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(input)
	if err != nil {
		fmt.Println(err)
		respBody, err := json.Marshal(validatorUtil.ToErrResponse(err))
		if err != nil {
			errors.ServerError(w, errors.RespJSONDecodeFailure)
		}
		errors.ValidationError(w, respBody)
		return
	}

	booking, err := a.repo.AddBooking(*input, r.Context())
	if err != nil {
		a.logger.Error().Str("error", err.Error()).Msg("")
		errors.ServerError(w, errors.RespDBDataInsertFailure)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(booking.ToDTO(booking.Ticket)); err != nil {
		errors.ServerError(w, errors.RespJSONEncodeFailure)
		return
	}
}
