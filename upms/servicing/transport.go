package servicing

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/geneseeq/authorize-system/upms/user"
	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

// MakeHandler returns a handler for the booking service.
func MakeHandler(bs Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	getServiceHandler := kithttp.NewServer(
		makeGetServiceEndpoint(bs),
		decodeGetServiceRequest,
		encodeResponse,
		opts...,
	)

	addServiceHandler := kithttp.NewServer(
		makePostServiceEndpoint(bs),
		decodePostServiceRequest,
		encodeResponse,
		opts...,
	)

	getAllServiceHandler := kithttp.NewServer(
		makeGetAllServiceEndpoint(bs),
		decodeGetAllServiceRequest,
		encodeResponse,
		opts...,
	)

	updateMultiServiceHandler := kithttp.NewServer(
		makePutMultiServiceEndpoint(bs),
		decodePutMultiServiceRequest,
		encodeResponse,
		opts...,
	)

	deleteMultiServiceHandler := kithttp.NewServer(
		makeDeleteMultiServiceEndpoint(bs),
		decodeDeleteMultiServiceRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/servicing/v1/service/{id}", getServiceHandler).Methods("GET")
	r.Handle("/servicing/v1/service", getAllServiceHandler).Methods("GET")
	r.Handle("/servicing/v1/service", addServiceHandler).Methods("POST")
	r.Handle("/servicing/v1/service", updateMultiServiceHandler).Methods("PUT")
	r.Handle("/servicing/v1/service", deleteMultiServiceHandler).Methods("DELETE")

	return r
}

var errBadRoute = errors.New("bad route")

func decodeGetServiceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}
	return getServiceRequest{ID: string(id)}, nil
}

func decodePostServiceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req postServiceRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Services); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeGetAllServiceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return listServiceRequest{}, nil
}

func decodeDeleteMultiServiceRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var req deleteMutliGroupRequest
	if e := json.NewDecoder(r.Body).Decode(&req.ListID); e != nil {
		return nil, e
	}
	return req, nil
}

func decodePutMultiServiceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req postServiceRequest
	if err := json.NewDecoder(r.Body).Decode(&req.Services); err != nil {
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
