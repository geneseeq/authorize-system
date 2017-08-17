package dataing

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

	getDataSetHandler := kithttp.NewServer(
		makeGetSetEndpoint(bs),
		decodeGetSetRequest,
		encodeResponse,
		opts...,
	)

	addDataSetHandler := kithttp.NewServer(
		makePostSetEndpoint(bs),
		decodePostSetRequest,
		encodeResponse,
		opts...,
	)

	getAllDataSetHandler := kithttp.NewServer(
		makeGetAllSetEndpoint(bs),
		decodeGetAllSetRequest,
		encodeResponse,
		opts...,
	)

	updateMultiDataSetHandler := kithttp.NewServer(
		makePutMultiSetEndpoint(bs),
		decodePutMultiSetRequest,
		encodeResponse,
		opts...,
	)

	deleteMultiDataSetHandler := kithttp.NewServer(
		makeDeleteMultiSetEndpoint(bs),
		decodeDeleteMultiSetRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/seting/v1/data/{id}", getDataSetHandler).Methods("GET")
	r.Handle("/seting/v1/data", getAllDataSetHandler).Methods("GET")
	r.Handle("/seting/v1/data", addDataSetHandler).Methods("POST")
	r.Handle("/seting/v1/data", updateMultiDataSetHandler).Methods("PUT")
	r.Handle("/seting/v1/data", deleteMultiDataSetHandler).Methods("DELETE")

	return r
}

var errBadRoute = errors.New("bad route")

func decodeGetSetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}
	return getSetRequest{ID: string(id)}, nil
}

func decodePostSetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req postSetRequest
	if e := json.NewDecoder(r.Body).Decode(&req.DataSet); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeGetAllSetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return listSetRequest{}, nil
}

func decodeDeleteMultiSetRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var req deleteMutliSetRequest
	if e := json.NewDecoder(r.Body).Decode(&req.ListID); e != nil {
		return nil, e
	}
	return req, nil
}

func decodePutMultiSetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req postSetRequest
	if err := json.NewDecoder(r.Body).Decode(&req.DataSet); err != nil {
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
