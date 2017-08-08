package users

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type baseRoleRequest struct {
	UserID string
}

type listRoleRequest struct{}

// type roleIDDict struct {
// 	UserID string
// 	RoleID []string
// }

// type baseMutliRoleRequest struct {
// 	ListID []Role
// }

type putRoleRequest struct {
	ID   string
	Role Role
}

type postRoleRequest struct {
	Role []Role
}

type baseResponse struct {
	Role []Role `json:"content,omitempty"`
	Err  error  `json:"error,omitempty"`
}

type roleResponse struct {
	//omitempty表示字段值为空，则不输出到json串
	Status     int      `json:"status"`
	SucessedID []string `json:"sucessedid,omitempty"`
	FailedID   []string `json:"failedid,omitempty"`
	Err        error    `json:"err,omitempty"`
}

func (r baseResponse) error() error { return r.Err }

func (r roleResponse) error() error { return r.Err }

func makeGetRoleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(baseRoleRequest)
		result, err := s.GetRoleFromUser(req.UserID)
		var roles []Role
		roles = append(roles, result)
		return baseResponse{Role: roles, Err: err}, nil
	}
}

func makeGetAllRoleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(listRoleRequest)
		result, err := s.GetAllRole()
		return baseResponse{Role: result, Err: err}, nil
	}
}

func makePostRoleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postRoleRequest)
		sucessedIds, failedIds, err := s.PostRole(req.Role)
		if err == nil {
			return roleResponse{
				SucessedID: sucessedIds,
				Err:        err,
				Status:     200,
				FailedID:   failedIds}, nil
		}
		return roleResponse{
			SucessedID: sucessedIds,
			Err:        err,
			Status:     300,
			FailedID:   failedIds}, nil
	}
}

// func makeDeleteRoleEndpoint(s Service) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(baseRoleRequest)
// 		err := s.DeleteRole(req.ID)
// 		if err == nil {
// 			return roleResponse{Err: err, Status: 200}, nil
// 		}
// 		return roleResponse{Err: err, Status: 300}, nil
// 	}
// }

func makeDeleteMultiRoleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postRoleRequest)
		sucessedIds, failedIds, err := s.DeleteMultiRole(req.Role)
		if err == nil {
			return roleResponse{
				SucessedID: sucessedIds,
				Err:        err,
				Status:     200,
				FailedID:   failedIds}, nil
		}
		return roleResponse{
			SucessedID: sucessedIds,
			Err:        err,
			Status:     300,
			FailedID:   failedIds}, nil
	}
}

// func makePutRoleEndpoint(s Service) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(putRoleRequest)
// 		err := s.PutRole(req.ID, req.Role)
// 		if err == nil {
// 			return roleResponse{Err: err, Status: 200}, nil
// 		}
// 		return roleResponse{Err: err, Status: 300}, nil
// 	}
// }

func makePutMultiRoleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postRoleRequest)
		sucessedIds, failedIds, err := s.PutMultiRole(req.Role)
		if err == nil {
			return roleResponse{
				SucessedID: sucessedIds,
				Err:        err,
				Status:     200,
				FailedID:   failedIds}, nil
		}
		return roleResponse{
			SucessedID: sucessedIds,
			Err:        err,
			Status:     300,
			FailedID:   failedIds}, nil
	}
}
