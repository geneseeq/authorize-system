package roleing

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

	getRoleHandler := kithttp.NewServer(
		makeGetRoleEndpoint(bs),
		decodeGetRoleRequest,
		encodeResponse,
		opts...,
	)

	getAllRoleHandler := kithttp.NewServer(
		makeGetAllRoleEndpoint(bs),
		decodeGetAllRoleRequest,
		encodeResponse,
		opts...,
	)

	addRoleHandler := kithttp.NewServer(
		makePostRoleEndpoint(bs),
		decodePostRoleRequest,
		encodeResponse,
		opts...,
	)

	deleteRoleHandler := kithttp.NewServer(
		makeDeleteRoleEndpoint(bs),
		decodeDeleteRoleRequest,
		encodeResponse,
		opts...,
	)

	deleteMultiRoleHandler := kithttp.NewServer(
		makeDeleteMultiRoleEndpoint(bs),
		decodeDeleteMultiRoleRequest,
		encodeResponse,
		opts...,
	)

	updateRoleHandler := kithttp.NewServer(
		makePutRoleEndpoint(bs),
		decodePutRoleRequest,
		encodeResponse,
		opts...,
	)

	updateMultiRoleHandler := kithttp.NewServer(
		makePutMultiRoleEndpoint(bs),
		decodePutMultiRoleRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/roleing/v1/role/{id}", getRoleHandler).Methods("GET")
	r.Handle("/roleing/v1/role", getAllRoleHandler).Methods("GET")
	r.Handle("/roleing/v1/role", addRoleHandler).Methods("POST")
	r.Handle("/roleing/v1/role/{id}", deleteRoleHandler).Methods("DELETE")
	r.Handle("/roleing/v1/role", deleteMultiRoleHandler).Methods("DELETE")
	r.Handle("/roleing/v1/role/{id}", updateRoleHandler).Methods("PUT")
	r.Handle("/roleing/v1/role", updateMultiRoleHandler).Methods("PUT")
	return r
}

var errBadRoute = errors.New("bad route")

func decodeGetRoleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}
	return baseRoleRequest{ID: string(id)}, nil
}

func decodeGetAllRoleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return listRoleRequest{}, nil
}

func decodePostRoleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req postRoleRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Role); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeDeleteRoleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}
	return baseRoleRequest{ID: string(id)}, nil
}

func decodeDeleteMultiRoleRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var req baseMutliRoleRequest
	if e := json.NewDecoder(r.Body).Decode(&req.ListID); e != nil {
		return nil, e
	}
	return req, nil
}

func decodePutRoleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}
	var role Role
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		return nil, err
	}
	return putRoleRequest{
		ID:   id,
		Role: role,
	}, nil
}

func decodePutMultiRoleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req postRoleRequest
	if err := json.NewDecoder(r.Body).Decode(&req.Role); err != nil {
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
