package baseing

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/geneseeq/authorize-system/dataService/data"
	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

// MakeHandler returns a handler for the booking service.
func MakeHandler(bs Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	getBaseDataHandler := kithttp.NewServer(
		makeGetBaseDataEndpoint(bs),
		decodeGetBaseDataRequest,
		encodeResponse,
		opts...,
	)

	addBaseDataHandler := kithttp.NewServer(
		makePostBaseDataEndpoint(bs),
		decodePostBaseDataRequest,
		encodeResponse,
		opts...,
	)

	getAllBaseDataHandler := kithttp.NewServer(
		makeGetAllBaseDataEndpoint(bs),
		decodeGetAllBaseDataRequest,
		encodeResponse,
		opts...,
	)

	updateMultiBaseDataHandler := kithttp.NewServer(
		makePutMultiBaseDataEndpoint(bs),
		decodePutMultiBaseDataRequest,
		encodeResponse,
		opts...,
	)

	deleteMultiBaseDataHandler := kithttp.NewServer(
		makeDeleteMultiBaseDataEndpoint(bs),
		decodeDeleteMultiBaseDataRequest,
		encodeResponse,
		opts...,
	)

	deleteMultiLabelHandler := kithttp.NewServer(
		makeDeleteMultiLabelEndpoint(bs),
		decodeDeleteMultiLabelRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/baseing/v1/data/{id}", getBaseDataHandler).Methods("GET")
	r.Handle("/baseing/v1/data", getAllBaseDataHandler).Methods("GET")
	r.Handle("/baseing/v1/data", addBaseDataHandler).Methods("POST")
	r.Handle("/baseing/v1/data", updateMultiBaseDataHandler).Methods("PUT")
	r.Handle("/baseing/v1/data", deleteMultiBaseDataHandler).Methods("DELETE")
	r.Handle("/baseing/v1/data/label", deleteMultiLabelHandler).Methods("DELETE")
	return r
}

var errBadRoute = errors.New("bad route")

func decodeGetBaseDataRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}
	return getBaseDataRequest{ID: string(id)}, nil
}

func decodePostBaseDataRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req postBaseDataRequest
	if e := json.NewDecoder(r.Body).Decode(&req.BaseData); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeGetAllBaseDataRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return listBaseDataRequest{}, nil
}

func decodeDeleteMultiBaseDataRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var req deleteMutliBaseDataRequest
	if e := json.NewDecoder(r.Body).Decode(&req.ListID); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeDeleteMultiLabelRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var req deleteMutliLabelRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Label); e != nil {
		return nil, e
	}
	return req, nil
}

func decodePutMultiBaseDataRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req postBaseDataRequest
	if err := json.NewDecoder(r.Body).Decode(&req.BaseData); err != nil {
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
	case data.ErrUnknown:
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
