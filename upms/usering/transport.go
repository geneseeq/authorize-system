package usering

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

	getUserHandler := kithttp.NewServer(
		makeGetUserEndpoint(bs),
		decodeGetUserRequest,
		encodeResponse,
		opts...,
	)

	addUserHandler := kithttp.NewServer(
		makePostUserEndpoint(bs),
		decodePostUserRequest,
		encodeResponse,
		opts...,
	)

	getAllUserHandler := kithttp.NewServer(
		makeGetAllUserEndpoint(bs),
		decodeGetAllUserRequest,
		encodeResponse,
		opts...,
	)

	deleteUserHandler := kithttp.NewServer(
		makeDeleteUserEndpoint(bs),
		decodeDeleteUserRequest,
		encodeResponse,
		opts...,
	)

	updateUserHandler := kithttp.NewServer(
		makePutUserEndpoint(bs),
		decodePutUserRequest,
		encodeResponse,
		opts...,
	)

	updateMultiUserHandler := kithttp.NewServer(
		makePutMultiUserEndpoint(bs),
		decodePutMultiUserRequest,
		encodeResponse,
		opts...,
	)

	deleteMultiUserHandler := kithttp.NewServer(
		makeDeleteMultiUserEndpoint(bs),
		decodeDeleteMultiUserRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/usering/v1/user/{id}", getUserHandler).Methods("GET")
	r.Handle("/usering/v1/user", getAllUserHandler).Methods("GET")
	r.Handle("/usering/v1/user", addUserHandler).Methods("POST")
	r.Handle("/usering/v1/user/{id}", updateUserHandler).Methods("PUT")
	r.Handle("/usering/v1/user", updateMultiUserHandler).Methods("PUT")
	r.Handle("/usering/v1/user/{id}", deleteUserHandler).Methods("DELETE")
	r.Handle("/usering/v1/user", deleteMultiUserHandler).Methods("DELETE")

	return r
}

var errBadRoute = errors.New("bad route")

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}
	return getUserRequest{ID: string(id)}, nil
}

func decodePostUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req postUserRequest
	if e := json.NewDecoder(r.Body).Decode(&req.User); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeGetAllUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return listUserRequest{}, nil
}

func decodeDeleteUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}
	return deleteUserRequest{ID: string(id)}, nil
}

func decodeDeleteMultiUserRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var req deleteMutliUserRequest
	if e := json.NewDecoder(r.Body).Decode(&req.ListID); e != nil {
		return nil, e
	}
	return req, nil
}

func decodePutUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return nil, err
	}
	return putUserRequest{
		ID:   id,
		User: user,
	}, nil
}

func decodePutMultiUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req postUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req.User); err != nil {
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
