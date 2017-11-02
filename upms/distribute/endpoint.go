package distribute

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getRoleDistributeRequest struct {
	ID string
}

type deleteRoleDistributeRequest struct {
	ID string
}

type deleteMutliRoleDistributeRequest struct {
	ListID []string
}

type listRoleDistributeRequest struct{}

type postRoleDistributeRequest struct {
	RoleDistribute []RoleDistribute
}

type putRoleDistributeRequest struct {
	ID             string
	RoleDistribute RoleDistribute
}

// roleDistributeResponse User must equal User type
type roleDistributeResponse struct {
	RoleDistribute []RoleDistribute `json:"content"`
	Err            error            `json:"error,omitempty"`
}

type postRoleDistributeResponse struct {
	//omitempty表示字段值为空，则不输出到json串
	Status     int      `json:"status"`
	Content    string   `json:"content"`
	SucessedId []string `json:"sucessedid,omitempty"`
	FailedID   []string `json:"failedid,omitempty"`
	Err        error    `json:"err,omitempty"`
}

func (r roleDistributeResponse) error() error { return r.Err }

func (r postRoleDistributeResponse) error() error { return r.Err }

func makeGetRoleDistributeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getRoleDistributeRequest)
		result, err := s.GetRoleDistribute(req.ID)
		var distributes []RoleDistribute
		distributes = append(distributes, result)
		return roleDistributeResponse{RoleDistribute: distributes, Err: err}, nil
	}
}

func makeGetAllRoleDistributeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(listRoleDistributeRequest)
		result, err := s.GetAllRoleDistribute()
		return roleDistributeResponse{RoleDistribute: result, Err: err}, nil
	}
}

func makePostRoleDistributeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postRoleDistributeRequest)
		sucessed, failed, err := s.PostRoleDistribute(req.RoleDistribute)
		if err == nil {
			return postRoleDistributeResponse{
				SucessedId: sucessed,
				FailedID:   failed,
				Err:        err,
				Status:     200,
				Content:    "add group sucessed"}, nil
		}
		return postRoleDistributeResponse{
			SucessedId: sucessed,
			FailedID:   failed,
			Err:        err,
			Status:     300,
			Content:    "add group failed"}, nil
	}
}

func makeDeleteMultiRoleDistributeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteMutliRoleDistributeRequest)
		sucessed, failed, err := s.DeleteMultiRoleDistribute(req.ListID)
		if err == nil {
			return postRoleDistributeResponse{
				SucessedId: sucessed,
				FailedID:   failed,
				Err:        err,
				Status:     200,
				Content:    "delete mutli group sucessed"}, nil
		}
		return postRoleDistributeResponse{
			SucessedId: sucessed,
			FailedID:   failed,
			Err:        err,
			Status:     300,
			Content:    "delete mutli group failed"}, nil
	}
}

func makePutMultiRoleDistributeEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postRoleDistributeRequest)
		sucessed, failed, err := s.PutMultiRoleDistribute(req.RoleDistribute)
		if err == nil {
			return postRoleDistributeResponse{
				SucessedId: sucessed,
				FailedID:   failed,
				Err:        err,
				Status:     200,
				Content:    "update group sucessed"}, nil
		}
		return postRoleDistributeResponse{
			SucessedId: sucessed,
			FailedID:   failed,
			Err:        err,
			Status:     300,
			Content:    "update group failed"}, nil
	}
}
