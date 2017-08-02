package grouping

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getGroupRequest struct {
	ID string
}

type listGroupRequest struct{}

// groupResponse User must equal User type
type groupResponse struct {
	Group []Group `json:"content,omitempty"`
	Err   error   `json:"error,omitempty"`
}

func (r groupResponse) error() error { return r.Err }

func makeGetGroupEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getGroupRequest)
		result, err := s.GetGroup(req.ID)
		var groups []Group
		groups = append(groups, result)
		return groupResponse{Group: groups, Err: err}, nil
	}
}

func makeGetAllGroupEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(listGroupRequest)
		result, err := s.GetAllGroup()
		return groupResponse{Group: result, Err: err}, nil
	}
}
