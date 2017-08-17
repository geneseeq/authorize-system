package roles

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

	getAuthorityHandler := kithttp.NewServer(
		makeGetAuthorityEndpoint(bs),
		decodeGetAuthorityRequest,
		encodeResponse,
		opts...,
	)

	getAllAuthorityHandler := kithttp.NewServer(
		makeGetAllAuthorityEndpoint(bs),
		decodeGetAllAuthorityRequest,
		encodeResponse,
		opts...,
	)

	addAuthorityHandler := kithttp.NewServer(
		makePostAuthorityEndpoint(bs),
		decodePostAuthorityRequest,
		encodeResponse,
		opts...,
	)

	deleteMultiAuthorityHandler := kithttp.NewServer(
		makeDeleteMultiAuthorityEndpoint(bs),
		decodeDeleteMultiAuthorityRequest,
		encodeResponse,
		opts...,
	)

	updateMultiAuthorityHandler := kithttp.NewServer(
		makePutMultiAuthorityEndpoint(bs),
		decodePutMultiAuthorityRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/releation/v1/role/{role_id}/authority", getAuthorityHandler).Methods("GET")
	r.Handle("/releation/v1/role/authority", getAllAuthorityHandler).Methods("GET")
	r.Handle("/releation/v1/role/authority", addAuthorityHandler).Methods("POST")
	r.Handle("/releation/v1/role/authority", deleteMultiAuthorityHandler).Methods("DELETE")
	r.Handle("/releation/v1/role/authority", updateMultiAuthorityHandler).Methods("PUT")
	return r
}

var errBadRoute = errors.New("bad route")

func decodeGetAuthorityRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	roleID, ok := vars["role_id"]
	if !ok {
		return nil, errBadRoute
	}
	return baseAuthorityRequest{RoleID: string(roleID)}, nil
}

func decodeGetAllAuthorityRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return listAuthorityRequest{}, nil
}

func decodePostAuthorityRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req postAuthorityRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Authority); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeDeleteMultiAuthorityRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var req deleteAuthorityRequest
	if e := json.NewDecoder(r.Body).Decode(&req.DeleteData); e != nil {
		return nil, e
	}
	return req, nil
}

func decodePutMultiAuthorityRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var req postAuthorityRequest
	if err := json.NewDecoder(r.Body).Decode(&req.Authority); err != nil {
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
