package baseing

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getBaseDataRequest struct {
	ID string
}

type deleteMutliBaseDataRequest struct {
	ListID []string
}

type listBaseDataRequest struct{}

type postBaseDataRequest struct {
	BaseData []BaseData
}

type postBaseDataResponse struct {
	//omitempty表示字段值为空，则不输出到json串
	Status     int      `json:"status"`
	Content    string   `json:"content"`
	SucessedID []string `json:"sucessedid,omitempty"`
	FailedID   []string `json:"failedid,omitempty"`
	Err        error    `json:"err,omitempty"`
}

// baseDataResponse User must equal User type
type baseDataResponse struct {
	BaseData []BaseData `json:"content"`
	Err      error      `json:"error,omitempty"`
}

func (r baseDataResponse) error() error { return r.Err }

func (r postBaseDataResponse) error() error { return r.Err }

func makeGetBaseDatarEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getBaseDataRequest)
		result, err := s.GetBaseData(req.ID)
		var datas []BaseData
		datas = append(datas, result)
		return baseDataResponse{BaseData: datas, Err: err}, nil
	}
}

func makeGetAllBaseDataEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(listBaseDataRequest)
		result, err := s.GetAllBaseData()
		return baseDataResponse{BaseData: result, Err: err}, nil
	}
}

func makePostBaseDataEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postBaseDataRequest)
		sucessed, failed, err := s.PostBaseData(req.BaseData)
		if err == nil {
			return postBaseDataResponse{
				SucessedID: sucessed,
				FailedID:   failed,
				Err:        err,
				Status:     200,
				Content:    "add data sucessed"}, nil
		}
		return postBaseDataResponse{
			SucessedID: sucessed,
			FailedID:   failed,
			Err:        err,
			Status:     300,
			Content:    "add user failed"}, nil
	}
}

func makeDeleteMultiBaseDataEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteMutliBaseDataRequest)
		sucessed, failed, err := s.DeleteMultiBaseData(req.ListID)
		if err == nil {
			return postBaseDataResponse{
				SucessedID: sucessed,
				FailedID:   failed,
				Err:        err,
				Status:     200,
				Content:    "add data sucessed"}, nil
		}
		return postBaseDataResponse{
			SucessedID: sucessed,
			FailedID:   failed,
			Err:        err,
			Status:     300,
			Content:    "add user failed"}, nil
	}
}

func makePutMultiBaseDataEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postBaseDataRequest)
		sucessed, failed, err := s.PutMultiBaseData(req.BaseData)
		if err == nil {
			return postBaseDataResponse{
				SucessedID: sucessed,
				FailedID:   failed,
				Err:        err,
				Status:     200,
				Content:    "add data sucessed"}, nil
		}
		return postBaseDataResponse{
			SucessedID: sucessed,
			FailedID:   failed,
			Err:        err,
			Status:     300,
			Content:    "add user failed"}, nil
	}
}
