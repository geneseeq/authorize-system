package distribute

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

// MakeHandler returns a handler for the grouping service.
func MakeHandler(bs Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	getRoleDistributeHandler := kithttp.NewServer(
		makeGetRoleDistributeEndpoint(bs),
		decodeGetRoleDistributeRequest,
		encodeResponse,
		opts...,
	)

	getAllRoleDistributeHandler := kithttp.NewServer(
		makeGetAllRoleDistributeEndpoint(bs),
		decodeGetAllRoleDistributeRequest,
		encodeResponse,
		opts...,
	)

	addRoleDistributeHandler := kithttp.NewServer(
		makePostRoleDistributeEndpoint(bs),
		decodePostRoleDistributeRequest,
		encodeResponse,
		opts...,
	)

	deleteMultiRoleDistributeHandler := kithttp.NewServer(
		makeDeleteMultiRoleDistributeEndpoint(bs),
		decodeDeleteMultiRoleDistributeRequest,
		encodeResponse,
		opts...,
	)

	updateMultiRoleDistributeHandler := kithttp.NewServer(
		makePutMultiRoleDistributeEndpoint(bs),
		decodePutMultiRoleDistributeRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/distribute/v1/role/{id}", getRoleDistributeHandler).Methods("GET")
	r.Handle("/distribute/v1/role", getAllRoleDistributeHandler).Methods("GET")
	r.Handle("/distribute/v1/role", addRoleDistributeHandler).Methods("POST")
	r.Handle("/distribute/v1/role", deleteMultiRoleDistributeHandler).Methods("DELETE")
	r.Handle("/distribute/v1/role", updateMultiRoleDistributeHandler).Methods("PUT")
	return r
}

var errBadRoute = errors.New("bad route")

func decodeGetRoleDistributeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}
	return getRoleDistributeRequest{ID: string(id)}, nil
}

func decodeGetAllRoleDistributeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return listRoleDistributeRequest{}, nil
}

func decodePostRoleDistributeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req postRoleDistributeRequest
	if e := json.NewDecoder(r.Body).Decode(&req.RoleDistribute); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeDeleteMultiRoleDistributeRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var req deleteMutliRoleDistributeRequest
	if e := json.NewDecoder(r.Body).Decode(&req.ListID); e != nil {
		return nil, e
	}
	return req, nil
}

func decodePutMultiRoleDistributeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req postRoleDistributeRequest
	if err := json.NewDecoder(r.Body).Decode(&req.RoleDistribute); err != nil {
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
