package fielding

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getFieldRequest struct {
	ID string
}

type deleteMutliFieldRequest struct {
	ListID []string
}

type listFieldRequest struct{}

type postFieldRequest struct {
	Field []Field
}

type postFieldResponse struct {
	//omitempty表示字段值为空，则不输出到json串
	Status     int      `json:"status"`
	Content    string   `json:"content"`
	SucessedID []string `json:"sucessedid,omitempty"`
	FailedID   []string `json:"failedid,omitempty"`
	Err        error    `json:"err,omitempty"`
}

type fieldResponse struct {
	Field []Field `json:"content"`
	Err   error   `json:"error,omitempty"`
}

func (r fieldResponse) error() error { return r.Err }

func (r postFieldResponse) error() error { return r.Err }

func makeGetFieldEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getFieldRequest)
		result, err := s.GetField(req.ID)
		var datas []Field
		datas = append(datas, result)
		return fieldResponse{Field: datas, Err: err}, nil
	}
}

func makeGetAllFieldEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(listFieldRequest)
		result, err := s.GetAllField()
		return fieldResponse{Field: result, Err: err}, nil
	}
}

func makePostFieldEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postFieldRequest)
		sucessed, failed, err := s.PostField(req.Field)
		if err == nil {
			return postFieldResponse{
				SucessedID: sucessed,
				FailedID:   failed,
				Err:        err,
				Status:     200,
				Content:    "add data sucessed"}, nil
		}
		return postFieldResponse{
			SucessedID: sucessed,
			FailedID:   failed,
			Err:        err,
			Status:     300,
			Content:    "add user failed"}, nil
	}
}

func makeDeleteMultiFieldEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteMutliFieldRequest)
		sucessed, failed, err := s.DeleteMultiField(req.ListID)
		if err == nil {
			return postFieldResponse{
				SucessedID: sucessed,
				FailedID:   failed,
				Err:        err,
				Status:     200,
				Content:    "add data sucessed"}, nil
		}
		return postFieldResponse{
			SucessedID: sucessed,
			FailedID:   failed,
			Err:        err,
			Status:     300,
			Content:    "add user failed"}, nil
	}
}

func makePutMultiFieldEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postFieldRequest)
		sucessed, failed, err := s.PutMultiField(req.Field)
		if err == nil {
			return postFieldResponse{
				SucessedID: sucessed,
				FailedID:   failed,
				Err:        err,
				Status:     200,
				Content:    "add data sucessed"}, nil
		}
		return postFieldResponse{
			SucessedID: sucessed,
			FailedID:   failed,
			Err:        err,
			Status:     300,
			Content:    "add user failed"}, nil
	}
}
