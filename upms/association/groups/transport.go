package groups

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
func MakeHandler(us Service, rs Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	getGroupUserHandler := kithttp.NewServer(
		makeGetGroupEndpoint(us),
		decodeGetGroupRequest,
		encodeResponse,
		opts...,
	)

	getAllGroupUserHandler := kithttp.NewServer(
		makeGetAllGroupEndpoint(us),
		decodeGetAllGroupRequest,
		encodeResponse,
		opts...,
	)

	addGroupUserHandler := kithttp.NewServer(
		makePostGroupEndpoint(us),
		decodePostGroupRequest,
		encodeResponse,
		opts...,
	)

	deleteMultiGroupUserHandler := kithttp.NewServer(
		makeDeleteMultiGroupEndpoint(us),
		decodeDeleteMultiGroupRequest,
		encodeResponse,
		opts...,
	)

	updateMultiGroupUserHandler := kithttp.NewServer(
		makePutMultiGroupEndpoint(us),
		decodePutMultiGroupRequest,
		encodeResponse,
		opts...,
	)

	getGroupRoleHandler := kithttp.NewServer(
		makeGetGroupEndpoint(rs),
		decodeGetGroupRequest,
		encodeResponse,
		opts...,
	)

	getAllGroupRoleHandler := kithttp.NewServer(
		makeGetAllGroupEndpoint(rs),
		decodeGetAllGroupRequest,
		encodeResponse,
		opts...,
	)

	addGroupRoleHandler := kithttp.NewServer(
		makePostGroupEndpoint(rs),
		decodePostGroupRequest,
		encodeResponse,
		opts...,
	)

	deleteMultiGroupRoleHandler := kithttp.NewServer(
		makeDeleteMultiGroupEndpoint(rs),
		decodeDeleteMultiGroupRequest,
		encodeResponse,
		opts...,
	)

	updateMultiGroupRoleHandler := kithttp.NewServer(
		makePutMultiGroupEndpoint(rs),
		decodePutMultiGroupRequest,
		encodeResponse,
		opts...,
	)
	r := mux.NewRouter()

	r.Handle("/releation/v1/group/{group_id}/user", getGroupUserHandler).Methods("GET")
	r.Handle("/releation/v1/group/user", getAllGroupUserHandler).Methods("GET")
	r.Handle("/releation/v1/group/user", addGroupUserHandler).Methods("POST")
	r.Handle("/releation/v1/group/user", deleteMultiGroupUserHandler).Methods("DELETE")
	r.Handle("/releation/v1/group/user", updateMultiGroupUserHandler).Methods("PUT")
	r.Handle("/releation/v1/group/{group_id}/role", getGroupRoleHandler).Methods("GET")
	r.Handle("/releation/v1/group/role", getAllGroupRoleHandler).Methods("GET")
	r.Handle("/releation/v1/group/role", addGroupRoleHandler).Methods("POST")
	r.Handle("/releation/v1/group/role", deleteMultiGroupRoleHandler).Methods("DELETE")
	r.Handle("/releation/v1/group/role", updateMultiGroupRoleHandler).Methods("PUT")
	return r
}

var errBadRoute = errors.New("bad route")

func decodeGetGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	groupID, ok := vars["group_id"]
	if !ok {
		return nil, errBadRoute
	}
	return baseGroupRequest{GroupID: string(groupID)}, nil
}

func decodeGetAllGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return listGroupRequest{}, nil
}

func decodePostGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req postGroupRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Groups); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeDeleteMultiGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var req postGroupRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Groups); e != nil {
		return nil, e
	}
	return req, nil
}

func decodePutMultiGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req postGroupRequest
	if err := json.NewDecoder(r.Body).Decode(&req.Groups); err != nil {
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
