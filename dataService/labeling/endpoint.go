package labeling

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getLabelRequest struct {
	LabelID string
}

type deleteMutliLabelRequest struct {
	ListID []string
}

type listLabelRequest struct{}

type postLabelRequest struct {
	Label []Label
}

type postLabelResponse struct {
	//omitempty表示字段值为空，则不输出到json串
	Status     int      `json:"status"`
	Content    string   `json:"content"`
	SucessedID []string `json:"sucessedid,omitempty"`
	FailedID   []string `json:"failedid,omitempty"`
	Err        error    `json:"err,omitempty"`
}

type labelResponse struct {
	Label []Label `json:"content"`
	Err   error   `json:"error,omitempty"`
}

func (r labelResponse) error() error { return r.Err }

func (r postLabelResponse) error() error { return r.Err }

func makeGetLabelEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getLabelRequest)
		result, err := s.GetLabel(req.LabelID)
		var datas []Label
		datas = append(datas, result)
		return labelResponse{Label: datas, Err: err}, nil
	}
}

func makeGetAllLabelEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(listLabelRequest)
		result, err := s.GetAllLabel()
		return labelResponse{Label: result, Err: err}, nil
	}
}

func makePostLabelEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postLabelRequest)
		sucessed, failed, err := s.PostLabel(req.Label)
		if err == nil {
			return postLabelResponse{
				SucessedID: sucessed,
				FailedID:   failed,
				Err:        err,
				Status:     200,
				Content:    "add data sucessed"}, nil
		}
		return postLabelResponse{
			SucessedID: sucessed,
			FailedID:   failed,
			Err:        err,
			Status:     300,
			Content:    "add user failed"}, nil
	}
}

func makeDeleteMultiLabelEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteMutliLabelRequest)
		sucessed, failed, err := s.DeleteMultiLabel(req.ListID)
		if err == nil {
			return postLabelResponse{
				SucessedID: sucessed,
				FailedID:   failed,
				Err:        err,
				Status:     200,
				Content:    "add data sucessed"}, nil
		}
		return postLabelResponse{
			SucessedID: sucessed,
			FailedID:   failed,
			Err:        err,
			Status:     300,
			Content:    "add user failed"}, nil
	}
}

func makePutMultiLabelEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postLabelRequest)
		sucessed, failed, err := s.PutMultiLabel(req.Label)
		if err == nil {
			return postLabelResponse{
				SucessedID: sucessed,
				FailedID:   failed,
				Err:        err,
				Status:     200,
				Content:    "add data sucessed"}, nil
		}
		return postLabelResponse{
			SucessedID: sucessed,
			FailedID:   failed,
			Err:        err,
			Status:     300,
			Content:    "add user failed"}, nil
	}
}
