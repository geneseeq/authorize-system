package roles

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type baseAuthorityRequest struct {
	RoleID string
}

type listAuthorityRequest struct{}

type putAuthorityRequest struct {
	ID        string
	Authority Authority
}

type postAuthorityRequest struct {
	Authority []Authority
}

type deleteAuthorityRequest struct {
	DeleteData []DeleteData
}

type baseResponse struct {
	Authority []Authority `json:"content,omitempty"`
	Err       error       `json:"error,omitempty"`
}

type authorityResponse struct {
	//omitempty表示字段值为空，则不输出到json串
	Status     int      `json:"status"`
	SucessedID []string `json:"sucessedid,omitempty"`
	FailedID   []string `json:"failedid,omitempty"`
	Err        error    `json:"err,omitempty"`
}

func (r baseResponse) error() error { return r.Err }

func (r authorityResponse) error() error { return r.Err }

func makeGetAuthorityEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(baseAuthorityRequest)
		result, err := s.GetAuthorityFromRole(req.RoleID)
		var a []Authority
		a = append(a, result)
		return baseResponse{Authority: a, Err: err}, nil
	}
}

func makeGetAllAuthorityEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(listAuthorityRequest)
		result, err := s.GetAllAuthority()
		return baseResponse{Authority: result, Err: err}, nil
	}
}

func makePostAuthorityEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postAuthorityRequest)
		sucessedIds, failedIds, err := s.PostAuthority(req.Authority)
		if err == nil {
			return authorityResponse{
				SucessedID: sucessedIds,
				Err:        err,
				Status:     200,
				FailedID:   failedIds}, nil
		}
		return authorityResponse{
			SucessedID: sucessedIds,
			Err:        err,
			Status:     300,
			FailedID:   failedIds}, nil
	}
}

func makeDeleteMultiAuthorityEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteAuthorityRequest)
		sucessedIds, failedIds, err := s.DeleteMultiAuthority(req.DeleteData)
		if err == nil {
			return authorityResponse{
				SucessedID: sucessedIds,
				Err:        err,
				Status:     200,
				FailedID:   failedIds}, nil
		}
		return authorityResponse{
			SucessedID: sucessedIds,
			Err:        err,
			Status:     300,
			FailedID:   failedIds}, nil
	}
}

func makePutMultiAuthorityEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postAuthorityRequest)
		sucessedIds, failedIds, err := s.PutMultiAuthority(req.Authority)
		if err == nil {
			return authorityResponse{
				SucessedID: sucessedIds,
				Err:        err,
				Status:     200,
				FailedID:   failedIds}, nil
		}
		return authorityResponse{
			SucessedID: sucessedIds,
			Err:        err,
			Status:     300,
			FailedID:   failedIds}, nil
	}
}
