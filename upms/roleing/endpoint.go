package roleing

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type baseRoleRequest struct {
	ID string
}

type listRoleRequest struct{}

type baseMutliRoleRequest struct {
	ListID []string
}

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
	Content    string   `json:"content"`
	SucessedID []string `json:"sucessedid,omitempty"`
	FailedID   []string `json:"failedid,omitempty"`
	Err        error    `json:"err,omitempty"`
}

func (r baseResponse) error() error { return r.Err }

func (r roleResponse) error() error { return r.Err }

func makeGetRoleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(baseRoleRequest)
		result, err := s.GetRole(req.ID)
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
				FailedID:   failedIds,
				Content:    "add role sucessed"}, nil
		}
		return roleResponse{
			SucessedID: sucessedIds,
			Err:        err,
			Status:     300,
			FailedID:   failedIds,
			Content:    "add role failed"}, nil
	}
}

func makeDeleteRoleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(baseRoleRequest)
		err := s.DeleteRole(req.ID)
		if err == nil {
			return roleResponse{
				Err:     err,
				Status:  200,
				Content: "delete role sucessed"}, nil
		}
		return roleResponse{
			Err:     err,
			Status:  300,
			Content: "delete role failed"}, nil
	}
}

func makeDeleteMultiRoleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(baseMutliRoleRequest)
		sucessedIds, failedIds, err := s.DeleteMultiRole(req.ListID)
		if err == nil {
			return roleResponse{
				SucessedID: sucessedIds,
				Err:        err,
				Status:     200,
				FailedID:   failedIds,
				Content:    "delete mutli role sucessed"}, nil
		}
		return roleResponse{
			SucessedID: sucessedIds,
			Err:        err,
			Status:     300,
			FailedID:   failedIds,
			Content:    "delete mutli role failed"}, nil
	}
}

func makePutRoleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(putRoleRequest)
		err := s.PutRole(req.ID, req.Role)
		if err == nil {
			return roleResponse{Err: err, Status: 200, Content: "update role sucessed"}, nil
		}
		return roleResponse{Err: err, Status: 300, Content: "update role failed"}, nil
	}
}

func makePutMultiRoleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postRoleRequest)
		sucessedIds, failedIds, err := s.PutMultiRole(req.Role)
		if err == nil {
			return roleResponse{
				SucessedID: sucessedIds,
				Err:        err,
				Status:     200,
				FailedID:   failedIds,
				Content:    "update role sucessed"}, nil
		}
		return roleResponse{
			SucessedID: sucessedIds,
			Err:        err,
			Status:     300,
			FailedID:   failedIds,
			Content:    "update role failed"}, nil
	}
}
