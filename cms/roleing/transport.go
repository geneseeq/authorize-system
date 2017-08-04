package roleing

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

	// deleteMultiGroupHandler := kithttp.NewServer(
	// 	makeDeleteMultiGroupEndpoint(bs),
	// 	decodeDeleteMultiGroupRequest,
	// 	encodeResponse,
	// 	opts...,
	// )

	// updateGroupHandler := kithttp.NewServer(
	// 	makePutGroupEndpoint(bs),
	// 	decodePutGroupRequest,
	// 	encodeResponse,
	// 	opts...,
	// )

	// updateMultiGroupHandler := kithttp.NewServer(
	// 	makePutMultiGroupEndpoint(bs),
	// 	decodePutMultiGroupRequest,
	// 	encodeResponse,
	// 	opts...,
	// )

	r := mux.NewRouter()

	r.Handle("/roleing/v1/role/{id}", getRoleHandler).Methods("GET")
	r.Handle("/roleing/v1/role", getAllRoleHandler).Methods("GET")
	r.Handle("/roleing/v1/role", addRoleHandler).Methods("POST")
	r.Handle("/roleing/v1/role/{id}", deleteRoleHandler).Methods("DELETE")
	// r.Handle("/grouping/v1/group", deleteMultiGroupHandler).Methods("DELETE")
	// r.Handle("/grouping/v1/group/{id}", updateGroupHandler).Methods("PUT")
	// r.Handle("/grouping/v1/group", updateMultiGroupHandler).Methods("PUT")
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

// func decodeDeleteMultiGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {

// 	var req deleteMutliGroupRequest
// 	if e := json.NewDecoder(r.Body).Decode(&req.ListId); e != nil {
// 		return nil, e
// 	}
// 	return req, nil
// }

// func decodePutGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {
// 	vars := mux.Vars(r)
// 	id, ok := vars["id"]
// 	if !ok {
// 		return nil, errBadRoute
// 	}
// 	var group Group
// 	if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
// 		return nil, err
// 	}
// 	return putGroupRequest{
// 		ID:    id,
// 		Group: group,
// 	}, nil
// }

// func decodePutMultiGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {
// 	var req postGroupRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req.Group); err != nil {
// 		return nil, err
// 	}
// 	return req, nil
// }

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
