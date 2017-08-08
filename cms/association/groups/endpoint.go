package groups

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type baseGroupRequest struct {
	GroupID string
}

type listGroupRequest struct{}

// type roleIDDict struct {
// 	UserID string
// 	RoleID []string
// }

// type baseMutliRoleRequest struct {
// 	ListID []Role
// }

// type putRoleRequest struct {
// 	ID   string
// 	Role Role
// }

type postGroupRequest struct {
	Groups []Groups
}

type baseResponse struct {
	Groups []Groups `json:"content,omitempty"`
	Err    error    `json:"error,omitempty"`
}

type groupResponse struct {
	//omitempty表示字段值为空，则不输出到json串
	Status     int      `json:"status"`
	SucessedID []string `json:"sucessedid,omitempty"`
	FailedID   []string `json:"failedid,omitempty"`
	Err        error    `json:"err,omitempty"`
}

func (r baseResponse) error() error { return r.Err }

func (r groupResponse) error() error { return r.Err }

func makeGetGroupEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(baseGroupRequest)
		result, err := s.GetDataFromGroup(req.GroupID)
		var groups []Groups
		groups = append(groups, result)
		return baseResponse{Groups: groups, Err: err}, nil
	}
}

func makeGetAllGroupEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(listGroupRequest)
		result, err := s.GetAllData()
		return baseResponse{Groups: result, Err: err}, nil
	}
}

func makePostGroupEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postGroupRequest)
		sucessedIds, failedIds, err := s.PostData(req.Groups)
		if err == nil {
			return groupResponse{
				SucessedID: sucessedIds,
				Err:        err,
				Status:     200,
				FailedID:   failedIds}, nil
		}
		return groupResponse{
			SucessedID: sucessedIds,
			Err:        err,
			Status:     300,
			FailedID:   failedIds}, nil
	}
}

func makeDeleteMultiGroupEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postGroupRequest)
		sucessedIds, failedIds, err := s.DeleteMultiData(req.Groups)
		if err == nil {
			return groupResponse{
				SucessedID: sucessedIds,
				Err:        err,
				Status:     200,
				FailedID:   failedIds}, nil
		}
		return groupResponse{
			SucessedID: sucessedIds,
			Err:        err,
			Status:     300,
			FailedID:   failedIds}, nil
	}
}

func makePutMultiGroupEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postGroupRequest)
		sucessedIds, failedIds, err := s.PutMultiData(req.Groups)
		if err == nil {
			return groupResponse{
				SucessedID: sucessedIds,
				Err:        err,
				Status:     200,
				FailedID:   failedIds}, nil
		}
		return groupResponse{
			SucessedID: sucessedIds,
			Err:        err,
			Status:     300,
			FailedID:   failedIds}, nil
	}
}
