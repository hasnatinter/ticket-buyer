package errors

import "net/http"

type Error struct {
	Error string `json:"error"`
}

type Errors struct {
	Errors []string `json:"errors"`
}

var (
	RespDBDataInsertFailure = []byte(`{"error": "db data insert failure"}`)
	RespDBDataAccessFailure = []byte(`{"error": "db data access failure"}`)
	RespDBDataUpdateFailure = []byte(`{"error": "db data update failure"}`)
	RespDBDataRemoveFailure = []byte(`{"error": "db data remove failure"}`)

	RespJSONEncodeFailure = []byte(`{"error": "json encode failure"}`)
	RespJSONDecodeFailure = []byte(`{"error": "json decode failure"}`)

	RespInvalidURLParamID = []byte(`{"error": "invalid url param-id"}`)
)

func ServerError(w http.ResponseWriter, resp []byte) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(resp)
}

func BadRequest(w http.ResponseWriter, resp []byte) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write(resp)
}

func ValidationError(w http.ResponseWriter, resp []byte) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write(resp)
}
