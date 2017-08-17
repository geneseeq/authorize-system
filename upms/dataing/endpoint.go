package dataing

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getSetRequest struct {
	ID string
}

type deleteUserRequest struct {
	ID string
}

type deleteMutliSetRequest struct {
	ListID []string
}

// type putUserRequest struct {
// 	ID   string
// 	User User
// }

type listSetRequest struct{}

type postSetRequest struct {
	DataSet []DataSet
}

type postSetResponse struct {
	//omitempty表示字段值为空，则不输出到json串
	Status     int      `json:"status"`
	Content    string   `json:"content"`
	SucessedID []string `json:"sucessedid,omitempty"`
	FailedID   []string `json:"failedid,omitempty"`
	Err        error    `json:"err,omitempty"`
}

type setResponse struct {
	DataSet []DataSet `json:"content"`
	Err     error     `json:"error,omitempty"`
}

func (r setResponse) error() error { return r.Err }

func (r postSetResponse) error() error { return r.Err }

func makeGetSetEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getSetRequest)
		result, err := s.GetDataSet(req.ID)
		var set []DataSet
		set = append(set, result)
		return setResponse{DataSet: set, Err: err}, nil
	}
}

func makeGetAllSetEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(listSetRequest)
		result, err := s.GetAllDataSet()
		return setResponse{DataSet: result, Err: err}, nil
	}
}

func makePostSetEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postSetRequest)
		sucessedID, failedID, err := s.PostDataSet(req.DataSet)
		if err == nil {
			return postSetResponse{
				SucessedID: sucessedID,
				FailedID:   failedID,
				Err:        err,
				Status:     200,
				Content:    "add user sucessed"}, nil
		}
		return postSetResponse{
			SucessedID: sucessedID,
			FailedID:   failedID,
			Err:        err,
			Status:     300,
			Content:    "add user failed"}, nil
	}
}

func makeDeleteMultiSetEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteMutliSetRequest)
		sucessedID, failedID, err := s.DeleteMultiDataSet(req.ListID)
		if err == nil {
			return postSetResponse{
				SucessedID: sucessedID,
				FailedID:   failedID,
				Err:        err,
				Status:     200,
				Content:    "delete dataset sucessed"}, nil
		}
		return postSetResponse{
			SucessedID: sucessedID,
			FailedID:   failedID,
			Err:        err,
			Status:     300,
			Content:    "delete dataset failed"}, nil
	}
}

func makePutMultiSetEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postSetRequest)
		sucessedID, failedID, err := s.PutMultiDataSet(req.DataSet)
		if err == nil {
			return postSetResponse{
				SucessedID: sucessedID,
				FailedID:   failedID,
				Err:        err,
				Status:     200,
				Content:    "update dataset sucessed"}, nil
		}
		return postSetResponse{
			SucessedID: sucessedID,
			FailedID:   failedID,
			Err:        err,
			Status:     300,
			Content:    "update dataset failed"}, nil
	}
}
