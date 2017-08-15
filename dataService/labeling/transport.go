package labeling

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

	getLabelHandler := kithttp.NewServer(
		makeGetLabelEndpoint(bs),
		decodeGetLabelRequest,
		encodeResponse,
		opts...,
	)

	addLabelHandler := kithttp.NewServer(
		makePostLabelEndpoint(bs),
		decodePostLabelRequest,
		encodeResponse,
		opts...,
	)

	getAllLabelHandler := kithttp.NewServer(
		makeGetAllLabelEndpoint(bs),
		decodeGetAllLabelRequest,
		encodeResponse,
		opts...,
	)

	updateMultiLabelHandler := kithttp.NewServer(
		makePutMultiLabelEndpoint(bs),
		decodePutMultiLabelRequest,
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

	r.Handle("/labeling/v1/label/{id}", getLabelHandler).Methods("GET")
	r.Handle("/labeling/v1/label", getAllLabelHandler).Methods("GET")
	r.Handle("/labeling/v1/label", addLabelHandler).Methods("POST")
	r.Handle("/labeling/v1/label", updateMultiLabelHandler).Methods("PUT")
	r.Handle("/labeling/v1/label", deleteMultiLabelHandler).Methods("DELETE")

	return r
}

var errBadRoute = errors.New("bad route")

func decodeGetLabelRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}
	return getLabelRequest{LabelID: string(id)}, nil
}

func decodePostLabelRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req postLabelRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Label); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeGetAllLabelRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return listLabelRequest{}, nil
}

func decodeDeleteMultiLabelRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var req deleteMutliLabelRequest
	if e := json.NewDecoder(r.Body).Decode(&req.ListID); e != nil {
		return nil, e
	}
	return req, nil
}

func decodePutMultiLabelRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req postLabelRequest
	if err := json.NewDecoder(r.Body).Decode(&req.Label); err != nil {
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
