package grouping

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

	getGroupHandler := kithttp.NewServer(
		makeGetGroupEndpoint(bs),
		decodeGetGroupRequest,
		encodeResponse,
		opts...,
	)

	getAllGroupHandler := kithttp.NewServer(
		makeGetAllGroupEndpoint(bs),
		decodeGetAllGroupRequest,
		encodeResponse,
		opts...,
	)

	addGroupHandler := kithttp.NewServer(
		makePostGroupEndpoint(bs),
		decodePostGroupRequest,
		encodeResponse,
		opts...,
	)

	deleteGroupHandler := kithttp.NewServer(
		makeDeleteGroupEndpoint(bs),
		decodeDeleteGroupRequest,
		encodeResponse,
		opts...,
	)

	deleteMultiGroupHandler := kithttp.NewServer(
		makeDeleteMultiGroupEndpoint(bs),
		decodeDeleteMultiGroupRequest,
		encodeResponse,
		opts...,
	)

	updateGroupHandler := kithttp.NewServer(
		makePutGroupEndpoint(bs),
		decodePutGroupRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/grouping/v1/group/{id}", getGroupHandler).Methods("GET")
	r.Handle("/grouping/v1/group", getAllGroupHandler).Methods("GET")
	r.Handle("/grouping/v1/group", addGroupHandler).Methods("POST")
	r.Handle("/grouping/v1/group/{id}", deleteGroupHandler).Methods("DELETE")
	r.Handle("/grouping/v1/group", deleteMultiGroupHandler).Methods("DELETE")
	r.Handle("/grouping/v1/group/{id}", updateGroupHandler).Methods("PUT")
	return r
}

var errBadRoute = errors.New("bad route")

func decodeGetGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}
	return getGroupRequest{ID: string(id)}, nil
}

func decodeGetAllGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return listGroupRequest{}, nil
}

func decodePostGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req postGroupRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Group); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeDeleteGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}
	return deleteGroupRequest{ID: string(id)}, nil
}

func decodeDeleteMultiGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var req deleteMutliGroupRequest
	if e := json.NewDecoder(r.Body).Decode(&req.ListId); e != nil {
		return nil, e
	}
	return req, nil
}

func decodePutGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}
	var group Group
	if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
		return nil, err
	}
	return putGroupRequest{
		ID:    id,
		Group: group,
	}, nil
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
