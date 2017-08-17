package servicing

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getServiceRequest struct {
	ID string
}

type deleteMutliGroupRequest struct {
	ListID []string
}

type listServiceRequest struct{}

type postServiceRequest struct {
	Services []Services
}

type postServiceResponse struct {
	//omitempty表示字段值为空，则不输出到json串
	Status     int      `json:"status"`
	Content    string   `json:"content"`
	SucessedID []string `json:"sucessedid,omitempty"`
	FailedID   []string `json:"failedid,omitempty"`
	Err        error    `json:"err,omitempty"`
}

// serviceResponse User must equal User type
type serviceResponse struct {
	Services []Services `json:"content"`
	Err      error      `json:"error,omitempty"`
}

func (r serviceResponse) error() error { return r.Err }

func (r postServiceResponse) error() error { return r.Err }

func makeGetServiceEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getServiceRequest)
		result, err := s.GetService(req.ID)
		var services []Services
		services = append(services, result)
		return serviceResponse{Services: services, Err: err}, nil
	}
}

func makeGetAllServiceEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(listServiceRequest)
		result, err := s.GetAllService()
		return serviceResponse{Services: result, Err: err}, nil
	}
}

func makePostServiceEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postServiceRequest)
		sucessed, failed, err := s.PostService(req.Services)
		if err == nil {
			return postServiceResponse{
				SucessedID: sucessed,
				FailedID:   failed,
				Err:        err,
				Status:     200,
				Content:    "add service sucessed"}, nil
		}
		return postServiceResponse{
			SucessedID: sucessed,
			FailedID:   failed,
			Err:        err,
			Status:     300,
			Content:    "add service failed"}, nil
	}
}

func makeDeleteMultiServiceEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteMutliGroupRequest)
		sucessed, failed, err := s.DeleteMultiService(req.ListID)
		if err == nil {
			return postServiceResponse{
				SucessedID: sucessed,
				FailedID:   failed,
				Err:        err,
				Status:     200,
				Content:    "delete servcie sucessed"}, nil
		}
		return postServiceResponse{
			SucessedID: sucessed,
			FailedID:   failed,
			Err:        err,
			Status:     300,
			Content:    "delete servcie failed"}, nil
	}
}

func makePutMultiServiceEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postServiceRequest)
		sucessed, failed, err := s.PutMultiService(req.Services)
		if err == nil {
			return postServiceResponse{
				SucessedID: sucessed,
				FailedID:   failed,
				Err:        err,
				Status:     200,
				Content:    "update service sucessed"}, nil
		}
		return postServiceResponse{
			SucessedID: sucessed,
			FailedID:   failed,
			Err:        err,
			Status:     300,
			Content:    "update service failed"}, nil
	}
}
