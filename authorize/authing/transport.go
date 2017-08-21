package authing

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

	getTokenHandler := kithttp.NewServer(
		makeGetTokenEndpoint(bs),
		decodeGetTokenRequest,
		encodeResponse,
		opts...,
	)

	addTokenHandler := kithttp.NewServer(
		makePostTokenEndpoint(bs),
		decodePostTokenRequest,
		encodeResponse,
		opts...,
	)

	getAllTokenHandler := kithttp.NewServer(
		makeGetAllTokenEndpoint(bs),
		decodeGetAllTokenRequest,
		encodeResponse,
		opts...,
	)

	updateMultiTokenHandler := kithttp.NewServer(
		makePutMultiTokenEndpoint(bs),
		decodePutMultiTokenRequest,
		encodeResponse,
		opts...,
	)

	deleteMultiTokenHandler := kithttp.NewServer(
		makeDeleteMultiTokenEndpoint(bs),
		decodeDeleteMultiTokenRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/authing/v1/authorize/user/{id}/token", getTokenHandler).Methods("GET")
	r.Handle("/authing/v1/authorize/user/token", getAllTokenHandler).Methods("GET")
	r.Handle("/authing/v1/authorize/user/token", addTokenHandler).Methods("POST")
	r.Handle("/authing/v1/authorize/user/token", updateMultiTokenHandler).Methods("PUT")
	r.Handle("/authing/v1/authorize/user/token", deleteMultiTokenHandler).Methods("DELETE")

	return r
}

var errBadRoute = errors.New("bad route")

func decodeGetTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}
	return getTokenRequest{ID: string(id)}, nil
}

func decodePostTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req postTokenRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Token); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeGetAllTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return listTokenRequest{}, nil
}

func decodeDeleteMultiTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var req deleteMutliTokenRequest
	if e := json.NewDecoder(r.Body).Decode(&req.ListID); e != nil {
		return nil, e
	}
	return req, nil
}

func decodePutMultiTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req postTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req.Token); err != nil {
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
