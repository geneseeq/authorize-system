package fielding

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/geneseeq/authorize-system/cms/user"
	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

// MakeHandler returns a handler for the booking service.
func MakeHandler(bs Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	getFieldHandler := kithttp.NewServer(
		makeGetFieldEndpoint(bs),
		decodeGetFieldRequest,
		encodeResponse,
		opts...,
	)

	addFieldHandler := kithttp.NewServer(
		makePostFieldEndpoint(bs),
		decodePostFieldRequest,
		encodeResponse,
		opts...,
	)

	getAllFieldHandler := kithttp.NewServer(
		makeGetAllFieldEndpoint(bs),
		decodeGetAllFieldRequest,
		encodeResponse,
		opts...,
	)

	updateMultiFieldHandler := kithttp.NewServer(
		makePutMultiFieldEndpoint(bs),
		decodePutMultiFieldRequest,
		encodeResponse,
		opts...,
	)

	deleteMultiFieldHandler := kithttp.NewServer(
		makeDeleteMultiFieldEndpoint(bs),
		decodeDeleteMultiFieldRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/fielding/v1/field/{id}", getFieldHandler).Methods("GET")
	r.Handle("/fielding/v1/field", getAllFieldHandler).Methods("GET")
	r.Handle("/fielding/v1/field", addFieldHandler).Methods("POST")
	r.Handle("/fielding/v1/field", updateMultiFieldHandler).Methods("PUT")
	r.Handle("/fielding/v1/field", deleteMultiFieldHandler).Methods("DELETE")

	return r
}

var errBadRoute = errors.New("bad route")

func decodeGetFieldRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}
	return getFieldRequest{ID: string(id)}, nil
}

func decodePostFieldRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req postFieldRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Field); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeGetAllFieldRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return listFieldRequest{}, nil
}

func decodeDeleteMultiFieldRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var req deleteMutliFieldRequest
	if e := json.NewDecoder(r.Body).Decode(&req.ListID); e != nil {
		return nil, e
	}
	return req, nil
}

func decodePutMultiFieldRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req postFieldRequest
	if err := json.NewDecoder(r.Body).Decode(&req.Field); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	switch err {
	case user.ErrUnknown:
		w.WriteHeader(http.StatusNotFound)
	case ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
